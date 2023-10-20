package seaking

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaseakingtitlerewriteresultAPIResponse 获取标题改写任务结果 API返回值
// alibaba.seaking.titlerewrite.result
//
// 获取标题改写任务结果
type AlibabaseakingtitlerewriteresultAPIResponse struct {
	model.CommonResponse
	AlibabaseakingtitlerewriteresultAPIResponseModel
}

// AlibabaseakingtitlerewriteresultAPIResponseModel is 获取标题改写任务结果 成功返回结果
type AlibabaseakingtitlerewriteresultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_titlerewrite_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
