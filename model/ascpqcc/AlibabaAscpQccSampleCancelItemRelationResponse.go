package ascpqcc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
魅力惠样品解除父子商品关系 APIResponse
alibaba.ascp.qcc.sample.cancel.item.relation

品控中心魅力惠样品解除父子商品关系
*/
type AlibabaAscpQccSampleCancelItemRelationAPIResponse struct {
    model.CommonResponse
    AlibabaAscpQccSampleCancelItemRelationResponse
}

type AlibabaAscpQccSampleCancelItemRelationResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_qcc_sample_cancel_item_relation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *SendResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
