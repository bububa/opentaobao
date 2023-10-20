package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardTpconfirmAPIResponse 确认服务完成 API返回值
// tmall.servicecenter.workcard.tpconfirm
//
// 服务商确认服务完成
type TmallServicecenterWorkcardTpconfirmAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardTpconfirmAPIResponseModel
}

// TmallServicecenterWorkcardTpconfirmAPIResponseModel is 确认服务完成 成功返回结果
type TmallServicecenterWorkcardTpconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_tpconfirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
