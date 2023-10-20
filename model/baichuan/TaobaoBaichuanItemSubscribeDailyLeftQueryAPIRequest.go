package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanitemsubscribedailyleftqueryAPIRequest 查询当天可添加的余量 API请求
// taobao.baichuan.item.subscribe.daily.left.query
//
// 查询当天可添加的余量
type TaobaobaichuanitemsubscribedailyleftqueryAPIRequest struct {
	model.Params
}

// NewTaobaobaichuanitemsubscribedailyleftqueryRequest 初始化TaobaobaichuanitemsubscribedailyleftqueryAPIRequest对象
func NewTaobaobaichuanitemsubscribedailyleftqueryRequest() *TaobaobaichuanitemsubscribedailyleftqueryAPIRequest {
	return &TaobaobaichuanitemsubscribedailyleftqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanitemsubscribedailyleftqueryAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.item.subscribe.daily.left.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanitemsubscribedailyleftqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanitemsubscribedailyleftqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
