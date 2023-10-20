package ticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTicketProductQueryAPIResponse 【门票API2.0】门票商品查询接口 API返回值
// alitrip.ticket.product.query
//
// 门票商品查询接口：返回商家上传的门票商品信息
type AlitripTicketProductQueryAPIResponse struct {
	model.CommonResponse
	AlitripTicketProductQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTicketProductQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTicketProductQueryAPIResponseModel).Reset()
}

// AlitripTicketProductQueryAPIResponseModel is 【门票API2.0】门票商品查询接口 成功返回结果
type AlitripTicketProductQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ticket_product_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门票商品详情
	FirstResult *TopTicketItemFullinfoResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTicketProductQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = nil
}

var poolAlitripTicketProductQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTicketProductQueryAPIResponse)
	},
}

// GetAlitripTicketProductQueryAPIResponse 从 sync.Pool 获取 AlitripTicketProductQueryAPIResponse
func GetAlitripTicketProductQueryAPIResponse() *AlitripTicketProductQueryAPIResponse {
	return poolAlitripTicketProductQueryAPIResponse.Get().(*AlitripTicketProductQueryAPIResponse)
}

// ReleaseAlitripTicketProductQueryAPIResponse 将 AlitripTicketProductQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripTicketProductQueryAPIResponse(v *AlitripTicketProductQueryAPIResponse) {
	v.Reset()
	poolAlitripTicketProductQueryAPIResponse.Put(v)
}
