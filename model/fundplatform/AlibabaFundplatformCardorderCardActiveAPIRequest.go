package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardorderCardActiveAPIRequest 储值卡激活 API请求
// alibaba.fundplatform.cardorder.card.active
//
// 储值卡接货接口，可以通过外部订单号或者卡号进行批量激活。如果储值卡已经被激活过仍然幂等返回成功。资金平台保证批量激活时一定全部成功或全部失败。
// 如果批量激活储值卡时，如果部分储值卡处于已激活，部分储值卡处于未激活，则会返回失败
type AlibabaFundplatformCardorderCardActiveAPIRequest struct {
	model.Params
	// 入参对象
	_paramCardActiveRequest *CardActiveRequest
}

// NewAlibabaFundplatformCardorderCardActiveRequest 初始化AlibabaFundplatformCardorderCardActiveAPIRequest对象
func NewAlibabaFundplatformCardorderCardActiveRequest() *AlibabaFundplatformCardorderCardActiveAPIRequest {
	return &AlibabaFundplatformCardorderCardActiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderCardActiveAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorder.card.active"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderCardActiveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamCardActiveRequest Setter
// 入参对象
func (r *AlibabaFundplatformCardorderCardActiveAPIRequest) SetParamCardActiveRequest(_paramCardActiveRequest *CardActiveRequest) error {
	r._paramCardActiveRequest = _paramCardActiveRequest
	r.Set("param_card_active_request", _paramCardActiveRequest)
	return nil
}

// Get ParamCardActiveRequest Getter
func (r AlibabaFundplatformCardorderCardActiveAPIRequest) GetParamCardActiveRequest() *CardActiveRequest {
	return r._paramCardActiveRequest
}
