package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse 关联关系处理查询 API返回值
// alibaba.alihealth.drug.kyt.scqy.codeprocess
//
// 关联关系处理查询
type AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqyCodeprocessAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytScqyCodeprocessAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytScqyCodeprocessAPIResponseModel is 关联关系处理查询 成功返回结果
type AlibabaAlihealthDrugKytScqyCodeprocessAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_codeprocess_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytScqyCodeprocessResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyCodeprocessAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytScqyCodeprocessAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytScqyCodeprocessAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse
func GetAlibabaAlihealthDrugKytScqyCodeprocessAPIResponse() *AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse {
	return poolAlibabaAlihealthDrugKytScqyCodeprocessAPIResponse.Get().(*AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytScqyCodeprocessAPIResponse 将 AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqyCodeprocessAPIResponse(v *AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqyCodeprocessAPIResponse.Put(v)
}
