package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardsuspendAPIResponse 工单挂起 API返回值
// tmall.servicecenter.workcard.suspend
//
// 工单挂起
type TmallservicecenterworkcardsuspendAPIResponse struct {
	model.CommonResponse
	TmallservicecenterworkcardsuspendAPIResponseModel
}

// TmallservicecenterworkcardsuspendAPIResponseModel is 工单挂起 成功返回结果
type TmallservicecenterworkcardsuspendAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_suspend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *TmallservicecenterworkcardsuspendResult `json:"result,omitempty" xml:"result,omitempty"`
}
