package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelGereralskuUpdateAPIRequest 发布SKU信息（如果properties重复 则更新） API请求
// alitrip.travel.gereralsku.update
//
// 发布SKU信息（如果properties重复 则更新）
type AlitripTravelGereralskuUpdateAPIRequest struct {
	model.Params
	// 淘宝商品ID
	_itemId int64
	// Sku的销售价格，普通商品使用。精确到2位小数;单位:分。如:20007，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
	_price int64
	// Sku的库存数量，普通商品使用。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
	_quantity int64
	// 商家编码
	_outerId string
	// sku销售属性别名；如套餐1 需要调整成其他 需要在这里修改
	_alias []PropertyAliasInfo
	// 商品属性列表；由类目的属性PID和VID组成，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对
	_properties []CatPropInfo
	// SKU的销售价格库存，日历商品使用
	_dateList []DateInventoryAndPrice
}

// NewAlitripTravelGereralskuUpdateRequest 初始化AlitripTravelGereralskuUpdateAPIRequest对象
func NewAlitripTravelGereralskuUpdateRequest() *AlitripTravelGereralskuUpdateAPIRequest {
	return &AlitripTravelGereralskuUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelGereralskuUpdateAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.gereralsku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelGereralskuUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// 淘宝商品ID
func (r *AlitripTravelGereralskuUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripTravelGereralskuUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetPrice is Price Setter
// Sku的销售价格，普通商品使用。精确到2位小数;单位:分。如:20007，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
func (r *AlitripTravelGereralskuUpdateAPIRequest) SetPrice(_price int64) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r AlitripTravelGereralskuUpdateAPIRequest) GetPrice() int64 {
	return r._price
}

// SetQuantity is Quantity Setter
// Sku的库存数量，普通商品使用。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
func (r *AlitripTravelGereralskuUpdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r AlitripTravelGereralskuUpdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}

// SetOuterId is OuterId Setter
// 商家编码
func (r *AlitripTravelGereralskuUpdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlitripTravelGereralskuUpdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetAlias is Alias Setter
// sku销售属性别名；如套餐1 需要调整成其他 需要在这里修改
func (r *AlitripTravelGereralskuUpdateAPIRequest) SetAlias(_alias []PropertyAliasInfo) error {
	r._alias = _alias
	r.Set("alias", _alias)
	return nil
}

// GetAlias Alias Getter
func (r AlitripTravelGereralskuUpdateAPIRequest) GetAlias() []PropertyAliasInfo {
	return r._alias
}

// SetProperties is Properties Setter
// 商品属性列表；由类目的属性PID和VID组成，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对
func (r *AlitripTravelGereralskuUpdateAPIRequest) SetProperties(_properties []CatPropInfo) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r AlitripTravelGereralskuUpdateAPIRequest) GetProperties() []CatPropInfo {
	return r._properties
}

// SetDateList is DateList Setter
// SKU的销售价格库存，日历商品使用
func (r *AlitripTravelGereralskuUpdateAPIRequest) SetDateList(_dateList []DateInventoryAndPrice) error {
	r._dateList = _dateList
	r.Set("date_list", _dateList)
	return nil
}

// GetDateList DateList Getter
func (r AlitripTravelGereralskuUpdateAPIRequest) GetDateList() []DateInventoryAndPrice {
	return r._dateList
}
