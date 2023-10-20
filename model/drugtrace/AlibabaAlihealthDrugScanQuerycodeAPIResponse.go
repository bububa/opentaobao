package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugScanQuerycodeAPIResponse 查询药监码对应的有效期和包装规格 API返回值
// alibaba.alihealth.drug.scan.querycode
//
// 查询药监码对应的有效期和包装规格
type AlibabaAlihealthDrugScanQuerycodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugScanQuerycodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugScanQuerycodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugScanQuerycodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugScanQuerycodeAPIResponseModel is 查询药监码对应的有效期和包装规格 成功返回结果
type AlibabaAlihealthDrugScanQuerycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_scan_querycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugScanQuerycodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugScanQuerycodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugScanQuerycodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugScanQuerycodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugScanQuerycodeAPIResponse
func GetAlibabaAlihealthDrugScanQuerycodeAPIResponse() *AlibabaAlihealthDrugScanQuerycodeAPIResponse {
	return poolAlibabaAlihealthDrugScanQuerycodeAPIResponse.Get().(*AlibabaAlihealthDrugScanQuerycodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugScanQuerycodeAPIResponse 将 AlibabaAlihealthDrugScanQuerycodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugScanQuerycodeAPIResponse(v *AlibabaAlihealthDrugScanQuerycodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugScanQuerycodeAPIResponse.Put(v)
}
