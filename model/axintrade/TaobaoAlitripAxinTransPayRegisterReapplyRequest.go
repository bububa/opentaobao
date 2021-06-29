package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿信支付入驻重新申请 API请求
taobao.alitrip.axin.trans.pay.register.reapply

阿信支付入驻重新申请
用于支付平台驳回，商户提交时的业务场景
*/
type TaobaoAlitripAxinTransPayRegisterReapplyRequest struct {
    model.Params
    // 阿信支付入驻重新申请参数
    _axinPayRegisterCreateDTO   *AxinPayRegisterCreateDTO
}

// 初始化TaobaoAlitripAxinTransPayRegisterReapplyRequest对象
func NewTaobaoAlitripAxinTransPayRegisterReapplyRequest() *TaobaoAlitripAxinTransPayRegisterReapplyRequest{
    return &TaobaoAlitripAxinTransPayRegisterReapplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPayRegisterReapplyRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.pay.register.reapply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPayRegisterReapplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AxinPayRegisterCreateDTO Setter
// 阿信支付入驻重新申请参数
func (r *TaobaoAlitripAxinTransPayRegisterReapplyRequest) SetAxinPayRegisterCreateDTO(_axinPayRegisterCreateDTO *AxinPayRegisterCreateDTO) error {
    r._axinPayRegisterCreateDTO = _axinPayRegisterCreateDTO
    r.Set("axin_pay_register_create_d_t_o", _axinPayRegisterCreateDTO)
    return nil
}

// AxinPayRegisterCreateDTO Getter
func (r TaobaoAlitripAxinTransPayRegisterReapplyRequest) GetAxinPayRegisterCreateDTO() *AxinPayRegisterCreateDTO {
    return r._axinPayRegisterCreateDTO
}
