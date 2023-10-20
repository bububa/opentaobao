package ticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTicketScenicBindAPIResponse 【门票API2.0】门票景点绑定接口 API返回值
// alitrip.ticket.scenic.bind
//
// 门票景点绑定接口，用于建立阿里标准景点id与商家系统景点id的映射关系。该接口同时支持新建和修改映射关系，当用户没有为ali_scenic_id建立过映射关系时，则判断为新建映射关系，否则为修改。可以通过设置update_out_scenic_id来修改ali_scenic_id与out_scenic_id的映射关系。
type AlitripTicketScenicBindAPIResponse struct {
	model.CommonResponse
	AlitripTicketScenicBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTicketScenicBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTicketScenicBindAPIResponseModel).Reset()
}

// AlitripTicketScenicBindAPIResponseModel is 【门票API2.0】门票景点绑定接口 成功返回结果
type AlitripTicketScenicBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ticket_scenic_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 景点绑定结果
	FirstResult *TicketScenicResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTicketScenicBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = nil
}

var poolAlitripTicketScenicBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTicketScenicBindAPIResponse)
	},
}

// GetAlitripTicketScenicBindAPIResponse 从 sync.Pool 获取 AlitripTicketScenicBindAPIResponse
func GetAlitripTicketScenicBindAPIResponse() *AlitripTicketScenicBindAPIResponse {
	return poolAlitripTicketScenicBindAPIResponse.Get().(*AlitripTicketScenicBindAPIResponse)
}

// ReleaseAlitripTicketScenicBindAPIResponse 将 AlitripTicketScenicBindAPIResponse 保存到 sync.Pool
func ReleaseAlitripTicketScenicBindAPIResponse(v *AlitripTicketScenicBindAPIResponse) {
	v.Reset()
	poolAlitripTicketScenicBindAPIResponse.Put(v)
}
