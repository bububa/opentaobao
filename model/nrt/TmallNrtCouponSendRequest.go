package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
券发放接口 APIRequest
tmall.nrt.coupon.send

新零售场景，商家自有渠道发放券
*/
type TmallNrtCouponSendRequest struct {
    model.Params

    // 发券dto
    nrtCouponSendDto   *NrtCouponSendDTO 

}

func NewTmallNrtCouponSendRequest() *TmallNrtCouponSendRequest{
    return &TmallNrtCouponSendRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtCouponSendRequest) GetApiMethodName() string {
    return "tmall.nrt.coupon.send"
}

func (r TmallNrtCouponSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtCouponSendRequest) SetNrtCouponSendDto(nrtCouponSendDto *NrtCouponSendDTO) error {
    r.nrtCouponSendDto = nrtCouponSendDto
    r.Set("nrt_coupon_send_dto", nrtCouponSendDto)
    return nil
}

func (r TmallNrtCouponSendRequest) GetNrtCouponSendDto() *NrtCouponSendDTO {
    return r.nrtCouponSendDto
}

