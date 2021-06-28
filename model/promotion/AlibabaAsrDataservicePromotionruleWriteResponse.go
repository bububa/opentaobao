package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
业务优惠规则写入 APIResponse
alibaba.asr.dataservice.promotionrule.write

星巴克优惠规则写入
*/
type AlibabaAsrDataservicePromotionruleWriteAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAsrDataservicePromotionruleWriteResponse `json:"alibaba_asr_dataservice_promotionrule_write_response,omitempty"` 
    AlibabaAsrDataservicePromotionruleWriteResponse
}

/* model for simplify = false
type AlibabaAsrDataservicePromotionruleWriteResponse struct {

    // 结果
    
    Result  *struct {
        DataServiceResponse  *DataServiceResponse `json:"data_service_response,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAsrDataservicePromotionruleWriteResponse struct {

    // 结果
    Result   *DataServiceResponse `json:"result,omitempty"`

}
