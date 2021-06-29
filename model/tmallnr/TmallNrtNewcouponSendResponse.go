package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
券发放接口 APIResponse
tmall.nrt.newcoupon.send

券发放接口
*/
type TmallNrtNewcouponSendAPIResponse struct {
    model.CommonResponse
    TmallNrtNewcouponSendResponse
}

type TmallNrtNewcouponSendResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_newcoupon_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TmallNrtNewcouponSendResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
