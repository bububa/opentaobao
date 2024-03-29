package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelProductGereralskuUpdateAPIRequest (供销)船票通用类目sku新增&编辑API API请求
// alitrip.travel.product.gereralsku.update
//
// 发布SKU信息（如果properties重复 则更新）
type AlitripTravelProductGereralskuUpdateAPIRequest struct {
	model.Params
	// sku销售属性别名；如套餐1 需要调整成其他 需要在这里修改
	_alias []PropertyAliasInfo
	// 商品属性列表；由类目的属性PID和VID组成，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对
	_properties []CatPropInfo
	// SKU的销售价格库存，日历商品使用
	_dateList []DateInventoryAndPrice
	// 商家编码
	_outerId string
	// 淘宝商品ID
	_itemId int64
	// Sku的销售价格。精确到2位小数;单位:分。如:20007，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
	_price int64
	// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
	_quantity int64
}

// NewAlitripTravelProductGereralskuUpdateRequest 初始化AlitripTravelProductGereralskuUpdateAPIRequest对象
func NewAlitripTravelProductGereralskuUpdateRequest() *AlitripTravelProductGereralskuUpdateAPIRequest {
	return &AlitripTravelProductGereralskuUpdateAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelProductGereralskuUpdateAPIRequest) Reset() {
	r._alias = r._alias[:0]
	r._properties = r._properties[:0]
	r._dateList = r._dateList[:0]
	r._outerId = ""
	r._itemId = 0
	r._price = 0
	r._quantity = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelProductGereralskuUpdateAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.product.gereralsku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelProductGereralskuUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelProductGereralskuUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlias is Alias Setter
// sku销售属性别名；如套餐1 需要调整成其他 需要在这里修改
func (r *AlitripTravelProductGereralskuUpdateAPIRequest) SetAlias(_alias []PropertyAliasInfo) error {
	r._alias = _alias
	r.Set("alias", _alias)
	return nil
}

// GetAlias Alias Getter
func (r AlitripTravelProductGereralskuUpdateAPIRequest) GetAlias() []PropertyAliasInfo {
	return r._alias
}

// SetProperties is Properties Setter
// 商品属性列表；由类目的属性PID和VID组成，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对
func (r *AlitripTravelProductGereralskuUpdateAPIRequest) SetProperties(_properties []CatPropInfo) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r AlitripTravelProductGereralskuUpdateAPIRequest) GetProperties() []CatPropInfo {
	return r._properties
}

// SetDateList is DateList Setter
// SKU的销售价格库存，日历商品使用
func (r *AlitripTravelProductGereralskuUpdateAPIRequest) SetDateList(_dateList []DateInventoryAndPrice) error {
	r._dateList = _dateList
	r.Set("date_list", _dateList)
	return nil
}

// GetDateList DateList Getter
func (r AlitripTravelProductGereralskuUpdateAPIRequest) GetDateList() []DateInventoryAndPrice {
	return r._dateList
}

// SetOuterId is OuterId Setter
// 商家编码
func (r *AlitripTravelProductGereralskuUpdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlitripTravelProductGereralskuUpdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetItemId is ItemId Setter
// 淘宝商品ID
func (r *AlitripTravelProductGereralskuUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripTravelProductGereralskuUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetPrice is Price Setter
// Sku的销售价格。精确到2位小数;单位:分。如:20007，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
func (r *AlitripTravelProductGereralskuUpdateAPIRequest) SetPrice(_price int64) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r AlitripTravelProductGereralskuUpdateAPIRequest) GetPrice() int64 {
	return r._price
}

// SetQuantity is Quantity Setter
// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
func (r *AlitripTravelProductGereralskuUpdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r AlitripTravelProductGereralskuUpdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}

var poolAlitripTravelProductGereralskuUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelProductGereralskuUpdateRequest()
	},
}

// GetAlitripTravelProductGereralskuUpdateRequest 从 sync.Pool 获取 AlitripTravelProductGereralskuUpdateAPIRequest
func GetAlitripTravelProductGereralskuUpdateAPIRequest() *AlitripTravelProductGereralskuUpdateAPIRequest {
	return poolAlitripTravelProductGereralskuUpdateAPIRequest.Get().(*AlitripTravelProductGereralskuUpdateAPIRequest)
}

// ReleaseAlitripTravelProductGereralskuUpdateAPIRequest 将 AlitripTravelProductGereralskuUpdateAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelProductGereralskuUpdateAPIRequest(v *AlitripTravelProductGereralskuUpdateAPIRequest) {
	v.Reset()
	poolAlitripTravelProductGereralskuUpdateAPIRequest.Put(v)
}
