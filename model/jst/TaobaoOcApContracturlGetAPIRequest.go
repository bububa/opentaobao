package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcApContracturlGetAPIRequest 按用户获取支付宝代扣协议链接地址 API请求
// taobao.oc.ap.contracturl.get
//
// 按用户获取支付宝代扣协议链接地址
type TaobaoOcApContracturlGetAPIRequest struct {
	model.Params
}

// NewTaobaoOcApContracturlGetRequest 初始化TaobaoOcApContracturlGetAPIRequest对象
func NewTaobaoOcApContracturlGetRequest() *TaobaoOcApContracturlGetAPIRequest {
	return &TaobaoOcApContracturlGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOcApContracturlGetAPIRequest) GetApiMethodName() string {
	return "taobao.oc.ap.contracturl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOcApContracturlGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
