package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
券发放接口 API请求
tmall.nrt.newcoupon.send

券发放接口
*/
type TmallNrtNewcouponSendAPIRequest struct {
    model.Params
    // 发券dto
    _nrtCouponSendDto   *NrtCouponSendDTO
}

// 初始化TmallNrtNewcouponSendAPIRequest对象
func NewTmallNrtNewcouponSendRequest() *TmallNrtNewcouponSendAPIRequest{
    return &TmallNrtNewcouponSendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtNewcouponSendAPIRequest) GetApiMethodName() string {
    return "tmall.nrt.newcoupon.send"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtNewcouponSendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtCouponSendDto Setter
// 发券dto
func (r *TmallNrtNewcouponSendAPIRequest) SetNrtCouponSendDto(_nrtCouponSendDto *NrtCouponSendDTO) error {
    r._nrtCouponSendDto = _nrtCouponSendDto
    r.Set("nrt_coupon_send_dto", _nrtCouponSendDto)
    return nil
}

// NrtCouponSendDto Getter
func (r TmallNrtNewcouponSendAPIRequest) GetNrtCouponSendDto() *NrtCouponSendDTO {
    return r._nrtCouponSendDto
}
