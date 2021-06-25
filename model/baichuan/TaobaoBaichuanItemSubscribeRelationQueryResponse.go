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
    Response *TaobaoBaichuanItemSubscribeRelationQueryResponse `json:"taobao_baichuan_item_subscribe_relation_query_response,omitempty"`
}

type TaobaoBaichuanItemSubscribeRelationQueryResponse struct {

    // 接口返回model
    Result   *TaobaoBaichuanItemSubscribeRelationQueryResult `json:"result,omitempty"`

}