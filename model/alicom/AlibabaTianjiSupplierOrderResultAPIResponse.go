package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianjiSupplierOrderResultAPIResponse 供应商处理订单接口（订购成功/失败、发货） API返回值
// alibaba.tianji.supplier.order.result
//
// 供应商处理订单接口（订购成功/失败、发货）
type AlibabaTianjiSupplierOrderResultAPIResponse struct {
	model.CommonResponse
	AlibabaTianjiSupplierOrderResultAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTianjiSupplierOrderResultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTianjiSupplierOrderResultAPIResponseModel).Reset()
}

// AlibabaTianjiSupplierOrderResultAPIResponseModel is 供应商处理订单接口（订购成功/失败、发货） 成功返回结果
type AlibabaTianjiSupplierOrderResultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianji_supplier_order_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTianjiSupplierOrderResultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = false
}

var poolAlibabaTianjiSupplierOrderResultAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTianjiSupplierOrderResultAPIResponse)
	},
}

// GetAlibabaTianjiSupplierOrderResultAPIResponse 从 sync.Pool 获取 AlibabaTianjiSupplierOrderResultAPIResponse
func GetAlibabaTianjiSupplierOrderResultAPIResponse() *AlibabaTianjiSupplierOrderResultAPIResponse {
	return poolAlibabaTianjiSupplierOrderResultAPIResponse.Get().(*AlibabaTianjiSupplierOrderResultAPIResponse)
}

// ReleaseAlibabaTianjiSupplierOrderResultAPIResponse 将 AlibabaTianjiSupplierOrderResultAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTianjiSupplierOrderResultAPIResponse(v *AlibabaTianjiSupplierOrderResultAPIResponse) {
	v.Reset()
	poolAlibabaTianjiSupplierOrderResultAPIResponse.Put(v)
}
