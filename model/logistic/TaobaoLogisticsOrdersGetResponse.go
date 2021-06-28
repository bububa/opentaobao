package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询物流订单 APIResponse
taobao.logistics.orders.get

批量查询物流订单。
*/
type TaobaoLogisticsOrdersGetAPIResponse struct {
    model.CommonResponse
    TaobaoLogisticsOrdersGetResponse
}

type TaobaoLogisticsOrdersGetResponse struct {
    XMLName xml.Name `xml:"logistics_orders_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 获取的物流订单详情列表 <br/>返回的Shipping包含的具体信息为入参fields请求的字段信息
    
    Shippings   []Shipping `json:"shippings,omitempty" xml:"shippings>shipping,omitempty"`
    
    
    // 搜索到的物流订单列表总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
}
