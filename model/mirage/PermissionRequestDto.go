package mirage

// PermissionRequestDto 结构体
type PermissionRequestDto struct {
	// 资源id
	ResourceIds []string `json:"resource_ids,omitempty" xml:"resource_ids>string,omitempty"`
	// 播放场景
	DisplayScene string `json:"display_scene,omitempty" xml:"display_scene,omitempty"`
	// pid
	Pid string `json:"pid,omitempty" xml:"pid,omitempty"`
	// 端对drm的能力
	ClientDrmAbility string `json:"client_drm_ability,omitempty" xml:"client_drm_ability,omitempty"`
	// app 版本
	AppVersion string `json:"app_version,omitempty" xml:"app_version,omitempty"`
	// 用户身份
	UserIdentity string `json:"user_identity,omitempty" xml:"user_identity,omitempty"`
	// ccode
	Ccode string `json:"ccode,omitempty" xml:"ccode,omitempty"`
	// ytid
	Ytid string `json:"ytid,omitempty" xml:"ytid,omitempty"`
	// 用户ip
	UserIp string `json:"user_ip,omitempty" xml:"user_ip,omitempty"`
	// 用户简化版userAgent
	Ua string `json:"ua,omitempty" xml:"ua,omitempty"`
	// 用户完整userAgent
	UserAgent string `json:"user_agent,omitempty" xml:"user_agent,omitempty"`
	// 设备类型
	DeviceType string `json:"device_type,omitempty" xml:"device_type,omitempty"`
	// 运营商
	DmaCode string `json:"dma_code,omitempty" xml:"dma_code,omitempty"`
	// 地区码
	AreaCode string `json:"area_code,omitempty" xml:"area_code,omitempty"`
	// 国家码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 网站
	Site string `json:"site,omitempty" xml:"site,omitempty"`
	// 资源类型
	ResourceType string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// 调用方
	Caller string `json:"caller,omitempty" xml:"caller,omitempty"`
	// 签名
	Signature string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 需要展示形式配置
	NeedDisplayConfig bool `json:"need_display_config,omitempty" xml:"need_display_config,omitempty"`
	// 需要同步返回drm配置信息
	NeedDrmConfig bool `json:"need_drm_config,omitempty" xml:"need_drm_config,omitempty"`
	// 苹果
	IsFromApple bool `json:"is_from_apple,omitempty" xml:"is_from_apple,omitempty"`
}
