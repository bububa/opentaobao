package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询后端商品 APIResponse
taobao.scitem.query

查询后端商品
*/
type TaobaoScitemQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoScitemQueryResponse `json:"scitem_query_response,omitempty"` 
    TaobaoScitemQueryResponse
}

/* model for simplify = false
type TaobaoScitemQueryResponse struct {

    // List<ScItemDO>
    
    ScItemList  struct {
        ScItem  []ScItem `json:"sc_item,omitempty"`
    } `json:"sc_item_list,omitempty"`
    

    // 商品条数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // 分页
    
    QueryPagination  *struct {
        QueryPagination  *QueryPagination `json:"query_pagination,omitempty"`
    } `json:"query_pagination,omitempty"`
    

}
*/

type TaobaoScitemQueryResponse struct {

    // List<ScItemDO>
    ScItemList   []ScItem `json:"sc_item_list,omitempty"`

    // 商品条数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 分页
    QueryPagination   *QueryPagination `json:"query_pagination,omitempty"`

}
