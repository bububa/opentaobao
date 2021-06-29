package tmalltrend

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全域新品店铺优惠券免除 APIResponse
tmall.brand.item.coupon.protect

全域新品店铺优惠券免除申请打标接口
*/
type TmallBrandItemCouponProtectAPIResponse struct {
    model.CommonResponse
    TmallBrandItemCouponProtectResponse
}

type TmallBrandItemCouponProtectResponse struct {
    XMLName xml.Name `xml:"tmall_brand_item_coupon_protect_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`

    
    // 店铺优惠券保护期设置是否成功
    
    RespSuccess   bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`

    
    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
    // 错误码
    
    RespErrorCode   string `json:"resp_error_code,omitempty" xml:"resp_error_code,omitempty"`

    
}
