package events

import (
	log "github.com/Sirupsen/logrus"
)

const startedGeoipDatabaseDownload = "started_geoip_database_download"
const finishedGeoipDatabaseDownload = "finished_geoip_database_download"
const startedIplookup = "started_ip_lookup"
const failedIplookup = "failed_ip_lookup"
const started_server = "started_server"

func LogStartedGeoipDatabaseDownload()  {
	logEvent(startedGeoipDatabaseDownload).Info()
}

func LogFinishedGeoipDatabaseDownload()  {
	logEvent(finishedGeoipDatabaseDownload).Info()
}

func LogStartedIpLookup(ip string)  {
	logEvent(startedIplookup).WithFields(log.Fields{
		"ip": ip,
	}).Info()
}

func LogFailedIpLookup(error error, ip string)  {
	logEvent(failedIplookup).WithFields(log.Fields{
		"ip": ip,
	}).Error(error)
}

func LogStartedServer(port int)  {
	logEvent(started_server).WithFields(log.Fields{
		"port": port,
	}).Info()
}

func logEvent(event string) *log.Entry {
	return log.WithField("event", event)
}