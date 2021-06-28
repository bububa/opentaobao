package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠删除 APIResponse
alibaba.wdk.marketing.expire.promotion.delete

短保优惠删除
*/
type AlibabaWdkMarketingExpirePromotionDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_expire_promotion_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // message
    
    Message   string `json:"message,omitempty" xml:"