package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterReservecondUpdateAPIResponse 主动预约条件更新 API返回值
// tmall.servicecenter.reservecond.update
//
// 1、设置主动预约开通条件
type TmallServicecenterReservecondUpdateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterReservecondUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterReservecondUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterReservecondUpdateAPIResponseModel).Reset()
}

// TmallServicecenterReservecondUpdateAPIResponseModel is 主动预约条件更新 成功返回结果
type TmallServicecenterReservecondUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_reservecond_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	MsgSuccess bool `json:"msg_success,omitempty" xml:"msg_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterReservecondUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.MsgSuccess = false
}

var poolTmallServicecenterReservecondUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterReservecondUpdateAPIResponse)
	},
}

// GetTmallServicecenterReservecondUpdateAPIResponse 从 sync.Pool 获取 TmallServicecenterReservecondUpdateAPIResponse
func GetTmallServicecenterReservecondUpdateAPIResponse() *TmallServicecenterReservecondUpdateAPIResponse {
	return poolTmallServicecenterReservecondUpdateAPIResponse.Get().(*TmallServicecenterReservecondUpdateAPIResponse)
}

// ReleaseTmallServicecenterReservecondUpdateAPIResponse 将 TmallServicecenterReservecondUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterReservecondUpdateAPIResponse(v *TmallServicecenterReservecondUpdateAPIResponse) {
	v.Reset()
	poolTmallServicecenterReservecondUpdateAPIResponse.Put(v)
}
