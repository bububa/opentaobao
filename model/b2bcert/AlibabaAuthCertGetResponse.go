package b2bcert

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取证书数据 APIResponse
alibaba.auth.cert.get

获取证书数据
*/
type AlibabaAuthCertGetAPIResponse struct {
    model.CommonResponse
    AlibabaAuthCertGetResponse
}

type AlibabaAuthCertGetResponse struct {
    XMLName xml.Name `xml:"alibaba_auth_cert_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
