package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytQuerycodeAPIResponse 查询追溯码对应的药品信息 API返回值
// alibaba.alihealth.drug.code.kyt.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytQuerycodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytQuerycodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytQuerycodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeKytQuerycodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeKytQuerycodeAPIResponseModel is 查询追溯码对应的药品信息 成功返回结果
type AlibabaAlihealthDrugCodeKytQuerycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_querycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugCodeKytQuerycodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytQuerycodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeKytQuerycodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytQuerycodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeKytQuerycodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytQuerycodeAPIResponse
func GetAlibabaAlihealthDrugCodeKytQuerycodeAPIResponse() *AlibabaAlihealthDrugCodeKytQuerycodeAPIResponse {
	return poolAlibabaAlihealthDrugCodeKytQuerycodeAPIResponse.Get().(*AlibabaAlihealthDrugCodeKytQuerycodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeKytQuerycodeAPIResponse 将 AlibabaAlihealthDrugCodeKytQuerycodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytQuerycodeAPIResponse(v *AlibabaAlihealthDrugCodeKytQuerycodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytQuerycodeAPIResponse.Put(v)
}
