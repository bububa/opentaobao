package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSearchstatusAPIResponse 单据处理状态查询 API返回值
// alibaba.alihealth.drug.kyt.searchstatus
//
// 单据处理状态查询
type AlibabaAlihealthDrugKytSearchstatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSearchstatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSearchstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSearchstatusAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSearchstatusAPIResponseModel is 单据处理状态查询 成功返回结果
type AlibabaAlihealthDrugKytSearchstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_searchstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytSearchstatusResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSearchstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytSearchstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSearchstatusAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSearchstatusAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSearchstatusAPIResponse
func GetAlibabaAlihealthDrugKytSearchstatusAPIResponse() *AlibabaAlihealthDrugKytSearchstatusAPIResponse {
	return poolAlibabaAlihealthDrugKytSearchstatusAPIResponse.Get().(*AlibabaAlihealthDrugKytSearchstatusAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSearchstatusAPIResponse 将 AlibabaAlihealthDrugKytSearchstatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSearchstatusAPIResponse(v *AlibabaAlihealthDrugKytSearchstatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSearchstatusAPIResponse.Put(v)
}
