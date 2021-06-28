package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询单个订阅关系 APIResponse
taobao.baichuan.item.subscribe.relation.query

查询单个订阅关系
*/
type TaobaoBaichuanItemSubscribeRelationQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaichuanItemSubscribeRelationQueryResponse `json:"baichuan_item_subscribe_relation_query_response,omitempty"` 
    TaobaoBaichuanItemSubscribeRelationQueryResponse
}

/* model for simplify = false
type TaobaoBaichuanItemSubscribeRelationQueryResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoBaichuanItemSubscribeRelationQueryResult  *TaobaoBaichuanItemSubscribeRelationQueryResult `json:"taobao_baichuan_item_subscribe_relation_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBaichuanItemSubscribeRelationQueryResponse struct {

    // 接口返回model
    Result   *TaobaoBaichuanItemSubscribeRelationQueryResult `json:"result,omitempty"`

}
