package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemSkuPriceUpdateAPIRequest 更新商品SKU的价格 API请求
// taobao.item.sku.price.update
//
// 更新商品SKU的价格
type TaobaoItemSkuPriceUpdateAPIRequest struct {
	model.Params
	// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
	_properties string
	// Sku的商家外部id
	_outerId string
	// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
	_lang string
	// 忽略警告提示.
	_ignorewarning string
	// Sku所属商品数字id，可通过 taobao.item.get 获取
	_numIid int64
	// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
	_quantity int64
	// Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
	_price float64
	// sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功
	_itemPrice float64
}

// NewTaobaoItemSkuPriceUpdateRequest 初始化TaobaoItemSkuPriceUpdateAPIRequest对象
func NewTaobaoItemSkuPriceUpdateRequest() *TaobaoItemSkuPriceUpdateAPIRequest {
	return &TaobaoItemSkuPriceUpdateAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemSkuPriceUpdateAPIRequest) Reset() {
	r._properties = ""
	r._outerId = ""
	r._lang = ""
	r._ignorewarning = ""
	r._numIid = 0
	r._quantity = 0
	r._price = 0
	r._itemPrice = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemSkuPriceUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.item.sku.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemSkuPriceUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemSkuPriceUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProperties is Properties Setter
// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
func (r *TaobaoItemSkuPriceUpdateAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoItemSkuPriceUpdateAPIRequest) GetProperties() string {
	return r._properties
}

// SetOuterId is OuterId Setter
// Sku的商家外部id
func (r *TaobaoItemSkuPriceUpdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoItemSkuPriceUpdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetLang is Lang Setter
// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
func (r *TaobaoItemSkuPriceUpdateAPIRequest) SetLang(_lang string) error {
	r._lang = _lang
	r.Set("lang", _lang)
	return nil
}

// GetLang Lang Getter
func (r TaobaoItemSkuPriceUpdateAPIRequest) GetLang() string {
	return r._lang
}

// SetIgnorewarning is Ignorewarning Setter
// 忽略警告提示.
func (r *TaobaoItemSkuPriceUpdateAPIRequest) SetIgnorewarning(_ignorewarning string) error {
	r._ignorewarning = _ignorewarning
	r.Set("ignorewarning", _ignorewarning)
	return nil
}

// GetIgnorewarning Ignorewarning Getter
func (r TaobaoItemSkuPriceUpdateAPIRequest) GetIgnorewarning() string {
	return r._ignorewarning
}

// SetNumIid is NumIid Setter
// Sku所属商品数字id，可通过 taobao.item.get 获取
func (r *TaobaoItemSkuPriceUpdateAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoItemSkuPriceUpdateAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetQuantity is Quantity Setter
// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
func (r *TaobaoItemSkuPriceUpdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaoItemSkuPriceUpdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}

// SetPrice is Price Setter
// Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
func (r *TaobaoItemSkuPriceUpdateAPIRequest) SetPrice(_price float64) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TaobaoItemSkuPriceUpdateAPIRequest) GetPrice() float64 {
	return r._price
}

// SetItemPrice is ItemPrice Setter
// sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功
func (r *TaobaoItemSkuPriceUpdateAPIRequest) SetItemPrice(_itemPrice float64) error {
	r._itemPrice = _itemPrice
	r.Set("item_price", _itemPrice)
	return nil
}

// GetItemPrice ItemPrice Getter
func (r TaobaoItemSkuPriceUpdateAPIRequest) GetItemPrice() float64 {
	return r._itemPrice
}

var poolTaobaoItemSkuPriceUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemSkuPriceUpdateRequest()
	},
}

// GetTaobaoItemSkuPriceUpdateRequest 从 sync.Pool 获取 TaobaoItemSkuPriceUpdateAPIRequest
func GetTaobaoItemSkuPriceUpdateAPIRequest() *TaobaoItemSkuPriceUpdateAPIRequest {
	return poolTaobaoItemSkuPriceUpdateAPIRequest.Get().(*TaobaoItemSkuPriceUpdateAPIRequest)
}

// ReleaseTaobaoItemSkuPriceUpdateAPIRequest 将 TaobaoItemSkuPriceUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemSkuPriceUpdateAPIRequest(v *TaobaoItemSkuPriceUpdateAPIRequest) {
	v.Reset()
	poolTaobaoItemSkuPriceUpdateAPIRequest.Put(v)
}
