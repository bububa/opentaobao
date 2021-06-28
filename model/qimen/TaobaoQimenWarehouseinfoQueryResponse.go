package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
货主仓库资源查询接口 APIResponse
taobao.qimen.warehouseinfo.query

货主仓库资源查询
*/
type TaobaoQimenWarehouseinfoQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenWarehouseinfoQueryResponse `json:"qimen_warehouseinfo_query_response,omitempty"` 
    TaobaoQimenWarehouseinfoQueryResponse
}

/* model for simplify = false
type TaobaoQimenWarehouseinfoQueryResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenWarehouseinfoQueryResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
