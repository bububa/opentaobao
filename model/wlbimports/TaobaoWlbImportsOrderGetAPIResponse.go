package wlbimports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流订单获取 API返回值 
taobao.wlb.imports.order.get

一般进口物流订单获取
*/
type TaobaoWlbImportsOrderGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbImportsOrderGetAPIResponseModel
}

// 物流订单获取 成功返回结果
type TaobaoWlbImportsOrderGetAPIResponseModel struct {
    XMLName xml.Name `xml:"wlb_imports_order_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 物流订单信息
    Orders   []LocOrder `json:"orders,omitempty" xml:"orders>loc_order,omitempty"`
    // 搜索到的总数量
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
