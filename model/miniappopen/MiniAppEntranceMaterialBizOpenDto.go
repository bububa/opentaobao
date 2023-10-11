package miniappopen

// MiniAppEntranceMaterialBizOpenDto 结构体
type MiniAppEntranceMaterialBizOpenDto struct {
	// 素材id列表
	MaterialIdList []int64 `json:"material_id_list,omitempty" xml:"material_id_list>int64,omitempty"`
	// 场景入口对应的素材动态表单值，必须符合表单的要求。每个场景、入口卡片都拥有不同的限制，请通过开发者控制台查看具体的表单字段要求。如果没有任何参数值，请传入空map
	DataStr string `json:"data_str,omitempty" xml:"data_str,omitempty"`
	// 对应页面的页面参数值map，如果没有任何参数值，请传入空map
	PParamsValueStr string `json:"p_params_value_str,omitempty" xml:"p_params_value_str,omitempty"`
	// 当前素材，对应的页面路径。必须为已经在开发者控制台-页面录入菜单下，已经录入过的页面路径之一。
	Path string `json:"path,omitempty" xml:"path,omitempty"`
	// 对应页面的全局参数值map，如果没有任何参数值，请传入空map
	QParamsValueStr string `json:"q_params_value_str,omitempty" xml:"q_params_value_str,omitempty"`
	// 素材的名字，主要用于方便商家辨识素材
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 对应的可投放的小程序id
	AppId int64 `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 素材适用的入口样式卡片id
	CardId int64 `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 素材适用的入口场景id
	SceneId int64 `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	// 素材id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
	// 要修改的素材id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
