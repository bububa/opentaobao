package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaodianserverdategetAPIRequest 服务器时间获取 API请求
// taobao.baodian.server.date.get
//
// 获取服务器时间
type TaobaobaodianserverdategetAPIRequest struct {
	model.Params
}

// NewTaobaobaodianserverdategetRequest 初始化TaobaobaodianserverdategetAPIRequest对象
func NewTaobaobaodianserverdategetRequest() *TaobaobaodianserverdategetAPIRequest {
	return &TaobaobaodianserverdategetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaodianserverdategetAPIRequest) GetApiMethodName() string {
	return "taobao.baodian.server.date.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaodianserverdategetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaodianserverdategetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
