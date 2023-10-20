package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeactivitydepositbindAPIRequest 销售活动绑定预存金商品id API请求
// alibaba.alihouse.newhome.activity.deposit.bind
//
// 销售活动绑定预存金商品id
type AlibabaalihousenewhomeactivitydepositbindAPIRequest struct {
	model.Params
	// 绑定预存金商品属性
	_preDepositGoldDto *PreDepositGoldDto
}

// NewAlibabaalihousenewhomeactivitydepositbindRequest 初始化AlibabaalihousenewhomeactivitydepositbindAPIRequest对象
func NewAlibabaalihousenewhomeactivitydepositbindRequest() *AlibabaalihousenewhomeactivitydepositbindAPIRequest {
	return &AlibabaalihousenewhomeactivitydepositbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeactivitydepositbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.deposit.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeactivitydepositbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeactivitydepositbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreDepositGoldDto is PreDepositGoldDto Setter
// 绑定预存金商品属性
func (r *AlibabaalihousenewhomeactivitydepositbindAPIRequest) SetPreDepositGoldDto(_preDepositGoldDto *PreDepositGoldDto) error {
	r._preDepositGoldDto = _preDepositGoldDto
	r.Set("pre_deposit_gold_dto", _preDepositGoldDto)
	return nil
}

// GetPreDepositGoldDto PreDepositGoldDto Getter
func (r AlibabaalihousenewhomeactivitydepositbindAPIRequest) GetPreDepositGoldDto() *PreDepositGoldDto {
	return r._preDepositGoldDto
}
