package nazca

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNazcaTokenIssuecertapplyGetAPIResponse
根据token获取出证申请信息 API返回值
alibaba.nazca.token.issuecertapply.get

根据token获取出证申请信息 */
type AlibabaNazcaTokenIssuecertapplyGetAPIResponse struct {
	model.CommonResponse
	AlibabaNazcaTokenIssuecertapplyGetAPIResponseModel
}

// AlibabaNazcaTokenIssuecertapplyGetAPIResponseModel is 根据token获取出证申请信息 成功返回结果
type AlibabaNazcaTokenIssuecertapplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nazca_token_issuecertapply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
