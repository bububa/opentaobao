package tmallhk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallHkClearanceInfoSendAPIResponse 清关信息回调通知 API返回值
// tmall.hk.clearance.info.send
//
// 清关信息回调通知
type TmallHkClearanceInfoSendAPIResponse struct {
	model.CommonResponse
	TmallHkClearanceInfoSendAPIResponseModel
}

// Reset 清空结构体
func (m *TmallHkClearanceInfoSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallHkClearanceInfoSendAPIResponseModel).Reset()
}

// TmallHkClearanceInfoSendAPIResponseModel is 清关信息回调通知 成功返回结果
type TmallHkClearanceInfoSendAPIResponseModel struct {
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

// Reset 清空结构体
func (m *TmallHkClearanceInfoSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.MsgSuccess = false
	m.MsgObj = false
}

var poolTmallHkClearanceInfoSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallHkClearanceInfoSendAPIResponse)
	},
}

// GetTmallHkClearanceInfoSendAPIResponse 从 sync.Pool 获取 TmallHkClearanceInfoSendAPIResponse
func GetTmallHkClearanceInfoSendAPIResponse() *TmallHkClearanceInfoSendAPIResponse {
	return poolTmallHkClearanceInfoSendAPIResponse.Get().(*TmallHkClearanceInfoSendAPIResponse)
}

// ReleaseTmallHkClearanceInfoSendAPIResponse 将 TmallHkClearanceInfoSendAPIResponse 保存到 sync.Pool
func ReleaseTmallHkClearanceInfoSendAPIResponse(v *TmallHkClearanceInfoSendAPIResponse) {
	v.Reset()
	poolTmallHkClearanceInfoSendAPIResponse.Put(v)
}
