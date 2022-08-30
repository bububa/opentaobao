package cloudgame

// GameEndPoint 结构体
type GameEndPoint struct {
	// H5或者Native的访问点
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 后端的isv类型
	Provider string `json:"provider,omitempty" xml:"provider,omitempty"`
	// websocket 域名 H5
	WsServer string `json:"ws_server,omitempty" xml:"ws_server,omitempty"`
	// websocket 端口 H5
	WsPort string `json:"ws_port,omitempty" xml:"ws_port,omitempty"`
	// websocket token
	WsToken string `json:"ws_token,omitempty" xml:"ws_token,omitempty"`
	// isp
	Isp string `json:"isp,omitempty" xml:"isp,omitempty"`
	// region_id
	RegionId string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// h5_domain
	H5Domain string `json:"h5_domain,omitempty" xml:"h5_domain,omitempty"`
	// area_id
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
}
