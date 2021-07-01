package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeStoreItemdetailGetAPIRequest
查询商品的详情信息 API请求
alibaba.nlife.store.itemdetail.get

查询零售加平台上单个商品的详情信息 */
type AlibabaNlifeStoreItemdetailGetAPIRequest struct {
	model.Params
	// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
	_storeIdType string
	// 门店ID
	_storeId string
	// 商品在外部商家的编码(与item_id不能同时为空)
	_outerId string
	// 商品Item的ID(与outer_id不能同时为空)
	_itemId int64
	// skuId列表-可查询指定的sku
	_skuIdList []int64
	// 商品来源类型: 0-线上商品; 1-商户导入的线下商品. 如果为空则默认值为0
	_itemType *model.File
	// 商家对商品的自用编码
	_code string
}

// NewAlibabaNlifeStoreItemdetailGetRequest 初始化AlibabaNlifeStoreItemdetailGetAPIRequest对象
func NewAlibabaNlifeStoreItemdetailGetRequest() *AlibabaNlifeStoreItemdetailGetAPIRequest {
	return &AlibabaNlifeStoreItemdetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreItemdetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.store.itemdetail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreItemdetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreIdType Setter
// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
func (r *AlibabaNlifeStoreItemdetailGetAPIRequest) SetStoreIdType(_storeIdType string) error {
	r._storeIdType = _storeIdType
	r.Set("store_id_type", _storeIdType)
	return nil
}

// Get StoreIdType Getter
func (r AlibabaNlifeStoreItemdetailGetAPIRequest) GetStoreIdType() string {
	return r._storeIdType
}

// Set is StoreId Setter
// 门店ID
func (r *AlibabaNlifeStoreItemdetailGetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaNlifeStoreItemdetailGetAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is OuterId Setter
// 商品在外部商家的编码(与item_id不能同时为空)
func (r *AlibabaNlifeStoreItemdetailGetAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r AlibabaNlifeStoreItemdetailGetAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is ItemId Setter
// 商品Item的ID(与outer_id不能同时为空)
func (r *AlibabaNlifeStoreItemdetailGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r AlibabaNlifeStoreItemdetailGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is SkuIdList Setter
// skuId列表-可查询指定的sku
func (r *AlibabaNlifeStoreItemdetailGetAPIRequest) SetSkuIdList(_skuIdList []int64) error {
	r._skuIdList = _skuIdList
	r.Set("sku_id_list", _skuIdList)
	return nil
}

// Get SkuIdList Getter
func (r AlibabaNlifeStoreItemdetailGetAPIRequest) GetSkuIdList() []int64 {
	return r._skuIdList
}

// Set is ItemType Setter
// 商品来源类型: 0-线上商品; 1-商户导入的线下商品. 如果为空则默认值为0
func (r *AlibabaNlifeStoreItemdetailGetAPIRequest) SetItemType(_itemType *model.File) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// Get ItemType Getter
func (r AlibabaNlifeStoreItemdetailGetAPIRequest) GetItemType() *model.File {
	return r._itemType
}

// Set is Code Setter
// 商家对商品的自用编码
func (r *AlibabaNlifeStoreItemdetailGetAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaNlifeStoreItemdetailGetAPIRequest) GetCode() string {
	return r._code
}
