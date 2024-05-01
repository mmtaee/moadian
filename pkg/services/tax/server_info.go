package tax

func ServerInfo(TaxUrl string) (ServerInfoResult, int, error) {
	p := Packet{
		PacketType: "GET_SERVER_INFORMATION",
	}
	packetMap := p.ToMap()
	request := Request{
		Path:   "GET_SERVER_INFORMATION",
		Method: "post",
		Sync:   true,
	}
	request.MakeHeader("")
	request.MakeUrl(TaxUrl)
	request.MakeBody(packetMap, "", "", 1)
	result := new(ServerInfoResult)
	send, status, err := request.Send(result)
	return send.(ServerInfoResult), status, err
}
