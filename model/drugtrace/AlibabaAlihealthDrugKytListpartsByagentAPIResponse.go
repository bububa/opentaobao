package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytListpartsByagentAPIResponse 物流代货主查往来单位接口 API返回值
// alibaba.alihealth.drug.kyt.listparts.byagent
//
// 代理企业查询往来单位列表
type AlibabaAlihealthDrugKytListpartsByagentAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytListpartsByagentAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytListpartsByagentAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytListpartsByagentAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytListpartsByagentAPIResponseModel is 物流代货主查往来单位接口 成功返回结果
type AlibabaAlihealthDrugKytListpartsByagentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_listparts_byagent_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytListpartsByagentResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytListpartsByagentAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytListpartsByagentAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytListpartsByagentAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytListpartsByagentAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytListpartsByagentAPIResponse
func GetAlibabaAlihealthDrugKytListpartsByagentAPIResponse() *AlibabaAlihealthDrugKytListpartsByagentAPIResponse {
	return poolAlibabaAlihealthDrugKytListpartsByagentAPIResponse.Get().(*AlibabaAlihealthDrugKytListpartsByagentAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytListpartsByagentAPIResponse 将 AlibabaAlihealthDrugKytListpartsByagentAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytListpartsByagentAPIResponse(v *AlibabaAlihealthDrugKytListpartsByagentAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytListpartsByagentAPIResponse.Put(v)
}
