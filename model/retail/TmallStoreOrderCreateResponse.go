package retail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店订单创建api API返回值 
tmall.store.order.create

门店订单创建api
*/
type TmallStoreOrderCreateAPIResponse struct {
    model.CommonResponse
    TmallStoreOrderCreateResponse
}

// 门店订单创建api 成功返回结果
type TmallStoreOrderCreateResponse struct {
    XMLName xml.Name `xml:"tmall_store_order_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // detailResults
    DetailResults   []Detailresults `json:"detail_results,omitempty" xml:"detail_results>detailresults,omitempty"`
}
