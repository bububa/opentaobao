package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse 物流代货主查找往来单位接口 API返回值
// alibaba.alihealth.drug.kyt.wes.listparts.byagent
//
// 代理企业查询往来单位列表
type AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesListpartsByagentAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesListpartsByagentAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesListpartsByagentAPIResponseModel is 物流代货主查找往来单位接口 成功返回结果
type AlibabaAlihealthDrugKytWesListpartsByagentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_listparts_byagent_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesListpartsByagentResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesListpartsByagentAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytWesListpartsByagentAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesListpartsByagentAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse
func GetAlibabaAlihealthDrugKytWesListpartsByagentAPIResponse() *AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse {
	return poolAlibabaAlihealthDrugKytWesListpartsByagentAPIResponse.Get().(*AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesListpartsByagentAPIResponse 将 AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesListpartsByagentAPIResponse(v *AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesListpartsByagentAPIResponse.Put(v)
}
