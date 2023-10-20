package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest 门店委托信息查询 API请求
// alibaba.alihouse.existinghome.house.storeentrust.query
//
// 门店委托信息查询
type AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest struct {
	model.Params
	// 外部小区id
	_outerCommunityId string
	// 外部房源id
	_outerHouseId string
	// 门店id
	_storeId int64
}

// NewAlibabaalihouseexistinghomehousestoreentrustqueryRequest 初始化AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest对象
func NewAlibabaalihouseexistinghomehousestoreentrustqueryRequest() *AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest {
	return &AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.storeentrust.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterCommunityId is OuterCommunityId Setter
// 外部小区id
func (r *AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest) SetOuterCommunityId(_outerCommunityId string) error {
	r._outerCommunityId = _outerCommunityId
	r.Set("outer_community_id", _outerCommunityId)
	return nil
}

// GetOuterCommunityId OuterCommunityId Getter
func (r AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest) GetOuterCommunityId() string {
	return r._outerCommunityId
}

// SetOuterHouseId is OuterHouseId Setter
// 外部房源id
func (r *AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest) SetOuterHouseId(_outerHouseId string) error {
	r._outerHouseId = _outerHouseId
	r.Set("outer_house_id", _outerHouseId)
	return nil
}

// GetOuterHouseId OuterHouseId Getter
func (r AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest) GetOuterHouseId() string {
	return r._outerHouseId
}

// SetStoreId is StoreId Setter
// 门店id
func (r *AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest) GetStoreId() int64 {
	return r._storeId
}
