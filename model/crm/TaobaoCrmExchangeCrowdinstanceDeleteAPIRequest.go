package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest 删除人群实例中的指定买家 API请求
// taobao.crm.exchange.crowdinstance.delete
//
// 删除人群实例中的指定买家
type TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest struct {
	model.Params
	// 操作原因
	_reason string
	// 买家昵称
	_buyerNick string
	// 人群实例ID
	_crowdInstanceId int64
}

// NewTaobaoCrmExchangeCrowdinstanceDeleteRequest 初始化TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest对象
func NewTaobaoCrmExchangeCrowdinstanceDeleteRequest() *TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest {
	return &TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.crm.exchange.crowdinstance.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReason is Reason Setter
// 操作原因
func (r *TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest) GetReason() string {
	return r._reason
}

// SetBuyerNick is BuyerNick Setter
// 买家昵称
func (r *TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetCrowdInstanceId is CrowdInstanceId Setter
// 人群实例ID
func (r *TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest) SetCrowdInstanceId(_crowdInstanceId int64) error {
	r._crowdInstanceId = _crowdInstanceId
	r.Set("crowd_instance_id", _crowdInstanceId)
	return nil
}

// GetCrowdInstanceId CrowdInstanceId Getter
func (r TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest) GetCrowdInstanceId() int64 {
	return r._crowdInstanceId
}
