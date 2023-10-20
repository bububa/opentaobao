package gameact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodeactivitymachineidgetAPIRequest 获取设备号 API请求
// taobao.de.activity.machineid.get
//
// 获取机器设备id
type TaobaodeactivitymachineidgetAPIRequest struct {
	model.Params
}

// NewTaobaodeactivitymachineidgetRequest 初始化TaobaodeactivitymachineidgetAPIRequest对象
func NewTaobaodeactivitymachineidgetRequest() *TaobaodeactivitymachineidgetAPIRequest {
	return &TaobaodeactivitymachineidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodeactivitymachineidgetAPIRequest) GetApiMethodName() string {
	return "taobao.de.activity.machineid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodeactivitymachineidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodeactivitymachineidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
