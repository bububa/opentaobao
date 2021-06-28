package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠删除 APIResponse
alibaba.wdk.marketing.expire.promotion.delete

短保优惠删除
*/
type AlibabaWdkMarketingExpirePromotionDeleteAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingExpirePromotionDeleteResponse `json:"alibaba_wdk_marketing_expire_promotion_delete_response,omitempty"` 
    AlibabaWdkMarketingExpirePromotionDeleteResponse
}

/* model for simplify = false
type AlibabaWdkMarketingExpirePromotionDeleteResponse struct {

    // message
    
    Message   string `json:"message,omitempty"`
    

    // data
    
    Datas  struct {
        ExpirePromotionResult  []ExpirePromotionResult `json:"expire_promotion_result,omitempty"`
    } `json:"datas,omitempty"`
    

    // errorCode
    
    FailCode   string `json:"fail_code,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type AlibabaWdkMarketingExpirePromotionDeleteResponse struct {

    // message
    Message   string `json:"message,omitempty"`

    // data
    Datas   []ExpirePromotionResult `json:"datas,omitempty"`

    // errorCode
    FailCode   string `json:"fail_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
