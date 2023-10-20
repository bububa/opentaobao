package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsordersgetAPIResponse 批量查询物流订单 API返回值
// taobao.logistics.orders.get
//
// 批量查询物流订单。
type TaobaologisticsordersgetAPIResponse struct {
	model.CommonResponse
	TaobaologisticsordersgetAPIResponseModel
}

// TaobaologisticsordersgetAPIResponseModel is 批量查询物流订单 成功返回结果
type TaobaologisticsordersgetAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_orders_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取的物流订单详情列表 &lt;br/&gt;返回的Shipping包含的具体信息为入参fields请求的字段信息
	Shippings []Shipping `json:"shippings,omitempty" xml:"shippings>shipping,omitempty"`
	// 搜索到的物流订单列表总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
