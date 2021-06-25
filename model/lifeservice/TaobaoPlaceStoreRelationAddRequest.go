package lifeservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店关系新增 APIRequest
taobao.place.store.relation.add

新增授权用户的门店关系信息
*/
type TaobaoPlaceStoreRelationAddRequest struct {
    model.Params

    // 关系B的门店ID
    outerId   string 

    // 关系类型(outer_id是主, store_id是从)
    relationType   int64 

    // 门店ID
    storeId   int64 

}

func NewTaobaoPlaceStoreRelationAddRequest() *TaobaoPlaceStoreRelationAddRequest{
    return &TaobaoPlaceStoreRelationAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoreRelationAddRequest) GetApiMethodName() string {
    return "taobao.place.store.relation.add"
}

func (r TaobaoPlaceStoreRelationAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoreRelationAddRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoPlaceStoreRelationAddRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoPlaceStoreRelationAddRequest) SetRelationType(relationType int64) error {
    r.relationType = relationType
    r.Set("relation_type", relationType)
    return nil
}

func (r TaobaoPlaceStoreRelationAddRequest) GetRelationType() int64 {
    return r.relationType
}

func (r *TaobaoPlaceStoreRelationAddRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoPlaceStoreRelationAddRequest) GetStoreId() int64 {
    return r.storeId
}

