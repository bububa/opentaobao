package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
券发放接口 APIResponse
tmall.nrt.coupon.send

新零售场景，商家自有渠道发放券
*/
type TmallNrtCouponSendAPIResponse struct {
    model.CommonResponse
    TmallNrtCouponSendResponse
}

type TmallNrtCouponSendResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_coupon_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 发券结果对象
    
    Model   *SendCouponResponse `json:"model,omitempty" xml:"model,omitempty"`

    
}
