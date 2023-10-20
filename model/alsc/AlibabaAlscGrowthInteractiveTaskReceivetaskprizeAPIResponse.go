package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscgrowthinteractivetaskreceivetaskprizeAPIResponse 任务领奖 API返回值
// alibaba.alsc.growth.interactive.task.receivetaskprize
//
// 任务领奖
type AlibabaalscgrowthinteractivetaskreceivetaskprizeAPIResponse struct {
	model.CommonResponse
	AlibabaalscgrowthinteractivetaskreceivetaskprizeAPIResponseModel
}

// AlibabaalscgrowthinteractivetaskreceivetaskprizeAPIResponseModel is 任务领奖 成功返回结果
type AlibabaalscgrowthinteractivetaskreceivetaskprizeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_task_receivetaskprize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发奖结果
	Result *BaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
