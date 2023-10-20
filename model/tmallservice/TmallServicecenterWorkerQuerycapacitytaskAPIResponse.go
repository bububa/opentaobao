package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkerquerycapacitytaskAPIResponse 查询需求容量 API返回值
// tmall.servicecenter.worker.querycapacitytask
//
// 查询需求容量
type TmallservicecenterworkerquerycapacitytaskAPIResponse struct {
	model.CommonResponse
	TmallservicecenterworkerquerycapacitytaskAPIResponseModel
}

// TmallservicecenterworkerquerycapacitytaskAPIResponseModel is 查询需求容量 成功返回结果
type TmallservicecenterworkerquerycapacitytaskAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_worker_querycapacitytask_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ResultBase
	ResultBase *ResultBase `json:"result_base,omitempty" xml:"result_base,omitempty"`
}
