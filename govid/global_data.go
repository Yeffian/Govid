package govid

// Stores global covid data
type GlobalData struct {
	Cases     int `json:"cases"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
}