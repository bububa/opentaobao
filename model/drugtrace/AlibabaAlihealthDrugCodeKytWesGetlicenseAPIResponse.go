package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodekytwesgetlicenseAPIResponse 获取licenseToken API返回值
// alibaba.alihealth.drug.code.kyt.wes.getlicense
//
// 获取licenseToken
type AlibabaalihealthdrugcodekytwesgetlicenseAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugcodekytwesgetlicenseAPIResponseModel
}

// AlibabaalihealthdrugcodekytwesgetlicenseAPIResponseModel is 获取licenseToken 成功返回结果
type AlibabaalihealthdrugcodekytwesgetlicenseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_wes_getlicense_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaalihealthdrugcodekytwesgetlicenseResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
