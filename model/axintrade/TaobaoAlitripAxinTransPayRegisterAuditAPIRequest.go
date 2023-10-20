package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintranspayregisterauditAPIRequest 阿信支付入驻审核通知 API请求
// taobao.alitrip.axin.trans.pay.register.audit
//
// 阿信支付入驻审核通知
type TaobaoalitripaxintranspayregisterauditAPIRequest struct {
	model.Params
	// 支付入驻审核对象
	_axinPayRegisterAuditDto *AxinPayRegisterAuditDto
}

// NewTaobaoalitripaxintranspayregisterauditRequest 初始化TaobaoalitripaxintranspayregisterauditAPIRequest对象
func NewTaobaoalitripaxintranspayregisterauditRequest() *TaobaoalitripaxintranspayregisterauditAPIRequest {
	return &TaobaoalitripaxintranspayregisterauditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripaxintranspayregisterauditAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.pay.register.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripaxintranspayregisterauditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripaxintranspayregisterauditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinPayRegisterAuditDto is AxinPayRegisterAuditDto Setter
// 支付入驻审核对象
func (r *TaobaoalitripaxintranspayregisterauditAPIRequest) SetAxinPayRegisterAuditDto(_axinPayRegisterAuditDto *AxinPayRegisterAuditDto) error {
	r._axinPayRegisterAuditDto = _axinPayRegisterAuditDto
	r.Set("axin_pay_register_audit_dto", _axinPayRegisterAuditDto)
	return nil
}

// GetAxinPayRegisterAuditDto AxinPayRegisterAuditDto Getter
func (r TaobaoalitripaxintranspayregisterauditAPIRequest) GetAxinPayRegisterAuditDto() *AxinPayRegisterAuditDto {
	return r._axinPayRegisterAuditDto
}
