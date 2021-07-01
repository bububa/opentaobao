package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿信创建退款单 API请求
taobao.alitrip.axin.trans.refund.create

阿信供销平台-创建退款单服务
*/
type TaobaoAlitripAxinTransRefundCreateAPIRequest struct {
    model.Params
    // 阿信创建退款单入参
    _axinRefundCreateDTO   *AxinRefundCreateDTO
}

// 初始化TaobaoAlitripAxinTransRefundCreateAPIRequest对象
func NewTaobaoAlitripAxinTransRefundCreateRequest() *TaobaoAlitripAxinTransRefundCreateAPIRequest{
    return &TaobaoAlitripAxinTransRefundCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransRefundCreateAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.refund.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransRefundCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AxinRefundCreateDTO Setter
// 阿信创建退款单入参
func (r *TaobaoAlitripAxinTransRefundCreateAPIRequest) SetAxinRefundCreateDTO(_axinRefundCreateDTO *AxinRefundCreateDTO) error {
    r._axinRefundCreateDTO = _axinRefundCreateDTO
    r.Set("axin_refund_create_d_t_o", _axinRefundCreateDTO)
    return nil
}

// AxinRefundCreateDTO Getter
func (r TaobaoAlitripAxinTransRefundCreateAPIRequest) GetAxinRefundCreateDTO() *AxinRefundCreateDTO {
    return r._axinRefundCreateDTO
}
