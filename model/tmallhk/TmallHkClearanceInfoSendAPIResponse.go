package tmallhk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallhkclearanceinfosendAPIResponse 清关信息回调通知 API返回值
// tmall.hk.clearance.info.send
//
// 清关信息回调通知
type TmallhkclearanceinfosendAPIResponse struct {
	model.CommonResponse
	TmallhkclearanceinfosendAPIResponseModel
}

// TmallhkclearanceinfosendAPIResponseModel is 清关信息回调通知 成功返回结果
type TmallhkclearanceinfosendAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_hk_clearance_info_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误原因
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 操作是否成功
	MsgSuccess bool `json:"msg_success,omitempty" xml:"msg_success,omitempty"`
	// 接口调用是否成功
	MsgObj bool `json:"msg_obj,omitempty" xml:"msg_obj,omitempty"`
}
