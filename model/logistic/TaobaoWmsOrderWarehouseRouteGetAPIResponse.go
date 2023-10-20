package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWmsOrderWarehouseRouteGetAPIResponse 获取订单仓库路由信息 API返回值
// taobao.wms.order.warehouse.route.get
//
// 获取订单仓库路由信息
type TaobaoWmsOrderWarehouseRouteGetAPIResponse struct {
	model.CommonResponse
	TaobaoWmsOrderWarehouseRouteGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWmsOrderWarehouseRouteGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWmsOrderWarehouseRouteGetAPIResponseModel).Reset()
}

// TaobaoWmsOrderWarehouseRouteGetAPIResponseModel is 获取订单仓库路由信息 成功返回结果
type TaobaoWmsOrderWarehouseRouteGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wms_order_warehouse_route_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品列表
	Items []OrderWarehouseRouteGetItems `json:"items,omitempty" xml:"items>order_warehouse_route_get_items,omitempty"`
	// 错误信息
	WlErrorCode string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`
	// 错误信息
	WlErrorMsg string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`
	// 订单编号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 是否成功
	WlSuccess bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWmsOrderWarehouseRouteGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Items = m.Items[:0]
	m.WlErrorCode = ""
	m.WlErrorMsg = ""
	m.OrderCode = ""
	m.WlSuccess = false
}

var poolTaobaoWmsOrderWarehouseRouteGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWmsOrderWarehouseRouteGetAPIResponse)
	},
}

// GetTaobaoWmsOrderWarehouseRouteGetAPIResponse 从 sync.Pool 获取 TaobaoWmsOrderWarehouseRouteGetAPIResponse
func GetTaobaoWmsOrderWarehouseRouteGetAPIResponse() *TaobaoWmsOrderWarehouseRouteGetAPIResponse {
	return poolTaobaoWmsOrderWarehouseRouteGetAPIResponse.Get().(*TaobaoWmsOrderWarehouseRouteGetAPIResponse)
}

// ReleaseTaobaoWmsOrderWarehouseRouteGetAPIResponse 将 TaobaoWmsOrderWarehouseRouteGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWmsOrderWarehouseRouteGetAPIResponse(v *TaobaoWmsOrderWarehouseRouteGetAPIResponse) {
	v.Reset()
	poolTaobaoWmsOrderWarehouseRouteGetAPIResponse.Put(v)
}
