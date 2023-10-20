package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeactivitydepositunbindAPIRequest 销售活动解绑预存金商品 API请求
// alibaba.alihouse.newhome.activity.deposit.unbind
//
// 销售活动解绑预存金商品
type AlibabaalihousenewhomeactivitydepositunbindAPIRequest struct {
	model.Params
	// 销售活动解绑预存金报文
	_preDepositGoldUnbindDto *PreDepositGoldUnbindDto
}

// NewAlibabaalihousenewhomeactivitydepositunbindRequest 初始化AlibabaalihousenewhomeactivitydepositunbindAPIRequest对象
func NewAlibabaalihousenewhomeactivitydepositunbindRequest() *AlibabaalihousenewhomeactivitydepositunbindAPIRequest {
	return &AlibabaalihousenewhomeactivitydepositunbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeactivitydepositunbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.deposit.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeactivitydepositunbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeactivitydepositunbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreDepositGoldUnbindDto is PreDepositGoldUnbindDto Setter
// 销售活动解绑预存金报文
func (r *AlibabaalihousenewhomeactivitydepositunbindAPIRequest) SetPreDepositGoldUnbindDto(_preDepositGoldUnbindDto *PreDepositGoldUnbindDto) error {
	r._preDepositGoldUnbindDto = _preDepositGoldUnbindDto
	r.Set("pre_deposit_gold_unbind_dto", _preDepositGoldUnbindDto)
	return nil
}

// GetPreDepositGoldUnbindDto PreDepositGoldUnbindDto Getter
func (r AlibabaalihousenewhomeactivitydepositunbindAPIRequest) GetPreDepositGoldUnbindDto() *PreDepositGoldUnbindDto {
	return r._preDepositGoldUnbindDto
}
