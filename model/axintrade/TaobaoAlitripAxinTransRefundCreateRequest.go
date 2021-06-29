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
type TaobaoAlitripAxinTransRefundCreateRequest struct {
    model.Params
    // 阿信创建退款单入参
    axinRefundCreateDTO   *AxinRefundCreateDto
}

// 初始化TaobaoAlitripAxinTransRefundCreateRequest对象
func NewTaobaoAlitripAxinTransRefundCreateRequest() *TaobaoAlitripAxinTransRefundCreateRequest{
    return &TaobaoAlitripAxinTransRefundCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransRefundCreateRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.refund.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransRefundCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AxinRefundCreateDTO Setter
// 阿信创建退款单入参
func (r *TaobaoAlitripAxinTransRefundCreateRequest) SetAxinRefundCreateDTO(axinRefundCreateDTO *AxinRefundCreateDto) error {
    r.axinRefundCreateDTO = axinRefundCreateDTO
    r.Set("axin_refund_create_d_t_o", axinRefundCreateDTO)
    return nil
}

// AxinRefundCreateDTO Getter
func (r TaobaoAlitripAxinTransRefundCreateRequest) GetAxinRefundCreateDTO() *AxinRefundCreateDto {
    return r.axinRefundCreateDTO
}
