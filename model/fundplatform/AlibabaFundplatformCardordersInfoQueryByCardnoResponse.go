package fundplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过卡号查询卡信息 APIResponse
alibaba.fundplatform.cardorders.info.query.by.cardno

该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
*/
type AlibabaFundplatformCardordersInfoQueryByCardnoAPIResponse struct {
    model.CommonResponse
    AlibabaFundplatformCardordersInfoQueryByCardnoResponse
}

type AlibabaFundplatformCardordersInfoQueryByCardnoResponse struct {
    XMLName xml.Name `xml:"alibaba_fundplatform_cardorders_info_query_by_cardno_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CardMakingInfoQueryResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
