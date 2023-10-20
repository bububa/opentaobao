package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseOrderidGetAPIResponse 天猫开新车查询订单id API返回值
// tmall.car.lease.orderid.get
//
// 天猫开新车查询订单id
type TmallCarLeaseOrderidGetAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseOrderidGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeaseOrderidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeaseOrderidGetAPIResponseModel).Reset()
}

// TmallCarLeaseOrderidGetAPIResponseModel is 天猫开新车查询订单id 成功返回结果
type TmallCarLeaseOrderidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_orderid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeaseOrderidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeaseOrderidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseOrderidGetAPIResponse)
	},
}

// GetTmallCarLeaseOrderidGetAPIResponse 从 sync.Pool 获取 TmallCarLeaseOrderidGetAPIResponse
func GetTmallCarLeaseOrderidGetAPIResponse() *TmallCarLeaseOrderidGetAPIResponse {
	return poolTmallCarLeaseOrderidGetAPIResponse.Get().(*TmallCarLeaseOrderidGetAPIResponse)
}

// ReleaseTmallCarLeaseOrderidGetAPIResponse 将 TmallCarLeaseOrderidGetAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeaseOrderidGetAPIResponse(v *TmallCarLeaseOrderidGetAPIResponse) {
	v.Reset()
	poolTmallCarLeaseOrderidGetAPIResponse.Put(v)
}
