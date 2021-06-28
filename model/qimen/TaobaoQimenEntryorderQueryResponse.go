package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
入库单查询接口 APIResponse
taobao.qimen.entryorder.query

ERP调用接口，查询入库单信息;
*/
type TaobaoQimenEntryorderQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenEntryorderQueryResponse `json:"qimen_entryorder_query_response,omitempty"` 
    TaobaoQimenEntryorderQueryResponse
}

/* model for simplify = false
type TaobaoQimenEntryorderQueryResponse struct {

    // 
    
    Response  *struct {
        EntryOrderQueryResponse  *EntryOrderQueryResponse `json:"entry_order_query_response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenEntryorderQueryResponse struct {

    // 
    Response   *EntryOrderQueryResponse `json:"response,omitempty"`

}
