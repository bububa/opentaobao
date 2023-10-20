package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardassignworkerAPIResponse 服务商分派工人 API返回值
// tmall.servicecenter.workcard.assignworker
//
// 服务商调用该接口分派工人给具体的工单
type TmallservicecenterworkcardassignworkerAPIResponse struct {
	model.CommonResponse
	TmallservicecenterworkcardassignworkerAPIResponseModel
}

// TmallservicecenterworkcardassignworkerAPIResponseModel is 服务商分派工人 成功返回结果
type TmallservicecenterworkcardassignworkerAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_assignworker_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// -
	Result *TmallservicecenterworkcardassignworkerResult `json:"result,omitempty" xml:"result,omitempty"`
}
