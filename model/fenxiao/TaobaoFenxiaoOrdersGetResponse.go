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
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_orders_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 搜索到的采购单记录总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"