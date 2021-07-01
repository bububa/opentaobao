package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransPayRegisterAuditAPIRequest
阿信支付入驻审核通知 API请求
taobao.alitrip.axin.trans.pay.register.audit

阿信支付入驻审核通知 */
type TaobaoAlitripAxinTransPayRegisterAuditAPIRequest struct {
	model.Params
	// 支付入驻审核对象
	_axinPayRegisterAuditDto *AxinPayRegisterAuditDto
}

// NewTaobaoAlitripAxinTransPayRegisterAuditRequest 初始化TaobaoAlitripAxinTransPayRegisterAuditAPIRequest对象
func NewTaobaoAlitripAxinTransPayRegisterAuditRequest() *TaobaoAlitripAxinTransPayRegisterAuditAPIRequest {
	return &TaobaoAlitripAxinTransPayRegisterAuditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPayRegisterAuditAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.pay.register.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPayRegisterAuditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AxinPayRegisterAuditDto Setter
// 支付入驻审核对象
func (r *TaobaoAlitripAxinTransPayRegisterAuditAPIRequest) SetAxinPayRegisterAuditDto(_axinPayRegisterAuditDto *AxinPayRegisterAuditDto) error {
	r._axinPayRegisterAuditDto = _axinPayRegisterAuditDto
	r.Set("axin_pay_register_audit_dto", _axinPayRegisterAuditDto)
	return nil
}

// Get AxinPayRegisterAuditDto Getter
func (r TaobaoAlitripAxinTransPayRegisterAuditAPIRequest) GetAxinPayRegisterAuditDto() *AxinPayRegisterAuditDto {
	return r._axinPayRegisterAuditDto
}
