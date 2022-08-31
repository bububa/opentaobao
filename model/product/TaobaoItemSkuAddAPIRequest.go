package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemSkuAddAPIRequest 添加SKU API请求
// taobao.item.sku.add
//
// 新增一个sku到num_iid指定的商品中 &lt;br/&gt;传入的iid所对应的商品必须属于当前会话的用户
type TaobaoItemSkuAddAPIRequest struct {
	model.Params
	// Sku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。
	_properties string
	// Sku的商家外部id
	_outerId string
	// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
	_lang string
	// 忽略警告提示.
	_ignorewarning string
	// Sku所属商品数字id。必选
	_numIid int64
	// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数
	_quantity int64
	// Sku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分
	_price float64
	// sku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功
	_itemPrice float64
}

// NewTaobaoItemSkuAddRequest 初始化TaobaoItemSkuAddAPIRequest对象
func NewTaobaoItemSkuAddRequest() *TaobaoItemSkuAddAPIRequest {
	return &TaobaoItemSkuAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemSkuAddAPIRequest) GetApiMethodName() string {
	return "taobao.item.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemSkuAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProperties is Properties Setter
// Sku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。
func (r *TaobaoItemSkuAddAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoItemSkuAddAPIRequest) GetProperties() string {
	return r._properties
}

// SetOuterId is OuterId Setter
// Sku的商家外部id
func (r *TaobaoItemSkuAddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoItemSkuAddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetLang is Lang Setter
// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
func (r *TaobaoItemSkuAddAPIRequest) SetLang(_lang string) error {
	r._lang = _lang
	r.Set("lang", _lang)
	return nil
}

// GetLang Lang Getter
func (r TaobaoItemSkuAddAPIRequest) GetLang() string {
	return r._lang
}

// SetIgnorewarning is Ignorewarning Setter
// 忽略警告提示.
func (r *TaobaoItemSkuAddAPIRequest) SetIgnorewarning(_ignorewarning string) error {
	r._ignorewarning = _ignorewarning
	r.Set("ignorewarning", _ignorewarning)
	return nil
}

// GetIgnorewarning Ignorewarning Getter
func (r TaobaoItemSkuAddAPIRequest) GetIgnorewarning() string {
	return r._ignorewarning
}

// SetNumIid is NumIid Setter
// Sku所属商品数字id。必选
func (r *TaobaoItemSkuAddAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoItemSkuAddAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetQuantity is Quantity Setter
// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数
func (r *TaobaoItemSkuAddAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaoItemSkuAddAPIRequest) GetQuantity() int64 {
	return r._quantity
}

// SetPrice is Price Setter
// Sku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分
func (r *TaobaoItemSkuAddAPIRequest) SetPrice(_price float64) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TaobaoItemSkuAddAPIRequest) GetPrice() float64 {
	return r._price
}

// SetItemPrice is ItemPrice Setter
// sku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功
func (r *TaobaoItemSkuAddAPIRequest) SetItemPrice(_itemPrice float64) error {
	r._itemPrice = _itemPrice
	r.Set("item_price", _itemPrice)
	return nil
}

// GetItemPrice ItemPrice Getter
func (r TaobaoItemSkuAddAPIRequest) GetItemPrice() float64 {
	return r._itemPrice
}
