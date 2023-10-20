package alihealthmedical

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalOrderQueryAPIRequest 三方机构查询订单详情接口 API请求
// alibaba.alihealth.medical.order.query
//
// 查询订单详情，包括评价
type AlibabaAlihealthMedicalOrderQueryAPIRequest struct {
	model.Params
	// 请求入参
	_requestInfo *OrderQueryRequestDto
}

// NewAlibabaAlihealthMedicalOrderQueryRequest 初始化AlibabaAlihealthMedicalOrderQueryAPIRequest对象
func NewAlibabaAlihealthMedicalOrderQueryRequest() *AlibabaAlihealthMedicalOrderQueryAPIRequest {
	return &AlibabaAlihealthMedicalOrderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalOrderQueryAPIRequest) Reset() {
	r._requestInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestInfo is RequestInfo Setter
// 请求入参
func (r *AlibabaAlihealthMedicalOrderQueryAPIRequest) SetRequestInfo(_requestInfo *OrderQueryRequestDto) error {
	r._requestInfo = _requestInfo
	r.Set("request_info", _requestInfo)
	return nil
}

// GetRequestInfo RequestInfo Getter
func (r AlibabaAlihealthMedicalOrderQueryAPIRequest) GetRequestInfo() *OrderQueryRequestDto {
	return r._requestInfo
}

var poolAlibabaAlihealthMedicalOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalOrderQueryRequest()
	},
}

// GetAlibabaAlihealthMedicalOrderQueryRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalOrderQueryAPIRequest
func GetAlibabaAlihealthMedicalOrderQueryAPIRequest() *AlibabaAlihealthMedicalOrderQueryAPIRequest {
	return poolAlibabaAlihealthMedicalOrderQueryAPIRequest.Get().(*AlibabaAlihealthMedicalOrderQueryAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalOrderQueryAPIRequest 将 AlibabaAlihealthMedicalOrderQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalOrderQueryAPIRequest(v *AlibabaAlihealthMedicalOrderQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalOrderQueryAPIRequest.Put(v)
}
