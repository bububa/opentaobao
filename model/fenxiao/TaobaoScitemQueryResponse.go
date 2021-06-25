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
    Response *TaobaoScitemQueryResponse `json:"taobao_scitem_query_response,omitempty"`
}

type TaobaoScitemQueryResponse struct {

    // List<ScItemDO>
    ScItemList   []ScItem `json:"sc_item_list,omitempty"`

    // 商品条数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 分页
    QueryPagination   *QueryPagination `json:"query_pagination,omitempty"`

}
