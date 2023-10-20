package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripRentcarOrderDetailQueryAPIResponse 租车订单详情查询 API返回值
// alitrip.rentcar.order.detail.query
//
// 租车订单详情查询
type AlitripRentcarOrderDetailQueryAPIResponse struct {
	model.CommonResponse
	AlitripRentcarOrderDetailQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripRentcarOrderDetailQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripRentcarOrderDetailQueryAPIResponseModel).Reset()
}

// AlitripRentcarOrderDetailQueryAPIResponseModel is 租车订单详情查询 成功返回结果
type AlitripRentcarOrderDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_rentcar_order_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RentCarOrderDetailRsp `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripRentcarOrderDetailQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripRentcarOrderDetailQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripRentcarOrderDetailQueryAPIResponse)
	},
}

// GetAlitripRentcarOrderDetailQueryAPIResponse 从 sync.Pool 获取 AlitripRentcarOrderDetailQueryAPIResponse
func GetAlitripRentcarOrderDetailQueryAPIResponse() *AlitripRentcarOrderDetailQueryAPIResponse {
	return poolAlitripRentcarOrderDetailQueryAPIResponse.Get().(*AlitripRentcarOrderDetailQueryAPIResponse)
}

// ReleaseAlitripRentcarOrderDetailQueryAPIResponse 将 AlitripRentcarOrderDetailQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripRentcarOrderDetailQueryAPIResponse(v *AlitripRentcarOrderDetailQueryAPIResponse) {
	v.Reset()
	poolAlitripRentcarOrderDetailQueryAPIResponse.Put(v)
}
