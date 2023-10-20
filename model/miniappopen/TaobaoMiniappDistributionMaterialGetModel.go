package miniappopen

// TaobaominiappdistributionmaterialgetModel 结构体
type TaobaominiappdistributionmaterialgetModel struct {
	// 素材名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 场景卡片对应的动态表单内容
	DataStr string `json:"data_str,omitempty" xml:"data_str,omitempty"`
	// 素材对应的页面路径
	Path string `json:"path,omitempty" xml:"path,omitempty"`
	// 素材对应的页面参数值：全局参数值
	QParamsValueStr string `json:"q_params_value_str,omitempty" xml:"q_params_value_str,omitempty"`
	// 素材对应的页面参数值：页面参数值
	PParamsValueStr string `json:"p_params_value_str,omitempty" xml:"p_params_value_str,omitempty"`
	// 素材id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 小程序appid
	AppId int64 `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 支持的投放入口卡片id
	CardId int64 `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 支持的投放入口场景id
	SceneId int64 `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
}
