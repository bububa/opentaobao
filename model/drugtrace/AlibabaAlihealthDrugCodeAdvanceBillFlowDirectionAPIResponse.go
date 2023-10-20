package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse 单据流向查询 API返回值
// alibaba.alihealth.drug.code.advance.bill.flow.direction
//
// 单据流向查询
type AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponseModel is 单据流向查询 成功返回结果
type AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_advance_bill_flow_direction_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse
func GetAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse() *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse {
	return poolAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse.Get().(*AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse 将 AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse(v *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse.Put(v)
}
