package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销导购员券页面二维码获取 APIResponse
alibaba.txcs.brandmarketing.coupon.qrcode.get

构建券页码二维码url
*/
type AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaTxcsBrandmarketingCouponQrcodeGetResponse `json:"alibaba_txcs_brandmarketing_coupon_qrcode_get_response,omitempty"`
}

type AlibabaTxcsBrandmarketingCouponQrcodeGetResponse struct {

    // 返回结果
    ApiResult   *AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult `json:"api_result,omitempty"`

}
