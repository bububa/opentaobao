package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianjiSupplierOrderQueryAPIResponse 查询供应商订单 API返回值
// alibaba.tianji.supplier.order.query
//
// 查询供应商订单
type AlibabaTianjiSupplierOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaTianjiSupplierOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTianjiSupplierOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTianjiSupplierOrderQueryAPIResponseModel).Reset()
}

// AlibabaTianjiSupplierOrderQueryAPIResponseModel is 查询供应商订单 成功返回结果
type AlibabaTianjiSupplierOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianji_supplier_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分销订单信息
	ModelList []DistributionOrderInfo `json:"model_list,omitempty" xml:"model_list>distribution_order_info,omitempty"`
	// 查询总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTianjiSupplierOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ModelList = m.ModelList[:0]
	m.TotalCount = 0
}

var poolAlibabaTianjiSupplierOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTianjiSupplierOrderQueryAPIResponse)
	},
}

// GetAlibabaTianjiSupplierOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaTianjiSupplierOrderQueryAPIResponse
func GetAlibabaTianjiSupplierOrderQueryAPIResponse() *AlibabaTianjiSupplierOrderQueryAPIResponse {
	return poolAlibabaTianjiSupplierOrderQueryAPIResponse.Get().(*AlibabaTianjiSupplierOrderQueryAPIResponse)
}

// ReleaseAlibabaTianjiSupplierOrderQueryAPIResponse 将 AlibabaTianjiSupplierOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTianjiSupplierOrderQueryAPIResponse(v *AlibabaTianjiSupplierOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaTianjiSupplierOrderQueryAPIResponse.Put(v)
}
