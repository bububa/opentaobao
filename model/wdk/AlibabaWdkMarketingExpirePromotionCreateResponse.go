package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠创建 APIResponse
alibaba.wdk.marketing.expire.promotion.create

过期优惠优惠信息录入
*/
type AlibabaWdkMarketingExpirePromotionCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingExpirePromotionCreateResponse `json:"alibaba_wdk_marketing_expire_promotion_create_response,omitempty"` 
    AlibabaWdkMarketingExpirePromotionCreateResponse
}

/* model for simplify = false
type AlibabaWdkMarketingExpirePromotionCreateResponse struct {

    // message
    
    Message   string `json:"message,omitempty"`
    

    // data
    
    Datas  struct {
        T  []T `json:"t,omitempty"`
    } `json:"datas,omitempty"`
    

    // errorCode
    
    FailCode   string `json:"fail_code,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type AlibabaWdkMarketingExpirePromotionCreateResponse struct {

    // message
    Message   string `json:"message,omitempty"`

    // data
    Datas   []T `json:"datas,omitempty"`

    // errorCode
    FailCode   string `json:"fail_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
