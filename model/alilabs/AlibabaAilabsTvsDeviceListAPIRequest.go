package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstvsdevicelistAPIRequest 获取TVS设备列表 API请求
// alibaba.ailabs.tvs.device.list
//
// 获取用户所绑定的TVS设备列表
type AlibabaailabstvsdevicelistAPIRequest struct {
	model.Params
}

// NewAlibabaailabstvsdevicelistRequest 初始化AlibabaailabstvsdevicelistAPIRequest对象
func NewAlibabaailabstvsdevicelistRequest() *AlibabaailabstvsdevicelistAPIRequest {
	return &AlibabaailabstvsdevicelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstvsdevicelistAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tvs.device.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstvsdevicelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstvsdevicelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}
