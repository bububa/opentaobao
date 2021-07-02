package alilabs

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTvsDeviceListAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tvs.device.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTvsDeviceListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
