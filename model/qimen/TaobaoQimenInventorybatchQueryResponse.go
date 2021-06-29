package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品单仓批次库存查询接口 APIResponse
taobao.qimen.inventorybatch.query

ERP 通过该接口查询指定商品的单仓批次库存
*/
type TaobaoQimenInventorybatchQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenInventorybatchQueryResponse
}

type TaobaoQimenInventorybatchQueryResponse struct {
    XMLName xml.Name `xml:"qimen_inventorybatch_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
