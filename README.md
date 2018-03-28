# vancouver-business-licence-api

## Background

This repo will contain a codebase for the continuing of our winning prototype from the 2018 Vancouver Open Data Day Hackathon where we tackled housing affordability (https://www.opendatabc.ca/pages/vodday-2018-vancouver-open-data-day-hackathon). We prototyped an API predicted geolocations across the Greater Vancouver Regional District that had similar 'surroundings'. Surroundings we defined as having similar infrastructure, e.g. building types, stores nearby, number of hospitals, number of schools, cost of rent, etc.. We felt this would be valuable to those struggling to find a suitable location to rent that would be similar to where they are currently renting. With increasing rent prices, many Vancouverites are forced to move our of their current rental unit to find cheaper rent. This tool might be useful.

### Concept

1. Build a Cassandra backend to store the 1997-2018 business licence dataset.
2. Build a Go API to access the Cassandra database.

## Progress

1. Setup an AWS EC2 to host our data and make it accessible.

```
aws ec2 run-instances --image-id ami-79873901 --count 1 \
--instance-type t2.micro --key-name *your-key-name* --security-groups *your-security-groups* \
--user-data file://user-data.sh
```

2. Create a simple cassandra keystore holding our data. Can be found in the cassandra-setup.sh file. The data is found in the data folder. I have added the Year column to all the files. Currently have some 2018 and 2017 data in the database.

3. Build a simple golang server to relay http requests to access the cassandra database. This was adapted from Ian Douglas (https://getstream.io/blog/building-a-performant-api-using-go-and-cassandra/).

The server is live at
```
http://34.218.47.133:8081/
```
