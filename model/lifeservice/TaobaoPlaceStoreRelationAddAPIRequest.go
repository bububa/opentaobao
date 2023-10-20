package lifeservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestorerelationaddAPIRequest 门店关系新增 API请求
// taobao.place.store.relation.add
//
// 新增授权用户的门店关系信息
type TaobaoplacestorerelationaddAPIRequest struct {
	model.Params
	// 关系B的门店ID
	_outerId string
	// 关系类型(outer_id是主, store_id是从)
	_relationType int64
	// 门店ID
	_storeId int64
}

// NewTaobaoplacestorerelationaddRequest 初始化TaobaoplacestorerelationaddAPIRequest对象
func NewTaobaoplacestorerelationaddRequest() *TaobaoplacestorerelationaddAPIRequest {
	return &TaobaoplacestorerelationaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestorerelationaddAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.relation.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestorerelationaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestorerelationaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 关系B的门店ID
func (r *TaobaoplacestorerelationaddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoplacestorerelationaddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetRelationType is RelationType Setter
// 关系类型(outer_id是主, store_id是从)
func (r *TaobaoplacestorerelationaddAPIRequest) SetRelationType(_relationType int64) error {
	r._relationType = _relationType
	r.Set("relation_type", _relationType)
	return nil
}

// GetRelationType RelationType Getter
func (r TaobaoplacestorerelationaddAPIRequest) GetRelationType() int64 {
	return r._relationType
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoplacestorerelationaddAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoplacestorerelationaddAPIRequest) GetStoreId() int64 {
	return r._storeId
}
