package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemskudeleteAPIRequest 删除SKU API请求
// taobao.item.sku.delete
//
// 删除一个sku的数据&lt;br/&gt;需要删除的sku通过属性properties进行匹配查找
type TaobaoitemskudeleteAPIRequest struct {
	model.Params
	// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
	_properties string
	// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
	_lang string
	// 忽略警告提示.
	_ignorewarning string
	// Sku所属商品数字id，可通过 taobao.item.get 获取。必选
	_numIid int64
	// sku所属商品的价格。当用户删除sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够删除成功
	_itemPrice float64
	// sku所属商品的数量,大于0的整数。当用户删除sku，使商品数量不等于sku数量之和时候，用于修改商品的数量，使sku能够删除成功。特别是删除最后一个sku的时候，一定要设置商品数量到正常的值，否则删除失败
	_itemNum int64
}

// NewTaobaoitemskudeleteRequest 初始化TaobaoitemskudeleteAPIRequest对象
func NewTaobaoitemskudeleteRequest() *TaobaoitemskudeleteAPIRequest {
	return &TaobaoitemskudeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemskudeleteAPIRequest) GetApiMethodName() string {
	return "taobao.item.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemskudeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemskudeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProperties is Properties Setter
// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
func (r *TaobaoitemskudeleteAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoitemskudeleteAPIRequest) GetProperties() string {
	return r._properties
}

// SetLang is Lang Setter
// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
func (r *TaobaoitemskudeleteAPIRequest) SetLang(_lang string) error {
	r._lang = _lang
	r.Set("lang", _lang)
	return nil
}

// GetLang Lang Getter
func (r TaobaoitemskudeleteAPIRequest) GetLang() string {
	return r._lang
}

// SetIgnorewarning is Ignorewarning Setter
// 忽略警告提示.
func (r *TaobaoitemskudeleteAPIRequest) SetIgnorewarning(_ignorewarning string) error {
	r._ignorewarning = _ignorewarning
	r.Set("ignorewarning", _ignorewarning)
	return nil
}

// GetIgnorewarning Ignorewarning Getter
func (r TaobaoitemskudeleteAPIRequest) GetIgnorewarning() string {
	return r._ignorewarning
}

// SetNumIid is NumIid Setter
// Sku所属商品数字id，可通过 taobao.item.get 获取。必选
func (r *TaobaoitemskudeleteAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoitemskudeleteAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetItemPrice is ItemPrice Setter
// sku所属商品的价格。当用户删除sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够删除成功
func (r *TaobaoitemskudeleteAPIRequest) SetItemPrice(_itemPrice float64) error {
	r._itemPrice = _itemPrice
	r.Set("item_price", _itemPrice)
	return nil
}

// GetItemPrice ItemPrice Getter
func (r TaobaoitemskudeleteAPIRequest) GetItemPrice() float64 {
	return r._itemPrice
}

// SetItemNum is ItemNum Setter
// sku所属商品的数量,大于0的整数。当用户删除sku，使商品数量不等于sku数量之和时候，用于修改商品的数量，使sku能够删除成功。特别是删除最后一个sku的时候，一定要设置商品数量到正常的值，否则删除失败
func (r *TaobaoitemskudeleteAPIRequest) SetItemNum(_itemNum int64) error {
	r._itemNum = _itemNum
	r.Set("item_num", _itemNum)
	return nil
}

// GetItemNum ItemNum Getter
func (r TaobaoitemskudeleteAPIRequest) GetItemNum() int64 {
	return r._itemNum
}
