package nazca

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据token获取认证申请信息 APIResponse
alibaba.nazca.token.authapply.get

根据token获取认证申请信息
*/
type AlibabaNazcaTokenAuthapplyGetAPIResponse struct {
    model.CommonResponse
    AlibabaNazcaTokenAuthapplyGetResponse
}

type AlibabaNazcaTokenAuthapplyGetResponse struct {
    XMLName xml.Name `xml:"alibaba_nazca_token_authapply_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ActionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
