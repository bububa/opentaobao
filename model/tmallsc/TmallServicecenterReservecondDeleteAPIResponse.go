package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterReservecondDeleteAPIResponse 删除主动预约开通条件 API返回值
// tmall.servicecenter.reservecond.delete
//
// 删除主动预约开通条件
type TmallServicecenterReservecondDeleteAPIResponse struct {
	model.CommonResponse
	TmallServicecenterReservecondDeleteAPIResponseModel
}

// TmallServicecenterReservecondDeleteAPIResponseModel is 删除主动预约开通条件 成功返回结果
type TmallServicecenterReservecondDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_reservecond_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	MsgSuccess bool `json:"msg_success,omitempty" xml:"msg_success,omitempty"`
}
