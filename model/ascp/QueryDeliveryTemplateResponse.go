package ascp

import (
	"sync"
)

// QueryDeliveryTemplateResponse 结构体
type QueryDeliveryTemplateResponse struct {
	// 运费模板列表
	DeliveryTemplatemodelList []DeliveryTemplateModel `json:"delivery_templatemodel_list,omitempty" xml:"delivery_templatemodel_list>delivery_template_model,omitempty"`
	// 查询结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 错误原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolQueryDeliveryTemplateResponse = sync.Pool{
	New: func() any {
		return new(QueryDeliveryTemplateResponse)
	},
}

// GetQueryDeliveryTemplateResponse() 从对象池中获取QueryDeliveryTemplateResponse
func GetQueryDeliveryTemplateResponse() *QueryDeliveryTemplateResponse {
	return poolQueryDeliveryTemplateResponse.Get().(*QueryDeliveryTemplateResponse)
}

// ReleaseQueryDeliveryTemplateResponse 释放QueryDeliveryTemplateResponse
func ReleaseQueryDeliveryTemplateResponse(v *QueryDeliveryTemplateResponse) {
	v.DeliveryTemplatemodelList = v.DeliveryTemplatemodelList[:0]
	v.Result = ""
	v.Message = ""
	poolQueryDeliveryTemplateResponse.Put(v)
}
