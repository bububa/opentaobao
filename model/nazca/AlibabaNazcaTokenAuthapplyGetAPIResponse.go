package nazca

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNazcaTokenAuthapplyGetAPIResponse
根据token获取认证申请信息 API返回值
alibaba.nazca.token.authapply.get

根据token获取认证申请信息 */
type AlibabaNazcaTokenAuthapplyGetAPIResponse struct {
	model.CommonResponse
	AlibabaNazcaTokenAuthapplyGetAPIResponseModel
}

// AlibabaNazcaTokenAuthapplyGetAPIResponseModel is 根据token获取认证申请信息 成功返回结果
type AlibabaNazcaTokenAuthapplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nazca_token_authapply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
