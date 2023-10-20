package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrugrescodeAPIResponse 查询药品码段信息 API返回值
// alibaba.alihealth.drug.kyt.drugrescode
//
// 查询药品码段信息
type AlibabaAlihealthDrugKytDrugrescodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrugrescodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrugrescodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrugrescodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrugrescodeAPIResponseModel is 查询药品码段信息 成功返回结果
type AlibabaAlihealthDrugKytDrugrescodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_drugrescode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytDrugrescodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrugrescodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrugrescodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrugrescodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrugrescodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrugrescodeAPIResponse
func GetAlibabaAlihealthDrugKytDrugrescodeAPIResponse() *AlibabaAlihealthDrugKytDrugrescodeAPIResponse {
	return poolAlibabaAlihealthDrugKytDrugrescodeAPIResponse.Get().(*AlibabaAlihealthDrugKytDrugrescodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrugrescodeAPIResponse 将 AlibabaAlihealthDrugKytDrugrescodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrugrescodeAPIResponse(v *AlibabaAlihealthDrugKytDrugrescodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrugrescodeAPIResponse.Put(v)
}
