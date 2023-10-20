package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentshoppingpushAPIRequest 国际机票大卖家Shopping推送 API请求
// taobao.alitrip.ie.agent.shopping.push
//
// 用于国际机票大卖家主动推送Shopping结果更新缓存报价。
type TaobaoalitripieagentshoppingpushAPIRequest struct {
	model.Params
	// 政策推送结构体
	_param0 *ShoppingPushRq
}

// NewTaobaoalitripieagentshoppingpushRequest 初始化TaobaoalitripieagentshoppingpushAPIRequest对象
func NewTaobaoalitripieagentshoppingpushRequest() *TaobaoalitripieagentshoppingpushAPIRequest {
	return &TaobaoalitripieagentshoppingpushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripieagentshoppingpushAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.shopping.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripieagentshoppingpushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripieagentshoppingpushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 政策推送结构体
func (r *TaobaoalitripieagentshoppingpushAPIRequest) SetParam0(_param0 *ShoppingPushRq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoalitripieagentshoppingpushAPIRequest) GetParam0() *ShoppingPushRq {
	return r._param0
}
