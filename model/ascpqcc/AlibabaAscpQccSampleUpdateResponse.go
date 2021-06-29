package ascpqcc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品控中心更新样品信息 APIResponse
alibaba.ascp.qcc.sample.update

品控中心更新样品信息
*/
type AlibabaAscpQccSampleUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAscpQccSampleUpdateResponse
}

type AlibabaAscpQccSampleUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_qcc_sample_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *SendResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
