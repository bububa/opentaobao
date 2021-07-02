package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemSkuDeleteAPIRequest 删除SKU API请求
// taobao.item.sku.delete
//
// 删除一个sku的数据<br/>需要删除的sku通过属性properties进行匹配查找
type TaobaoItemSkuDeleteAPIRequest struct {
	model.Params
	// Sku所属商品数字id，可通过 taobao.item.get 获取。必选
	_numIid int64
	// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
	_properties string
	// sku所属商品的价格。当用户删除sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够删除成功
	_itemPrice float64
	// sku所属商品的数量,大于0的整数。当用户删除sku，使商品数量不等于sku数量之和时候，用于修改商品的数量，使sku能够删除成功。特别是删除最后一个sku的时候，一定要设置商品数量到正常的值，否则删除失败
	_itemNum int64
	// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
	_lang string
	// 忽略警告提示.
	_ignorewarning string
}

// NewTaobaoItemSkuDeleteRequest 初始化TaobaoItemSkuDeleteAPIRequest对象
func NewTaobaoItemSkuDeleteRequest() *TaobaoItemSkuDeleteAPIRequest {
	return &TaobaoItemSkuDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemSkuDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.item.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemSkuDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNumIid is NumIid Setter
// Sku所属商品数字id，可通过 taobao.item.get 获取。必选
func (r *TaobaoItemSkuDeleteAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoItemSkuDeleteAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetProperties is Properties Setter
// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
func (r *TaobaoItemSkuDeleteAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoItemSkuDeleteAPIRequest) GetProperties() string {
	return r._properties
}

// SetItemPrice is ItemPrice Setter
// sku所属商品的价格。当用户删除sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够删除成功
func (r *TaobaoItemSkuDeleteAPIRequest) SetItemPrice(_itemPrice float64) error {
	r._itemPrice = _itemPrice
	r.Set("item_price", _itemPrice)
	return nil
}

// GetItemPrice ItemPrice Getter
func (r TaobaoItemSkuDeleteAPIRequest) GetItemPrice() float64 {
	return r._itemPrice
}

// SetItemNum is ItemNum Setter
// sku所属商品的数量,大于0的整数。当用户删除sku，使商品数量不等于sku数量之和时候，用于修改商品的数量，使sku能够删除成功。特别是删除最后一个sku的时候，一定要设置商品数量到正常的值，否则删除失败
func (r *TaobaoItemSkuDeleteAPIRequest) SetItemNum(_itemNum int64) error {
	r._itemNum = _itemNum
	r.Set("item_num", _itemNum)
	return nil
}

// GetItemNum ItemNum Getter
func (r TaobaoItemSkuDeleteAPIRequest) GetItemNum() int64 {
	return r._itemNum
}

// SetLang is Lang Setter
// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
func (r *TaobaoItemSkuDeleteAPIRequest) SetLang(_lang string) error {
	r._lang = _lang
	r.Set("lang", _lang)
	return nil
}

// GetLang Lang Getter
func (r TaobaoItemSkuDeleteAPIRequest) GetLang() string {
	return r._lang
}

// SetIgnorewarning is Ignorewarning Setter
// 忽略警告提示.
func (r *TaobaoItemSkuDeleteAPIRequest) SetIgnorewarning(_ignorewarning string) error {
	r._ignorewarning = _ignorewarning
	r.Set("ignorewarning", _ignorewarning)
	return nil
}

// GetIgnorewarning Ignorewarning Getter
func (r TaobaoItemSkuDeleteAPIRequest) GetIgnorewarning() string {
	return r._ignorewarning
}
