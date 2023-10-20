package b2bcert

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaauthcertgetAPIResponse 获取证书数据 API返回值
// alibaba.auth.cert.get
//
// 获取证书数据
type AlibabaauthcertgetAPIResponse struct {
	model.CommonResponse
	AlibabaauthcertgetAPIResponseModel
}

// AlibabaauthcertgetAPIResponseModel is 获取证书数据 成功返回结果
type AlibabaauthcertgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_auth_cert_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaauthcertgetResponse `json:"result,omitempty" xml:"result,omitempty"`
}
