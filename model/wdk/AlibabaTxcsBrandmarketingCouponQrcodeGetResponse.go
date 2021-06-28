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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_txcs_brandmarketing_coupon_qrcode_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    ApiResult   *AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult `json:"api_result,omitempty" xml:"