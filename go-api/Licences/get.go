package Licences

import (
  "net/http"
  "github.com/gocql/gocql"
  "encoding/json"
  "go-api/Cassandra"
  "github.com/gorilla/mux"
)

func GetLicence(w http.ResponseWriter, r *http.Request) {
  var licence Licence
  var errs []string
  var found bool = false

  vars := mux.Vars(r)
  rsn := vars["licence_LicenceRSN"]

  m := map[string]interface{}{}
  // query := "SELECT LicenceRSN,LicenceNumber,LicenceRevisionNumber,BusinessName,BusinessTradeName,Status,IssuedDate,ExpiredDate,BusinessType,BusinessSubType,Unit,UnitType,House,Street,City,Province,Country,PostalCode,LocalArea,NumberOfEmployees,Latitude,Longitude,ExtractDate,Year FROM vancouver_opendata.business_licence WHERE LicenceRSN=? AND Year=2001 LIMIT 1"
  //
  query := "SELECT LicenceRSN,BusinessName,Status,Year FROM business_licence WHERE LicenceRSN=? AND Year=2018 LIMIT 1"
  iterable := Cassandra.Session.Query(query, rsn).Consistency(gocql.One).Iter()
  for iterable.MapScan(m) {
    found = true
    licence = Licence {
      LicenceRSN: m["licencersn"].(int),
      // LicenceNumber: m["licencenumber"].(string),
      // LicenceRevisionNumber: m["licencerevisionNumber"].(int),
      BusinessName: m["businessname"].(string),
      // BusinessTradeName: m["businesstradeName"].(string),
      Status: m["status"].(string),
      // IssuedDate: m["issueddate"].(string),
      // ExpiredDate: m["expireddate"].(string),
      // BusinessType: m["businesstype"].(string),
      // BusinessSubType: m["businesssubtype"].(string),
      // Unit: m["unit"].(string),
      // UnitType: m["unittype"].(string),
      // House: m["house"].(string),
      // Street: m["street"].(string),
      // City: m["city"].(string),
      // Province: m["province"].(string),
      // Country: m["country"].(string),
      // PostalCode: m["postalcode"].(string),
      // LocalArea: m["localarea"].(string),
      // NumberOfEmployees: m["numberofemployees"].(float64),
      // Latitude: m["latitude"].(float64),
      // Longitude: m["longitude"].(float64),
      // ExtractDate: m["extractdate"].(string),
      Year: m["year"].(int),
    }
  }

  if !found {
    errs = append(errs, "Licence not found")
  }

  if found {
    json.NewEncoder(w).Encode(GetLicenceResponse{Licence: licence})
  } else {
    json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
  }
}
