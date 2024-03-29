package ascp

import (
	"sync"
)

// QueryScItemResponse 结构体
type QueryScItemResponse struct {
	// 货品列表
	ScitemModels []ScItemModel `json:"scitem_models,omitempty" xml:"scitem_models>sc_item_model,omitempty"`
	// 业务是否成功，success为成功，否则为错误码
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 业务错误信息，成功则为空
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
}

var poolQueryScItemResponse = sync.Pool{
	New: func() any {
		return new(QueryScItemResponse)
	},
}

// GetQueryScItemResponse() 从对象池中获取QueryScItemResponse
func GetQueryScItemResponse() *QueryScItemResponse {
	return poolQueryScItemResponse.Get().(*QueryScItemResponse)
}

// ReleaseQueryScItemResponse 释放QueryScItemResponse
func ReleaseQueryScItemResponse(v *QueryScItemResponse) {
	v.ScitemModels = v.ScitemModels[:0]
	v.Result = ""
	v.ErrorMessage = ""
	poolQueryScItemResponse.Put(v)
}
