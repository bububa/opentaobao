package ascp

// QueryScItemResponse 结构体
type QueryScItemResponse struct {
	// 货品列表
	ScitemModels []ScItemModel `json:"scitem_models,omitempty" xml:"scitem_models>sc_item_model,omitempty"`
	// 调用返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务是否成功，success为成功，否则为错误码
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 业务错误信息，成功则为空
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 业务处理结果
	Data *QueryScItemResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
