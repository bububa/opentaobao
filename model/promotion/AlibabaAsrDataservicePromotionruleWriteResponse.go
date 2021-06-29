package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
业务优惠规则写入 APIResponse
alibaba.asr.dataservice.promotionrule.write

星巴克优惠规则写入
*/
type AlibabaAsrDataservicePromotionruleWriteAPIResponse struct {
    model.CommonResponse
    AlibabaAsrDataservicePromotionruleWriteResponse
}

type AlibabaAsrDataservicePromotionruleWriteResponse struct {
    XMLName xml.Name `xml:"alibaba_asr_dataservice_promotionrule_write_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *DataServiceResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
