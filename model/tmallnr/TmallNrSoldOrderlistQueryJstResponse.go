package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
已入塔商家查询订单列表 APIResponse
tmall.nr.sold.orderlist.query.jst

该服务用于已入聚石塔的商家，获取订单列表信息；
*/
type TmallNrSoldOrderlistQueryJstAPIResponse struct {
    model.CommonResponse
    TmallNrSoldOrderlistQueryJstResponse
}

type TmallNrSoldOrderlistQueryJstResponse struct {
    XMLName xml.Name `xml:"tmall_nr_sold_orderlist_query_jst_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *NewRetailResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
