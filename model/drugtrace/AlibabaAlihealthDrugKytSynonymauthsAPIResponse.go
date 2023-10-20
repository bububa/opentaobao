package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSynonymauthsAPIResponse 物流企业查询货主企业信息 API返回值
// alibaba.alihealth.drug.kyt.synonymauths
//
// 物流企业查询货主企业信息
type AlibabaAlihealthDrugKytSynonymauthsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSynonymauthsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSynonymauthsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSynonymauthsAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSynonymauthsAPIResponseModel is 物流企业查询货主企业信息 成功返回结果
type AlibabaAlihealthDrugKytSynonymauthsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_synonymauths_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytSynonymauthsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSynonymauthsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytSynonymauthsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSynonymauthsAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSynonymauthsAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSynonymauthsAPIResponse
func GetAlibabaAlihealthDrugKytSynonymauthsAPIResponse() *AlibabaAlihealthDrugKytSynonymauthsAPIResponse {
	return poolAlibabaAlihealthDrugKytSynonymauthsAPIResponse.Get().(*AlibabaAlihealthDrugKytSynonymauthsAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSynonymauthsAPIResponse 将 AlibabaAlihealthDrugKytSynonymauthsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSynonymauthsAPIResponse(v *AlibabaAlihealthDrugKytSynonymauthsAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSynonymauthsAPIResponse.Put(v)
}
