package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销导购员券页面二维码获取 APIResponse
alibaba.txcs.brandmarketing.coupon.qrcode.get

构建券页码二维码url
*/
type AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse struct {
    model.CommonResponse
    AlibabaTxcsBrandmarketingCouponQrcodeGetResponse
}

type AlibabaTxcsBrandmarketingCouponQrcodeGetResponse struct {
    XMLName xml.Name `xml:"alibaba_txcs_brandmarketing_coupon_qrcode_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    ApiResult   *AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
