package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthstorecertificatecreateAPIResponse 仓库换证审批 API返回值
// alibaba.alihealth.store.certificate.create
//
// 仓库侧换证发起审批
type AlibabaalihealthstorecertificatecreateAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthstorecertificatecreateAPIResponseModel
}

// AlibabaalihealthstorecertificatecreateAPIResponseModel is 仓库换证审批 成功返回结果
type AlibabaalihealthstorecertificatecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_store_certificate_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AlibabaalihealthstorecertificatecreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
