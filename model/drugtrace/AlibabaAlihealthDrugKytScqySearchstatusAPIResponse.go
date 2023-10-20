package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqySearchstatusAPIResponse 单据处理状态查询 API返回值
// alibaba.alihealth.drug.kyt.scqy.searchstatus
//
// 单据处理状态查询
type AlibabaAlihealthDrugKytScqySearchstatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqySearchstatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqySearchstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytScqySearchstatusAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytScqySearchstatusAPIResponseModel is 单据处理状态查询 成功返回结果
type AlibabaAlihealthDrugKytScqySearchstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_searchstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytScqySearchstatusResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqySearchstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytScqySearchstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqySearchstatusAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytScqySearchstatusAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqySearchstatusAPIResponse
func GetAlibabaAlihealthDrugKytScqySearchstatusAPIResponse() *AlibabaAlihealthDrugKytScqySearchstatusAPIResponse {
	return poolAlibabaAlihealthDrugKytScqySearchstatusAPIResponse.Get().(*AlibabaAlihealthDrugKytScqySearchstatusAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytScqySearchstatusAPIResponse 将 AlibabaAlihealthDrugKytScqySearchstatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqySearchstatusAPIResponse(v *AlibabaAlihealthDrugKytScqySearchstatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqySearchstatusAPIResponse.Put(v)
}
