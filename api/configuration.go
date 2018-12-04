package api

type Configuration struct {
	Admin *User `json:"Admin"`
	Mail  *Mail `json:"Mail"`
}

type Mail struct {
	Host     string `json:"Host"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}
