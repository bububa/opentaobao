package nazca

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabanazcatokenchangeauthapplygetAPIResponse 根据token获取变更认证申请信息 API返回值
// alibaba.nazca.token.changeauthapply.get
//
// 根据token获取变更认证申请信息
type AlibabanazcatokenchangeauthapplygetAPIResponse struct {
	model.CommonResponse
	AlibabanazcatokenchangeauthapplygetAPIResponseModel
}

// AlibabanazcatokenchangeauthapplygetAPIResponseModel is 根据token获取变更认证申请信息 成功返回结果
type AlibabanazcatokenchangeauthapplygetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nazca_token_changeauthapply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
