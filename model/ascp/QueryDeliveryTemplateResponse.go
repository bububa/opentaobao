package ascp

// QueryDeliveryTemplateResponse 结构体
type QueryDeliveryTemplateResponse struct {
	// 运费模板列表
	DeliveryTemplatemodelList []DeliveryTemplateModel `json:"delivery_templatemodel_list,omitempty" xml:"delivery_templatemodel_list>delivery_template_model,omitempty"`
	// 查询结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 错误原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
