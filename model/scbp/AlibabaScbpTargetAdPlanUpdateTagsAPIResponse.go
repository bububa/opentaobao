package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplanupdatetagsAPIResponse 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新 API返回值
// alibaba.scbp.target.ad.plan.update.tags
//
// 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新
type AlibabascbptargetadplanupdatetagsAPIResponse struct {
	model.CommonResponse
	AlibabascbptargetadplanupdatetagsAPIResponseModel
}

// AlibabascbptargetadplanupdatetagsAPIResponseModel is 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新 成功返回结果
type AlibabascbptargetadplanupdatetagsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_update_tags_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改记录数量
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
