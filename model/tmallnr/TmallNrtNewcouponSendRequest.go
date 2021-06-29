package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
券发放接口 APIRequest
tmall.nrt.newcoupon.send

券发放接口
*/
type TmallNrtNewcouponSendRequest struct {
    model.Params

    // 发券dto
    nrtCouponSendDto   *NrtCouponSendDto 

}

func NewTmallNrtNewcouponSendRequest() *TmallNrtNewcouponSendRequest{
    return &TmallNrtNewcouponSendRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtNewcouponSendRequest) GetApiMethodName() string {
    return "tmall.nrt.newcoupon.send"
}

func (r TmallNrtNewcouponSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtNewcouponSendRequest) SetNrtCouponSendDto(nrtCouponSendDto *NrtCouponSendDto) error {
    r.nrtCouponSendDto = nrtCouponSendDto
    r.Set("nrt_coupon_send_dto", nrtCouponSendDto)
    return nil
}

func (r TmallNrtNewcouponSendRequest) GetNrtCouponSendDto() *NrtCouponSendDto {
    return r.nrtCouponSendDto
}

