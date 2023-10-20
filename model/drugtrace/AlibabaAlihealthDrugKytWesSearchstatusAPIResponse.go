package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesSearchstatusAPIResponse 单据处理状态查询 API返回值
// alibaba.alihealth.drug.kyt.wes.searchstatus
//
// 单据处理状态查询
type AlibabaAlihealthDrugKytWesSearchstatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesSearchstatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesSearchstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesSearchstatusAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesSearchstatusAPIResponseModel is 单据处理状态查询 成功返回结果
type AlibabaAlihealthDrugKytWesSearchstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_searchstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesSearchstatusResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesSearchstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytWesSearchstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesSearchstatusAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesSearchstatusAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesSearchstatusAPIResponse
func GetAlibabaAlihealthDrugKytWesSearchstatusAPIResponse() *AlibabaAlihealthDrugKytWesSearchstatusAPIResponse {
	return poolAlibabaAlihealthDrugKytWesSearchstatusAPIResponse.Get().(*AlibabaAlihealthDrugKytWesSearchstatusAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesSearchstatusAPIResponse 将 AlibabaAlihealthDrugKytWesSearchstatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesSearchstatusAPIResponse(v *AlibabaAlihealthDrugKytWesSearchstatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesSearchstatusAPIResponse.Put(v)
}
