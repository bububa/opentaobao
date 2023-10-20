package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMsAreaProvinceListAPIRequest 疫苗预约省份列表查询 API请求
// alibaba.alihealth.ms.area.province.list
//
// 微信小程序疫苗预约省份列表查询
type AlibabaAlihealthMsAreaProvinceListAPIRequest struct {
	model.Params
}

// NewAlibabaAlihealthMsAreaProvinceListRequest 初始化AlibabaAlihealthMsAreaProvinceListAPIRequest对象
func NewAlibabaAlihealthMsAreaProvinceListRequest() *AlibabaAlihealthMsAreaProvinceListAPIRequest {
	return &AlibabaAlihealthMsAreaProvinceListAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMsAreaProvinceListAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMsAreaProvinceListAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.ms.area.province.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMsAreaProvinceListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMsAreaProvinceListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaAlihealthMsAreaProvinceListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMsAreaProvinceListRequest()
	},
}

// GetAlibabaAlihealthMsAreaProvinceListRequest 从 sync.Pool 获取 AlibabaAlihealthMsAreaProvinceListAPIRequest
func GetAlibabaAlihealthMsAreaProvinceListAPIRequest() *AlibabaAlihealthMsAreaProvinceListAPIRequest {
	return poolAlibabaAlihealthMsAreaProvinceListAPIRequest.Get().(*AlibabaAlihealthMsAreaProvinceListAPIRequest)
}

// ReleaseAlibabaAlihealthMsAreaProvinceListAPIRequest 将 AlibabaAlihealthMsAreaProvinceListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMsAreaProvinceListAPIRequest(v *AlibabaAlihealthMsAreaProvinceListAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMsAreaProvinceListAPIRequest.Put(v)
}
