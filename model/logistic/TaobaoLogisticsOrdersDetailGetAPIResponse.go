package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询物流订单,返回详细信息 API返回值 
taobao.logistics.orders.detail.get

查询物流订单的详细信息，涉及用户隐私字段。
*/
type TaobaoLogisticsOrdersDetailGetAPIResponse struct {
    model.CommonResponse
    TaobaoLogisticsOrdersDetailGetAPIResponseModel
}

// 批量查询物流订单,返回详细信息 成功返回结果
type TaobaoLogisticsOrdersDetailGetAPIResponseModel struct {
    XMLName xml.Name `xml:"logistics_orders_detail_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 获取的物流订单详情列表.返回的Shipping包含的具体信息为入参fields请求的字段信息.
    Shippings   []Shipping `json:"shippings,omitempty" xml:"shippings>shipping,omitempty"`
    // 搜索到的物流订单列表总数
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
