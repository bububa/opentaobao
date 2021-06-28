package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠规则删除 APIResponse
alibaba.asr.dataservice.promotionrule.delete

删除优惠规则，例如星巴克删除优惠规则
*/
type AlibabaAsrDataservicePromotionruleDeleteAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAsrDataservicePromotionruleDeleteResponse `json:"alibaba_asr_dataservice_promotionrule_delete_response,omitempty"` 
    AlibabaAsrDataservicePromotionruleDeleteResponse
}

/* model for simplify = false
type AlibabaAsrDataservicePromotionruleDeleteResponse struct {

    // 结果
    
    Result  *struct {
        DataServiceResponse  *DataServiceResponse `json:"data_service_response,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAsrDataservicePromotionruleDeleteResponse struct {

    // 结果
    Result   *DataServiceResponse `json:"result,omitempty"`

}
