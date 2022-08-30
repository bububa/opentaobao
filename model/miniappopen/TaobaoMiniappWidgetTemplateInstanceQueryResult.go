package miniappopen

// TaobaoMiniappWidgetTemplateInstanceQueryResult 结构体
type TaobaoMiniappWidgetTemplateInstanceQueryResult struct {
	// 返回结果
	Model []MiniappInstanceAppAllVersionsDto `json:"model,omitempty" xml:"model>miniapp_instance_app_all_versions_dto,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}
