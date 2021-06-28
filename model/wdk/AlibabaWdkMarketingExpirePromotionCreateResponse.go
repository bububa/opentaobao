package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠创建 APIResponse
alibaba.wdk.marketing.expire.promotion.create

过期优惠优惠信息录入
*/
type AlibabaWdkMarketingExpirePromotionCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_expire_promotion_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // message
    
    Message   string `json:"message,omitempty" xml:"