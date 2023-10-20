package ticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTicketScenicQueryAPIResponse 【门票API2.0】卖家已发布门票商品列表查询接口（根据景点维度查询） API返回值
// alitrip.ticket.scenic.query
//
// 查询卖家已发布过的门票商品列表，根据景点维度聚合查询。如果卖家在该景点下未发布过任何商品，则查询不到数据！
type AlitripTicketScenicQueryAPIResponse struct {
	model.CommonResponse
	AlitripTicketScenicQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTicketScenicQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTicketScenicQueryAPIResponseModel).Reset()
}

// AlitripTicketScenicQueryAPIResponseModel is 【门票API2.0】卖家已发布门票商品列表查询接口（根据景点维度查询） 成功返回结果
type AlitripTicketScenicQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ticket_scenic_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	FirstResult *ScenicAndProductResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTicketScenicQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = nil
}

var poolAlitripTicketScenicQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTicketScenicQueryAPIResponse)
	},
}

// GetAlitripTicketScenicQueryAPIResponse 从 sync.Pool 获取 AlitripTicketScenicQueryAPIResponse
func GetAlitripTicketScenicQueryAPIResponse() *AlitripTicketScenicQueryAPIResponse {
	return poolAlitripTicketScenicQueryAPIResponse.Get().(*AlitripTicketScenicQueryAPIResponse)
}

// ReleaseAlitripTicketScenicQueryAPIResponse 将 AlitripTicketScenicQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripTicketScenicQueryAPIResponse(v *AlitripTicketScenicQueryAPIResponse) {
	v.Reset()
	poolAlitripTicketScenicQueryAPIResponse.Put(v)
}
