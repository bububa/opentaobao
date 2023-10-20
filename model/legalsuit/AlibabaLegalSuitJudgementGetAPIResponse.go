package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitJudgementGetAPIResponse 获取裁判登记信息 API返回值
// alibaba.legal.suit.judgement.get
//
// 供ISV供应商获取集团法务系统的裁判登记信息
type AlibabaLegalSuitJudgementGetAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitJudgementGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalSuitJudgementGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalSuitJudgementGetAPIResponseModel).Reset()
}

// AlibabaLegalSuitJudgementGetAPIResponseModel is 获取裁判登记信息 成功返回结果
type AlibabaLegalSuitJudgementGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_judgement_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalSuitJudgementGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalSuitJudgementGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalSuitJudgementGetAPIResponse)
	},
}

// GetAlibabaLegalSuitJudgementGetAPIResponse 从 sync.Pool 获取 AlibabaLegalSuitJudgementGetAPIResponse
func GetAlibabaLegalSuitJudgementGetAPIResponse() *AlibabaLegalSuitJudgementGetAPIResponse {
	return poolAlibabaLegalSuitJudgementGetAPIResponse.Get().(*AlibabaLegalSuitJudgementGetAPIResponse)
}

// ReleaseAlibabaLegalSuitJudgementGetAPIResponse 将 AlibabaLegalSuitJudgementGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalSuitJudgementGetAPIResponse(v *AlibabaLegalSuitJudgementGetAPIResponse) {
	v.Reset()
	poolAlibabaLegalSuitJudgementGetAPIResponse.Put(v)
}
