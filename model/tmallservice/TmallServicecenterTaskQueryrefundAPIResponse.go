package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentertaskqueryrefundAPIResponse 查询任务类工单是否退款 API返回值
// tmall.servicecenter.task.queryrefund
//
// 查询任务类工单是否退款
type TmallservicecentertaskqueryrefundAPIResponse struct {
	model.CommonResponse
	TmallservicecentertaskqueryrefundAPIResponseModel
}

// TmallservicecentertaskqueryrefundAPIResponseModel is 查询任务类工单是否退款 成功返回结果
type TmallservicecentertaskqueryrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_task_queryrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
