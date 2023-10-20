package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTicketSetAPIResponse 出票接口 API返回值
// taobao.bus.ticket.set
//
// 提供给汽车票商家出票使用
type TaobaoBusTicketSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusTicketSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusTicketSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusTicketSetAPIResponseModel).Reset()
}

// TaobaoBusTicketSetAPIResponseModel is 出票接口 成功返回结果
type TaobaoBusTicketSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_ticket_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	ErrorCode1 string `json:"error_code1,omitempty" xml:"error_code1,omitempty"`
	// errorMsg
	ErrorMsg1 string `json:"error_msg1,omitempty" xml:"error_msg1,omitempty"`
	// success1
	Success1 bool `json:"success1,omitempty" xml:"success1,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusTicketSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorCode1 = ""
	m.ErrorMsg1 = ""
	m.Success1 = false
}

var poolTaobaoBusTicketSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusTicketSetAPIResponse)
	},
}

// GetTaobaoBusTicketSetAPIResponse 从 sync.Pool 获取 TaobaoBusTicketSetAPIResponse
func GetTaobaoBusTicketSetAPIResponse() *TaobaoBusTicketSetAPIResponse {
	return poolTaobaoBusTicketSetAPIResponse.Get().(*TaobaoBusTicketSetAPIResponse)
}

// ReleaseTaobaoBusTicketSetAPIResponse 将 TaobaoBusTicketSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusTicketSetAPIResponse(v *TaobaoBusTicketSetAPIResponse) {
	v.Reset()
	poolTaobaoBusTicketSetAPIResponse.Put(v)
}
