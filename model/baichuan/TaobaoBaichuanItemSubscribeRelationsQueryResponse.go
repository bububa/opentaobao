package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
按条件查询订阅关系 APIResponse
taobao.baichuan.item.subscribe.relations.query

按条件查询订阅关系
*/
type TaobaoBaichuanItemSubscribeRelationsQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanItemSubscribeRelationsQueryResponse `json:"taobao_baichuan_item_subscribe_relations_query_response,omitempty"`
}

type TaobaoBaichuanItemSubscribeRelationsQueryResponse struct {

    // 接口返回model
    Result   *TaobaoBaichuanItemSubscribeRelationsQueryResult `json:"result,omitempty"`

}
