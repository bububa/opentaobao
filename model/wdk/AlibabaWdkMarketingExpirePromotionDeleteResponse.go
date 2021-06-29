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
    AlibabaWdkMarketingExpirePromotionDeleteResponse
}

type AlibabaWdkMarketingExpirePromotionDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_expire_promotion_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // data
    
    Datas   []ExpirePromotionResult `json:"datas,omitempty" xml:"datas>expire_promotion_result,omitempty"`
    
    
    // errorCode
    
    FailCode   string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
