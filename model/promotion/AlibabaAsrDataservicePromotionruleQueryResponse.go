package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
星巴克优惠规则查询 APIResponse
alibaba.asr.dataservice.promotionrule.query

查询优惠规则，例如星巴克查询优惠规则
*/
type AlibabaAsrDataservicePromotionruleQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAsrDataservicePromotionruleQueryResponse `json:"alibaba_asr_dataservice_promotionrule_query_response,omitempty"`
}

type AlibabaAsrDataservicePromotionruleQueryResponse struct {

    // 结果
    Result   *DataServiceResponse `json:"result,omitempty"`

}
