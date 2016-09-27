package main

import (
	"github.com/gorilla/mux"
	"net/http"
	log "github.com/Sirupsen/logrus"
	"github.com/asaskevich/govalidator"
	"github.com/oschwald/geoip2-golang"
	"net"
	"encoding/json"
	"os/exec"
	"github.com/21stio/go-ip-microservice/events"
	"flag"
	"fmt"
	"github.com/robfig/cron"
	"os"
)

type IpRequest struct {
	Ip string `valid:"required,ipv4"`
}

const (
	GEOIP_DATABASE_FILE="geoip_database.mmdb"
)

var (
	APPLICATION_PORT=flag.Int("APPLICATION_PORT", 8000, "port to listen on")
	CRON_SCHEDULE=flag.String("CRON_SCHEDULE", "@daily", "port to listen on")
)

func init()  {
	flag.Parse()
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	if(!isGeoipDatabasePresent()) {
		downloadGeoipDatabase()
	}

	startCronjob()

	router := mux.NewRouter()
	router.HandleFunc("/v1/{ip}", IpHandler)

	events.LogStartedServer(*APPLICATION_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *APPLICATION_PORT), router))
}

func IpHandler(writer http.ResponseWriter, request *http.Request) {
	ip := mux.Vars(request)["ip"]

	events.LogStartedIpLookup(ip)

	ip_request := &IpRequest{
		Ip: ip,
	}

	_, err := govalidator.ValidateStruct(ip_request)
	if err != nil {
		events.LogFailedIpLookup(err, ip)

		http.Error(writer, "not a valid ip adress", 400)

		return
	}

	geo_information := *getGeoInformation(ip_request.Ip)

	ip_lookup := map[string]interface{}{
		"location": map[string]interface{}{
			"country": map[string]string{
				"iso_code": geo_information.Country.IsoCode,
				"name": geo_information.Country.Names["en"],
			},
			"continent": map[string]string{
				"iso_code": geo_information.Continent.Code,
				"name": geo_information.Continent.Names["en"],
			},
			"city": map[string]string{
				"name": geo_information.City.Names["en"],
			},
		},
		"traits": map[string]bool{
			"is_anonymous_proxy": geo_information.Traits.IsAnonymousProxy,
			"is_satellite_provider": geo_information.Traits.IsSatelliteProvider,
		},
	}

	json.NewEncoder(writer).Encode(ip_lookup)
}

func getGeoInformation(raw_ip string) *geoip2.City {
	db, err := geoip2.Open(GEOIP_DATABASE_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	parsed_ip := net.ParseIP(raw_ip)
	record, err := db.City(parsed_ip)
	if err != nil {
		log.Fatal(err)
	}

	return record;
}

func startCronjob()  {
	c := cron.New()
	c.AddFunc(*CRON_SCHEDULE, downloadGeoipDatabase)
	c.Start()
}

func isGeoipDatabasePresent() bool  {
	if _, err := os.Stat(GEOIP_DATABASE_FILE); os.IsNotExist(err) {
		return false
	}

	return true
}

func downloadGeoipDatabase() {
	events.LogStartedGeoipDatabaseDownload()
	exec.Command("/bin/sh", "downloadGeoipDatabase.sh").Run()
	events.LogFinishedGeoipDatabaseDownload()
}