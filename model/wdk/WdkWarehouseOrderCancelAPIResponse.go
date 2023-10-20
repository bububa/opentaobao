package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkWarehouseOrderCancelAPIResponse 仓作业取消下发 API返回值
// wdk.warehouse.order.cancel
//
// 仓作业取消下发
type WdkWarehouseOrderCancelAPIResponse struct {
	model.CommonResponse
	WdkWarehouseOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *WdkWarehouseOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkWarehouseOrderCancelAPIResponseModel).Reset()
}

// WdkWarehouseOrderCancelAPIResponseModel is 仓作业取消下发 成功返回结果
type WdkWarehouseOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_warehouse_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	DataResult *DataResult `json:"data_result,omitempty" xml:"data_result,omitempty"`
}

// Reset 清空结构体
func (m *WdkWarehouseOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DataResult = nil
}

var poolWdkWarehouseOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkWarehouseOrderCancelAPIResponse)
	},
}

// GetWdkWarehouseOrderCancelAPIResponse 从 sync.Pool 获取 WdkWarehouseOrderCancelAPIResponse
func GetWdkWarehouseOrderCancelAPIResponse() *WdkWarehouseOrderCancelAPIResponse {
	return poolWdkWarehouseOrderCancelAPIResponse.Get().(*WdkWarehouseOrderCancelAPIResponse)
}

// ReleaseWdkWarehouseOrderCancelAPIResponse 将 WdkWarehouseOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseWdkWarehouseOrderCancelAPIResponse(v *WdkWarehouseOrderCancelAPIResponse) {
	v.Reset()
	poolWdkWarehouseOrderCancelAPIResponse.Put(v)
}
