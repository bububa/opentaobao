package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse 获取licenseToken API返回值
// alibaba.alihealth.drug.code.kyt.wes.getlicense
//
// 获取licenseToken
type AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponseModel is 获取licenseToken 成功返回结果
type AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_wes_getlicense_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse
func GetAlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse() *AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse {
	return poolAlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse.Get().(*AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse 将 AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse(v *AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse.Put(v)
}
