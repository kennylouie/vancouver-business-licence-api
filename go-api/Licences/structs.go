package Licences

type Licence struct {
  LicenceRSN  int `json:"LicenceRSN"`
  // LicenceNumber string `json:"LicenceNumber"`
  // LicenceRevisionNumber  int `json:"LicenceRevisionNumber"`
  BusinessName  string `json:"BusinessName"`
  // BusinessTradeName       string `json:"BusinessTradeName"`
  Status  string `json:"Status"`
  // IssuedDate      string `json:"IssuedDate"`
  // ExpiredDate      string `json:"ExpiredDate"`
  // BusinessType  string `json:"BusinessType"`
  // BusinessSubType      string `json:"BusinessSubType"`
  // Unit      string `json:"Unit"`
  // UnitType      string `json:"UnitType"`
  // House      string `json:"House"`
  // Street      string `json:"Street"`
  // City      string `json:"City"`
  // Province      string `json:"Province"`
  // Country      string `json:"Country"`
  // PostalCode      string `json:"PostalCode"`
  // LocalArea      string `json:"LocalArea"`
  // NumberOfEmployees      float64 `json:"NumberOfEmployees"`
  // Latitude      float64 `json:"Latitude"`
  // Longitude      float64 `json:"Longitude"`
  // ExtractDate      string `json:"ExtractDate"`
  Year  int `json:"Year"`
}

type GetLicenceResponse struct {
  Licence Licence `json:"licence"`
}

// ErrorResponse returns an array of error strings if appropriate
type ErrorResponse struct {
  Errors []string `json:"errors"`
}
