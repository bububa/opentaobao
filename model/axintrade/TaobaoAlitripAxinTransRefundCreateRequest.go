package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿信创建退款单 APIRequest
taobao.alitrip.axin.trans.refund.create

阿信供销平台-创建退款单服务
*/
type TaobaoAlitripAxinTransRefundCreateRequest struct {
    model.Params

    // 阿信创建退款单入参
    axinRefundCreateDTO   *AxinRefundCreateDto 

}

func NewTaobaoAlitripAxinTransRefundCreateRequest() *TaobaoAlitripAxinTransRefundCreateRequest{
    return &TaobaoAlitripAxinTransRefundCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripAxinTransRefundCreateRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.refund.create"
}

func (r TaobaoAlitripAxinTransRefundCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripAxinTransRefundCreateRequest) SetAxinRefundCreateDTO(axinRefundCreateDTO *AxinRefundCreateDto) error {
    r.axinRefundCreateDTO = axinRefundCreateDTO
    r.Set("axin_refund_create_d_t_o", axinRefundCreateDTO)
    return nil
}

func (r TaobaoAlitripAxinTransRefundCreateRequest) GetAxinRefundCreateDTO() *AxinRefundCreateDto {
    return r.axinRefundCreateDTO
}

