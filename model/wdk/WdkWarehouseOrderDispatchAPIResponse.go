package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// WdkWarehouseOrderDispatchAPIResponse 仓作业下发 API返回值
// wdk.warehouse.order.dispatch
//
// 牵牛花仓作业下发接口提供
type WdkWarehouseOrderDispatchAPIResponse struct {
	model.CommonResponse
	WdkWarehouseOrderDispatchAPIResponseModel
}

// WdkWarehouseOrderDispatchAPIResponseModel is 仓作业下发 成功返回结果
type WdkWarehouseOrderDispatchAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_warehouse_order_dispatch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	BaseResult *BaseResult `json:"base_result,omitempty" xml:"base_result,omitempty"`
}
