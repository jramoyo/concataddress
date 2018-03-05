package concataddress

// Input ...
type Input struct {
	Payload []struct {
		Address struct {
			BuildingNumber string
			Lat            float32
			Lon            float32
			Postcode       string
			State          string
			Street         string
			Suburb         string
		}
		PropertyTypeID int
		ReadyState     string
		Reference      string
		ShortID        string
		status         int
		Kind           string `json:"type"`
		Workflow       string
	}
}

// Output ...
type Output struct {
	Response []concataddress
}

type concataddress struct {
	Concataddress string
	Kind          string `json:"type"`
	Workflow      string
}
