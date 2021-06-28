package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠规则删除 APIResponse
alibaba.asr.dataservice.promotionrule.delete

删除优惠规则，例如星巴克删除优惠规则
*/
type AlibabaAsrDataservicePromotionruleDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_asr_dataservice_promotionrule_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *DataServiceResponse `json:"result,omitempty" xml:"