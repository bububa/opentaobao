package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentermsftasksupdateAPIResponse 喵师傅工人任务批量完成接口 API返回值
// tmall.servicecenter.msf.tasks.update
//
// 喵师傅工人任务批量完成接口
type TmallservicecentermsftasksupdateAPIResponse struct {
	model.CommonResponse
	TmallservicecentermsftasksupdateAPIResponseModel
}

// TmallservicecentermsftasksupdateAPIResponseModel is 喵师傅工人任务批量完成接口 成功返回结果
type TmallservicecentermsftasksupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_msf_tasks_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
