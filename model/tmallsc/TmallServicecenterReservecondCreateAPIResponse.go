package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterReservecondCreateAPIResponse 创建主动预约开通条件 API返回值
// tmall.servicecenter.reservecond.create
//
// 1、设置主动预约开通条件
type TmallServicecenterReservecondCreateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterReservecondCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterReservecondCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterReservecondCreateAPIResponseModel).Reset()
}

// TmallServicecenterReservecondCreateAPIResponseModel is 创建主动预约开通条件 成功返回结果
type TmallServicecenterReservecondCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_reservecond_create_response"`
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
func (m *TmallServicecenterReservecondCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.MsgSuccess = false
}

var poolTmallServicecenterReservecondCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterReservecondCreateAPIResponse)
	},
}

// GetTmallServicecenterReservecondCreateAPIResponse 从 sync.Pool 获取 TmallServicecenterReservecondCreateAPIResponse
func GetTmallServicecenterReservecondCreateAPIResponse() *TmallServicecenterReservecondCreateAPIResponse {
	return poolTmallServicecenterReservecondCreateAPIResponse.Get().(*TmallServicecenterReservecondCreateAPIResponse)
}

// ReleaseTmallServicecenterReservecondCreateAPIResponse 将 TmallServicecenterReservecondCreateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterReservecondCreateAPIResponse(v *TmallServicecenterReservecondCreateAPIResponse) {
	v.Reset()
	poolTmallServicecenterReservecondCreateAPIResponse.Put(v)
}
