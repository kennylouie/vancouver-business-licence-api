#!/bin/bash

## create keystore
echo "CREATE KEYSPACE IF NOT EXISTS vancouver_opendata WITH REPLICATION =
{ 'class' : 'SimpleStrategy', 'replication_factor' : 1};" | cqlsh


## create table to match business licence data
echo "CREATE TABLE vancouver_opendata.business_licence (
  LicenceRSN int,
  LicenceNumber text,
  LicenceRevisionNumber int,
  BusinessName text,
  BusinessTradeName text,
  Status text,
  IssuedDate text,
  ExpiredDate text,
  BusinessType text,
  BusinessSubType text,
  Unit text,
  UnitType text,
  House text,
  Street text,
  City text,
  Province text,
  Country text,
  PostalCode text,
  LocalArea text,
  NumberOfEmployees float,
  Latitude float,
  Longitude float,
  FeePaid float,
  ExtractDate text,
  Year int,
  PRIMARY KEY ((LicenceRSN, Year), BusinessTradeName)
);" | cqlsh

## injest csv files into the table
echo "COPY vancouver_opendata.business_licence (
  LicenceRSN,
  LicenceNumber,
  LicenceRevisionNumber,
  BusinessName,
  BusinessTradeName,
  Status,
  IssuedDate,
  ExpiredDate,
  BusinessType,
  BusinessSubType,
  Unit,
  UnitType,
  House,
  Street,
  City,
  Province,
  Country,
  PostalCode,
  LocalArea,
  NumberOfEmployees,
  Latitude,
  Longitude,
  FeePaid,
  ExtractDate,
  Year
  ) FROM './2001business_licences.csv' WITH HEADER=TRUE AND NULL='null';
" | cqlsh
