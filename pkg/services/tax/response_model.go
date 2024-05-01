package tax

type ServerResponseLayout struct {
	Signature      string `json:"signature,omitempty"`
	SignatureKeyId string `json:"signatureKeyId,omitempty"`
	Timestamp      int    `json:"timestamp,omitempty"`
}

type ServerInfoResult struct {
	ServerResponseLayout
	Result struct {
		Data struct {
			PublicKeys []struct {
				Algorithm string `json:"algorithm"`
				ID        string `json:"id"`
				Key       string `json:"key"`
				Purpose   int    `json:"purpose"`
			} `json:"publicKeys"`
		} `json:"data"`
		ServerTime int `json:"serverTime"`
	} `json:"result"`
}
