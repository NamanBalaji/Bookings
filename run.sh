#!/bin/bash

go build -o bookings.exe cmd/web/*.go && ./bookings -dbname=bookings_database -dbuser=postgres -cache=false -production=false