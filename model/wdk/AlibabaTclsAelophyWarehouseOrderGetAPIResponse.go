package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyWarehouseOrderGetAPIResponse 仓作业单获取 API返回值
// alibaba.tcls.aelophy.warehouse.order.get
//
// 仓作业单获取
type AlibabaTclsAelophyWarehouseOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyWarehouseOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyWarehouseOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyWarehouseOrderGetAPIResponseModel).Reset()
}

// AlibabaTclsAelophyWarehouseOrderGetAPIResponseModel is 仓作业单获取 成功返回结果
type AlibabaTclsAelophyWarehouseOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_warehouse_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *TopBaseResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyWarehouseOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyWarehouseOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyWarehouseOrderGetAPIResponse)
	},
}

// GetAlibabaTclsAelophyWarehouseOrderGetAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyWarehouseOrderGetAPIResponse
func GetAlibabaTclsAelophyWarehouseOrderGetAPIResponse() *AlibabaTclsAelophyWarehouseOrderGetAPIResponse {
	return poolAlibabaTclsAelophyWarehouseOrderGetAPIResponse.Get().(*AlibabaTclsAelophyWarehouseOrderGetAPIResponse)
}

// ReleaseAlibabaTclsAelophyWarehouseOrderGetAPIResponse 将 AlibabaTclsAelophyWarehouseOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyWarehouseOrderGetAPIResponse(v *AlibabaTclsAelophyWarehouseOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyWarehouseOrderGetAPIResponse.Put(v)
}
