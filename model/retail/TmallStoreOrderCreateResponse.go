package retail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店订单创建api APIResponse
tmall.store.order.create

门店订单创建api
*/
type TmallStoreOrderCreateAPIResponse struct {
    model.CommonResponse
    TmallStoreOrderCreateResponse
}

type TmallStoreOrderCreateResponse struct {
    XMLName xml.Name `xml:"tmall_store_order_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // detailResults
    
    DetailResults   []Detailresults `json:"detail_results,omitempty" xml:"detail_results>detailresults,omitempty"`
    
    
}
