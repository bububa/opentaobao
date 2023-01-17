package alihealthoutflow

// QuickAppRequest 结构体
type QuickAppRequest struct {
	// 场景id
	SceneId string `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	// 卡片当前展示的时间
	ShowTime string `json:"show_time,omitempty" xml:"show_time,omitempty"`
	// 渠道标识
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 应用版本
	AppVersion string `json:"app_version,omitempty" xml:"app_version,omitempty"`
	// 城市信息编码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
}
