package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店关系查询 APIResponse
taobao.place.store.relation.query

查询门店关系
*/
type TaobaoPlaceStoreRelationQueryAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoreRelationQueryResponse
}

type TaobaoPlaceStoreRelationQueryResponse struct {
    XMLName xml.Name `xml:"place_store_relation_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
