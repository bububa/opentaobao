package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentShoppingPushAPIRequest
国际机票大卖家Shopping推送 API请求
taobao.alitrip.ie.agent.shopping.push

用于国际机票大卖家主动推送Shopping结果更新缓存报价。 */
type TaobaoAlitripIeAgentShoppingPushAPIRequest struct {
	model.Params
	// 政策推送结构体
	_param0 *ShoppingPushRq
}

// NewTaobaoAlitripIeAgentShoppingPushRequest 初始化TaobaoAlitripIeAgentShoppingPushAPIRequest对象
func NewTaobaoAlitripIeAgentShoppingPushRequest() *TaobaoAlitripIeAgentShoppingPushAPIRequest {
	return &TaobaoAlitripIeAgentShoppingPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentShoppingPushAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.shopping.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentShoppingPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 政策推送结构体
func (r *TaobaoAlitripIeAgentShoppingPushAPIRequest) SetParam0(_param0 *ShoppingPushRq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoAlitripIeAgentShoppingPushAPIRequest) GetParam0() *ShoppingPushRq {
	return r._param0
}
