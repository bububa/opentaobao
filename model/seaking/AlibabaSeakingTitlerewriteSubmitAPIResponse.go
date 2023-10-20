package seaking

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaseakingtitlerewritesubmitAPIResponse 提交标题改写任务 API返回值
// alibaba.seaking.titlerewrite.submit
//
// 提交标题改写任务
type AlibabaseakingtitlerewritesubmitAPIResponse struct {
	model.CommonResponse
	AlibabaseakingtitlerewritesubmitAPIResponseModel
}

// AlibabaseakingtitlerewritesubmitAPIResponseModel is 提交标题改写任务 成功返回结果
type AlibabaseakingtitlerewritesubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_titlerewrite_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}
