package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseorderGetAPIResponse 获取租赁订单信息 API返回值
// tmall.car.leaseorder.get
//
// 卖家通过供销平台获取代销商的订单信息，但是部分情况下网商银行订单号获取不到，需要提供接口或者工具给卖家
type TmallCarLeaseorderGetAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseorderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeaseorderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeaseorderGetAPIResponseModel).Reset()
}

// TmallCarLeaseorderGetAPIResponseModel is 获取租赁订单信息 成功返回结果
type TmallCarLeaseorderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_leaseorder_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeaseorderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeaseorderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseorderGetAPIResponse)
	},
}

// GetTmallCarLeaseorderGetAPIResponse 从 sync.Pool 获取 TmallCarLeaseorderGetAPIResponse
func GetTmallCarLeaseorderGetAPIResponse() *TmallCarLeaseorderGetAPIResponse {
	return poolTmallCarLeaseorderGetAPIResponse.Get().(*TmallCarLeaseorderGetAPIResponse)
}

// ReleaseTmallCarLeaseorderGetAPIResponse 将 TmallCarLeaseorderGetAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeaseorderGetAPIResponse(v *TmallCarLeaseorderGetAPIResponse) {
	v.Reset()
	poolTmallCarLeaseorderGetAPIResponse.Put(v)
}
