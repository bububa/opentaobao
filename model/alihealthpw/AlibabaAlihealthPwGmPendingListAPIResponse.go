package alihealthpw

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwgmpendinglistAPIResponse 同情用药待审核工单查询接口 API返回值
// alibaba.alihealth.pw.gm.pending.list
//
// 同情用药待审核工单查询接口，提供给合作方用来查询待处理工单列表
type AlibabaalihealthpwgmpendinglistAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthpwgmpendinglistAPIResponseModel
}

// AlibabaalihealthpwgmpendinglistAPIResponseModel is 同情用药待审核工单查询接口 成功返回结果
type AlibabaalihealthpwgmpendinglistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_gm_pending_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}
