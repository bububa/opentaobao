package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTuanHotelItemSkuDeleteAPIRequest
酒店团购套餐商品SKU删除 API请求
alitrip.tuan.hotel.item.sku.delete

商户对发布的宝贝套餐价格库存信息进行删除 */
type AlitripTuanHotelItemSkuDeleteAPIRequest struct {
	model.Params
	// 宝贝ID
	_itemId int64
	// 宝贝所属类目
	_catId int64
	// 要删除的skuId列表
	_itemDeletedSkuIdList []int64
}

// NewAlitripTuanHotelItemSkuDeleteRequest 初始化AlitripTuanHotelItemSkuDeleteAPIRequest对象
func NewAlitripTuanHotelItemSkuDeleteRequest() *AlitripTuanHotelItemSkuDeleteAPIRequest {
	return &AlitripTuanHotelItemSkuDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelItemSkuDeleteAPIRequest) GetApiMethodName() string {
	return "alitrip.tuan.hotel.item.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelItemSkuDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 宝贝ID
func (r *AlitripTuanHotelItemSkuDeleteAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r AlitripTuanHotelItemSkuDeleteAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is CatId Setter
// 宝贝所属类目
func (r *AlitripTuanHotelItemSkuDeleteAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// Get CatId Getter
func (r AlitripTuanHotelItemSkuDeleteAPIRequest) GetCatId() int64 {
	return r._catId
}

// Set is ItemDeletedSkuIdList Setter
// 要删除的skuId列表
func (r *AlitripTuanHotelItemSkuDeleteAPIRequest) SetItemDeletedSkuIdList(_itemDeletedSkuIdList []int64) error {
	r._itemDeletedSkuIdList = _itemDeletedSkuIdList
	r.Set("item_deleted_sku_id_list", _itemDeletedSkuIdList)
	return nil
}

// Get ItemDeletedSkuIdList Getter
func (r AlitripTuanHotelItemSkuDeleteAPIRequest) GetItemDeletedSkuIdList() []int64 {
	return r._itemDeletedSkuIdList
}
