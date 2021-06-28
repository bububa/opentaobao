package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星巴克优惠规则查询 APIResponse
alibaba.asr.dataservice.promotionrule.query

查询优惠规则，例如星巴克查询优惠规则
*/
type AlibabaAsrDataservicePromotionruleQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_asr_dataservice_promotionrule_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *DataServiceResponse `json:"result,omitempty" xml:"