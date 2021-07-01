package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse
定向推广-新增或删除屏蔽词 API返回值
alibaba.scbp.target.ad.plan.forbidden.word.modify

定向推广-新增或删除屏蔽词 */
type AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponseModel
}

// AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponseModel is 定向推广-新增或删除屏蔽词 成功返回结果
type AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_forbidden_word_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true修改成功，fasle修改失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
