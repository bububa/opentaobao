package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿信支付入驻审核通知 API请求
taobao.alitrip.axin.trans.pay.register.audit

阿信支付入驻审核通知
*/
type TaobaoAlitripAxinTransPayRegisterAuditRequest struct {
    model.Params
    // 支付入驻审核对象
    _axinPayRegisterAuditDto   *AxinPayRegisterAuditDto
}

// 初始化TaobaoAlitripAxinTransPayRegisterAuditRequest对象
func NewTaobaoAlitripAxinTransPayRegisterAuditRequest() *TaobaoAlitripAxinTransPayRegisterAuditRequest{
    return &TaobaoAlitripAxinTransPayRegisterAuditRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPayRegisterAuditRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.pay.register.audit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPayRegisterAuditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AxinPayRegisterAuditDto Setter
// 支付入驻审核对象
func (r *TaobaoAlitripAxinTransPayRegisterAuditRequest) SetAxinPayRegisterAuditDto(_axinPayRegisterAuditDto *AxinPayRegisterAuditDto) error {
    r._axinPayRegisterAuditDto = _axinPayRegisterAuditDto
    r.Set("axin_pay_register_audit_dto", _axinPayRegisterAuditDto)
    return nil
}

// AxinPayRegisterAuditDto Getter
func (r TaobaoAlitripAxinTransPayRegisterAuditRequest) GetAxinPayRegisterAuditDto() *AxinPayRegisterAuditDto {
    return r._axinPayRegisterAuditDto
}
