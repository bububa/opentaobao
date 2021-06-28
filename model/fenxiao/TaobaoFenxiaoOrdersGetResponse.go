package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询采购单信息 APIResponse
taobao.fenxiao.orders.get

分销商或供应商均可用此接口查询采购单信息（代销）； (发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。
*/
type TaobaoFenxiaoOrdersGetAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoOrdersGetResponse
}

type TaobaoFenxiaoOrdersGetResponse struct {
    XMLName xml.Name `xml:"fenxiao_orders_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 搜索到的采购单记录总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
    // 采购单及子采购单信息。返回 PurchaseOrder 包含的字段信息。
    
    PurchaseOrders   []TopDpOrderDo `json:"purchase_orders,omitempty" xml:"purchase_orders>top_dp_order_do,omitempty"`
    
    
}
