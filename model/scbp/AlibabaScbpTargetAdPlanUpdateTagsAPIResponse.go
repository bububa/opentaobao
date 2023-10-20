package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanUpdateTagsAPIResponse 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新 API返回值
// alibaba.scbp.target.ad.plan.update.tags
//
// 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新
type AlibabaScbpTargetAdPlanUpdateTagsAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanUpdateTagsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanUpdateTagsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTargetAdPlanUpdateTagsAPIResponseModel).Reset()
}

// AlibabaScbpTargetAdPlanUpdateTagsAPIResponseModel is 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新 成功返回结果
type AlibabaScbpTargetAdPlanUpdateTagsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_update_tags_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改记录数量
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanUpdateTagsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpTargetAdPlanUpdateTagsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTargetAdPlanUpdateTagsAPIResponse)
	},
}

// GetAlibabaScbpTargetAdPlanUpdateTagsAPIResponse 从 sync.Pool 获取 AlibabaScbpTargetAdPlanUpdateTagsAPIResponse
func GetAlibabaScbpTargetAdPlanUpdateTagsAPIResponse() *AlibabaScbpTargetAdPlanUpdateTagsAPIResponse {
	return poolAlibabaScbpTargetAdPlanUpdateTagsAPIResponse.Get().(*AlibabaScbpTargetAdPlanUpdateTagsAPIResponse)
}

// ReleaseAlibabaScbpTargetAdPlanUpdateTagsAPIResponse 将 AlibabaScbpTargetAdPlanUpdateTagsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanUpdateTagsAPIResponse(v *AlibabaScbpTargetAdPlanUpdateTagsAPIResponse) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanUpdateTagsAPIResponse.Put(v)
}
