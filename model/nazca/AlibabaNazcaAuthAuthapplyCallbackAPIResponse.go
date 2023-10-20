package nazca

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabanazcaauthauthapplycallbackAPIResponse 认证的统一回调接口 API返回值
// alibaba.nazca.auth.authapply.callback
//
// 认证的统一回调接口
type AlibabanazcaauthauthapplycallbackAPIResponse struct {
	model.CommonResponse
	AlibabanazcaauthauthapplycallbackAPIResponseModel
}

// AlibabanazcaauthauthapplycallbackAPIResponseModel is 认证的统一回调接口 成功返回结果
type AlibabanazcaauthauthapplycallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nazca_auth_authapply_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
