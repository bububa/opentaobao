package lifeservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店关系新增 API请求
taobao.place.store.relation.add

新增授权用户的门店关系信息
*/
type TaobaoPlaceStoreRelationAddRequest struct {
    model.Params
    // 关系B的门店ID
    _outerId   string
    // 关系类型(outer_id是主, store_id是从)
    _relationType   int64
    // 门店ID
    _storeId   int64
}

// 初始化TaobaoPlaceStoreRelationAddRequest对象
func NewTaobaoPlaceStoreRelationAddRequest() *TaobaoPlaceStoreRelationAddRequest{
    return &TaobaoPlaceStoreRelationAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreRelationAddRequest) GetApiMethodName() string {
    return "taobao.place.store.relation.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreRelationAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 关系B的门店ID
func (r *TaobaoPlaceStoreRelationAddRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoPlaceStoreRelationAddRequest) GetOuterId() string {
    return r._outerId
}
// RelationType Setter
// 关系类型(outer_id是主, store_id是从)
func (r *TaobaoPlaceStoreRelationAddRequest) SetRelationType(_relationType int64) error {
    r._relationType = _relationType
    r.Set("relation_type", _relationType)
    return nil
}

// RelationType Getter
func (r TaobaoPlaceStoreRelationAddRequest) GetRelationType() int64 {
    return r._relationType
}
// StoreId Setter
// 门店ID
func (r *TaobaoPlaceStoreRelationAddRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStoreRelationAddRequest) GetStoreId() int64 {
    return r._storeId
}
