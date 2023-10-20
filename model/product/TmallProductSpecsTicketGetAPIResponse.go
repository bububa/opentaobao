package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSpecsTicketGetAPIResponse 产品规格审核信息获取接口 API返回值
// tmall.product.specs.ticket.get
//
// 批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等
type TmallProductSpecsTicketGetAPIResponse struct {
	model.CommonResponse
	TmallProductSpecsTicketGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallProductSpecsTicketGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallProductSpecsTicketGetAPIResponseModel).Reset()
}

// TmallProductSpecsTicketGetAPIResponseModel is 产品规格审核信息获取接口 成功返回结果
type TmallProductSpecsTicketGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_specs_ticket_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品规格审核单信息
	Tickets []Ticket `json:"tickets,omitempty" xml:"tickets>ticket,omitempty"`
}

// Reset 清空结构体
func (m *TmallProductSpecsTicketGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Tickets = m.Tickets[:0]
}

var poolTmallProductSpecsTicketGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallProductSpecsTicketGetAPIResponse)
	},
}

// GetTmallProductSpecsTicketGetAPIResponse 从 sync.Pool 获取 TmallProductSpecsTicketGetAPIResponse
func GetTmallProductSpecsTicketGetAPIResponse() *TmallProductSpecsTicketGetAPIResponse {
	return poolTmallProductSpecsTicketGetAPIResponse.Get().(*TmallProductSpecsTicketGetAPIResponse)
}

// ReleaseTmallProductSpecsTicketGetAPIResponse 将 TmallProductSpecsTicketGetAPIResponse 保存到 sync.Pool
func ReleaseTmallProductSpecsTicketGetAPIResponse(v *TmallProductSpecsTicketGetAPIResponse) {
	v.Reset()
	poolTmallProductSpecsTicketGetAPIResponse.Put(v)
}
