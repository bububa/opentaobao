package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelItemSkuCalendarUpdateAPIRequest 酒店非标套餐商品日历库存宝贝SKU更新接口 API请求
// alitrip.tuan.hotel.item.sku.calendar.update
//
// 商户对发布的日历库存类型的宝贝套餐价格库存信息进行更新，仅提供日历库存的宝贝SKU的更新功能，skuId须传递商品已存在的skuId，若想进行SKU新增操作，请选择使用alitrip.tuan.hotel.item.sku.update接口。提供增量更新SKU功能，对于日历库存若传递日期信息，参数中若包含某一日期的价格和库存，则对此日期的数据进行覆盖更新，若不传递则保留此日期的价格库存信息。
type AlitripTuanHotelItemSkuCalendarUpdateAPIRequest struct {
	model.Params
	// 暂不支持此接口对SKU的部分属性进行更新，包括以下属性： 套餐名称、价格、原价、库存、间夜、商家编码、人数、使用次数等
	_itemSkuList []TopTuanItemSkuVo
	// 宝贝ID
	_itemId int64
	// 宝贝所属类目
	_catId int64
}

// NewAlitripTuanHotelItemSkuCalendarUpdateRequest 初始化AlitripTuanHotelItemSkuCalendarUpdateAPIRequest对象
func NewAlitripTuanHotelItemSkuCalendarUpdateRequest() *AlitripTuanHotelItemSkuCalendarUpdateAPIRequest {
	return &AlitripTuanHotelItemSkuCalendarUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelItemSkuCalendarUpdateAPIRequest) GetApiMethodName() string {
	return "alitrip.tuan.hotel.item.sku.calendar.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelItemSkuCalendarUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTuanHotelItemSkuCalendarUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemSkuList is ItemSkuList Setter
// 暂不支持此接口对SKU的部分属性进行更新，包括以下属性： 套餐名称、价格、原价、库存、间夜、商家编码、人数、使用次数等
func (r *AlitripTuanHotelItemSkuCalendarUpdateAPIRequest) SetItemSkuList(_itemSkuList []TopTuanItemSkuVo) error {
	r._itemSkuList = _itemSkuList
	r.Set("item_sku_list", _itemSkuList)
	return nil
}

// GetItemSkuList ItemSkuList Getter
func (r AlitripTuanHotelItemSkuCalendarUpdateAPIRequest) GetItemSkuList() []TopTuanItemSkuVo {
	return r._itemSkuList
}

// SetItemId is ItemId Setter
// 宝贝ID
func (r *AlitripTuanHotelItemSkuCalendarUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripTuanHotelItemSkuCalendarUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCatId is CatId Setter
// 宝贝所属类目
func (r *AlitripTuanHotelItemSkuCalendarUpdateAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlitripTuanHotelItemSkuCalendarUpdateAPIRequest) GetCatId() int64 {
	return r._catId
}
