package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletGradeAPIResponse 获取流量档位 API返回值
// alibaba.aliqin.flow.wallet.grade
//
// 获取直充流量档位
type AlibabaAliqinFlowWalletGradeAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFlowWalletGradeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletGradeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFlowWalletGradeAPIResponseModel).Reset()
}

// AlibabaAliqinFlowWalletGradeAPIResponseModel is 获取流量档位 成功返回结果
type AlibabaAliqinFlowWalletGradeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_grade_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 档位
	Grade string `json:"grade,omitempty" xml:"grade,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletGradeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Grade = ""
}

var poolAlibabaAliqinFlowWalletGradeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFlowWalletGradeAPIResponse)
	},
}

// GetAlibabaAliqinFlowWalletGradeAPIResponse 从 sync.Pool 获取 AlibabaAliqinFlowWalletGradeAPIResponse
func GetAlibabaAliqinFlowWalletGradeAPIResponse() *AlibabaAliqinFlowWalletGradeAPIResponse {
	return poolAlibabaAliqinFlowWalletGradeAPIResponse.Get().(*AlibabaAliqinFlowWalletGradeAPIResponse)
}

// ReleaseAlibabaAliqinFlowWalletGradeAPIResponse 将 AlibabaAliqinFlowWalletGradeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFlowWalletGradeAPIResponse(v *AlibabaAliqinFlowWalletGradeAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFlowWalletGradeAPIResponse.Put(v)
}
