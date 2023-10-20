package alilabs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTvsDeviceListAPIRequest 获取TVS设备列表 API请求
// alibaba.ailabs.tvs.device.list
//
// 获取用户所绑定的TVS设备列表
type AlibabaAilabsTvsDeviceListAPIRequest struct {
	model.Params
}

// NewAlibabaAilabsTvsDeviceListRequest 初始化AlibabaAilabsTvsDeviceListAPIRequest对象
func NewAlibabaAilabsTvsDeviceListRequest() *AlibabaAilabsTvsDeviceListAPIRequest {
	return &AlibabaAilabsTvsDeviceListAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsTvsDeviceListAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTvsDeviceListAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tvs.device.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTvsDeviceListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTvsDeviceListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaAilabsTvsDeviceListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsTvsDeviceListRequest()
	},
}

// GetAlibabaAilabsTvsDeviceListRequest 从 sync.Pool 获取 AlibabaAilabsTvsDeviceListAPIRequest
func GetAlibabaAilabsTvsDeviceListAPIRequest() *AlibabaAilabsTvsDeviceListAPIRequest {
	return poolAlibabaAilabsTvsDeviceListAPIRequest.Get().(*AlibabaAilabsTvsDeviceListAPIRequest)
}

// ReleaseAlibabaAilabsTvsDeviceListAPIRequest 将 AlibabaAilabsTvsDeviceListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsTvsDeviceListAPIRequest(v *AlibabaAilabsTvsDeviceListAPIRequest) {
	v.Reset()
	poolAlibabaAilabsTvsDeviceListAPIRequest.Put(v)
}
