package nazca

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据token获取出证申请信息 APIResponse
alibaba.nazca.token.issuecertapply.get

根据token获取出证申请信息
*/
type AlibabaNazcaTokenIssuecertapplyGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_nazca_token_issuecertapply_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ActionResult `json:"result,omitempty" xml:"