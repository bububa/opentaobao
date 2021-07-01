package gameact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDeActivityMachineidGetAPIRequest
获取设备号 API请求
taobao.de.activity.machineid.get

获取机器设备id */
type TaobaoDeActivityMachineidGetAPIRequest struct {
	model.Params
}

// NewTaobaoDeActivityMachineidGetRequest 初始化TaobaoDeActivityMachineidGetAPIRequest对象
func NewTaobaoDeActivityMachineidGetRequest() *TaobaoDeActivityMachineidGetAPIRequest {
	return &TaobaoDeActivityMachineidGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDeActivityMachineidGetAPIRequest) GetApiMethodName() string {
	return "taobao.de.activity.machineid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDeActivityMachineidGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
