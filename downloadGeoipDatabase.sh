#!/bin/bash

set -e

geoip_database_location=http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.mmdb.gz
destination=geoip_database

wget ${geoip_database_location} -O ${destination}.gz
gzip --decompress --stdout ${destination}.gz > ${destination}.mmdb
rm ${destination}.gz