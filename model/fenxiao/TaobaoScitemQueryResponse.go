package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询后端商品 APIResponse
taobao.scitem.query

查询后端商品
*/
type TaobaoScitemQueryAPIResponse struct {
    model.CommonResponse
    TaobaoScitemQueryResponse
}

type TaobaoScitemQueryResponse struct {
    XMLName xml.Name `xml:"scitem_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // List<ScItemDO>
    
    ScItemList   []ScItem `json:"sc_item_list,omitempty" xml:"sc_item_list>sc_item,omitempty"`
    
    
    // 商品条数
    
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`

    
    // 分页
    
    QueryPagination   *QueryPagination `json:"query_pagination,omitempty" xml:"query_pagination,omitempty"`

    
}
