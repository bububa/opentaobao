package idle

// UserGrantRequest 结构体
type UserGrantRequest struct {
	// 当前用户的所属业务类型编码，优品&开放平台业务 默认使用 IDLE_TOP
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 场景码，标识品类。22:虚拟卡券/账号
	SceneType string `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
}
