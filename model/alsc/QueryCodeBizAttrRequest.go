package alsc

// QueryCodeBizAttrRequest 结构体
type QueryCodeBizAttrRequest struct {
	// 码使用的业务场景
	BizScenePrefix string `json:"biz_scene_prefix,omitempty" xml:"biz_scene_prefix,omitempty"`
	// 码值
	CodeValue string `json:"code_value,omitempty" xml:"code_value,omitempty"`
}
