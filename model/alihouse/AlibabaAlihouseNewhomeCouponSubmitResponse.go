package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交专车优惠券活动 APIResponse
alibaba.alihouse.newhome.coupon.submit

提交专车优惠券活动
*/
type AlibabaAlihouseNewhomeCouponSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeCouponSubmitResponse
}

type AlibabaAlihouseNewhomeCouponSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_coupon_submit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeCouponSubmitResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
