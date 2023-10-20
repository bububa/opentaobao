package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse 定向推广-新增或删除屏蔽词 API返回值
// alibaba.scbp.target.ad.plan.forbidden.word.modify
//
// 定向推广-新增或删除屏蔽词
type AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponseModel).Reset()
}

// AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponseModel is 定向推广-新增或删除屏蔽词 成功返回结果
type AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_forbidden_word_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true修改成功，fasle修改失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse)
	},
}

// GetAlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse 从 sync.Pool 获取 AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse
func GetAlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse() *AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse {
	return poolAlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse.Get().(*AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse)
}

// ReleaseAlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse 将 AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse(v *AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse.Put(v)
}
