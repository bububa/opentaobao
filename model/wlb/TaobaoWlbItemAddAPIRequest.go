package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemAddAPIRequest 添加单个物流宝商品 API请求
// taobao.wlb.item.add
//
// 添加物流宝商品，支持物流宝子商品和属性添加
type TaobaoWlbItemAddAPIRequest struct {
	model.Params
	// 商品名称
	_name string
	// 商品标题
	_title string
	// 商品编码
	_itemCode string
	// 商品备注
	_remark string
	// NORMAL--普通商品<br/>COMBINE--组合商品<br/>DISTRIBUTION--分销
	_type string
	// 是否sku
	_isSku string
	// 属性名列表,目前支持的属性：<br/>毛重:GWeight	<br/>净重:Nweight<br/>皮重: Tweight<br/>自定义属性：<br/>paramkey1<br/>paramkey2<br/>paramkey3<br/>paramkey4
	_proNameList string
	// 属性值列表：<br/>10,8
	_proValueList string
	// 是否易碎品
	_isFriable bool
	// 是否危险品
	_isDangerous bool
	// 商品颜色
	_color string
	// 商品重量，单位G
	_weight int64
	// 商品长度，单位毫米
	_length int64
	// 商品宽度，单位毫米
	_width int64
	// 商品高度，单位毫米
	_height int64
	// 商品体积，单位立方厘米
	_volume int64
	// 货类
	_goodsCat string
	// 计价货类
	_pricingCat string
	// 商品包装材料类型
	_packageMaterial string
	// 商品价格，单位：分
	_price int64
	// 是否支持批次
	_supportBatch bool
}

// NewTaobaoWlbItemAddRequest 初始化TaobaoWlbItemAddAPIRequest对象
func NewTaobaoWlbItemAddRequest() *TaobaoWlbItemAddAPIRequest {
	return &TaobaoWlbItemAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemAddAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Name Setter
// 商品名称
func (r *TaobaoWlbItemAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoWlbItemAddAPIRequest) GetName() string {
	return r._name
}

// Set is Title Setter
// 商品标题
func (r *TaobaoWlbItemAddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// Get Title Getter
func (r TaobaoWlbItemAddAPIRequest) GetTitle() string {
	return r._title
}

// Set is ItemCode Setter
// 商品编码
func (r *TaobaoWlbItemAddAPIRequest) SetItemCode(_itemCode string) error {
	r._itemCode = _itemCode
	r.Set("item_code", _itemCode)
	return nil
}

// Get ItemCode Getter
func (r TaobaoWlbItemAddAPIRequest) GetItemCode() string {
	return r._itemCode
}

// Set is Remark Setter
// 商品备注
func (r *TaobaoWlbItemAddAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// Get Remark Getter
func (r TaobaoWlbItemAddAPIRequest) GetRemark() string {
	return r._remark
}

// Set is Type Setter
// NORMAL--普通商品<br/>COMBINE--组合商品<br/>DISTRIBUTION--分销
func (r *TaobaoWlbItemAddAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TaobaoWlbItemAddAPIRequest) GetType() string {
	return r._type
}

// Set is IsSku Setter
// 是否sku
func (r *TaobaoWlbItemAddAPIRequest) SetIsSku(_isSku string) error {
	r._isSku = _isSku
	r.Set("is_sku", _isSku)
	return nil
}

// Get IsSku Getter
func (r TaobaoWlbItemAddAPIRequest) GetIsSku() string {
	return r._isSku
}

// Set is ProNameList Setter
// 属性名列表,目前支持的属性：<br/>毛重:GWeight	<br/>净重:Nweight<br/>皮重: Tweight<br/>自定义属性：<br/>paramkey1<br/>paramkey2<br/>paramkey3<br/>paramkey4
func (r *TaobaoWlbItemAddAPIRequest) SetProNameList(_proNameList string) error {
	r._proNameList = _proNameList
	r.Set("pro_name_list", _proNameList)
	return nil
}

// Get ProNameList Getter
func (r TaobaoWlbItemAddAPIRequest) GetProNameList() string {
	return r._proNameList
}

// Set is ProValueList Setter
// 属性值列表：<br/>10,8
func (r *TaobaoWlbItemAddAPIRequest) SetProValueList(_proValueList string) error {
	r._proValueList = _proValueList
	r.Set("pro_value_list", _proValueList)
	return nil
}

// Get ProValueList Getter
func (r TaobaoWlbItemAddAPIRequest) GetProValueList() string {
	return r._proValueList
}

// Set is IsFriable Setter
// 是否易碎品
func (r *TaobaoWlbItemAddAPIRequest) SetIsFriable(_isFriable bool) error {
	r._isFriable = _isFriable
	r.Set("is_friable", _isFriable)
	return nil
}

// Get IsFriable Getter
func (r TaobaoWlbItemAddAPIRequest) GetIsFriable() bool {
	return r._isFriable
}

// Set is IsDangerous Setter
// 是否危险品
func (r *TaobaoWlbItemAddAPIRequest) SetIsDangerous(_isDangerous bool) error {
	r._isDangerous = _isDangerous
	r.Set("is_dangerous", _isDangerous)
	return nil
}

// Get IsDangerous Getter
func (r TaobaoWlbItemAddAPIRequest) GetIsDangerous() bool {
	return r._isDangerous
}

// Set is Color Setter
// 商品颜色
func (r *TaobaoWlbItemAddAPIRequest) SetColor(_color string) error {
	r._color = _color
	r.Set("color", _color)
	return nil
}

// Get Color Getter
func (r TaobaoWlbItemAddAPIRequest) GetColor() string {
	return r._color
}

// Set is Weight Setter
// 商品重量，单位G
func (r *TaobaoWlbItemAddAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// Get Weight Getter
func (r TaobaoWlbItemAddAPIRequest) GetWeight() int64 {
	return r._weight
}

// Set is Length Setter
// 商品长度，单位毫米
func (r *TaobaoWlbItemAddAPIRequest) SetLength(_length int64) error {
	r._length = _length
	r.Set("length", _length)
	return nil
}

// Get Length Getter
func (r TaobaoWlbItemAddAPIRequest) GetLength() int64 {
	return r._length
}

// Set is Width Setter
// 商品宽度，单位毫米
func (r *TaobaoWlbItemAddAPIRequest) SetWidth(_width int64) error {
	r._width = _width
	r.Set("width", _width)
	return nil
}

// Get Width Getter
func (r TaobaoWlbItemAddAPIRequest) GetWidth() int64 {
	return r._width
}

// Set is Height Setter
// 商品高度，单位毫米
func (r *TaobaoWlbItemAddAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// Get Height Getter
func (r TaobaoWlbItemAddAPIRequest) GetHeight() int64 {
	return r._height
}

// Set is Volume Setter
// 商品体积，单位立方厘米
func (r *TaobaoWlbItemAddAPIRequest) SetVolume(_volume int64) error {
	r._volume = _volume
	r.Set("volume", _volume)
	return nil
}

// Get Volume Getter
func (r TaobaoWlbItemAddAPIRequest) GetVolume() int64 {
	return r._volume
}

// Set is GoodsCat Setter
// 货类
func (r *TaobaoWlbItemAddAPIRequest) SetGoodsCat(_goodsCat string) error {
	r._goodsCat = _goodsCat
	r.Set("goods_cat", _goodsCat)
	return nil
}

// Get GoodsCat Getter
func (r TaobaoWlbItemAddAPIRequest) GetGoodsCat() string {
	return r._goodsCat
}

// Set is PricingCat Setter
// 计价货类
func (r *TaobaoWlbItemAddAPIRequest) SetPricingCat(_pricingCat string) error {
	r._pricingCat = _pricingCat
	r.Set("pricing_cat", _pricingCat)
	return nil
}

// Get PricingCat Getter
func (r TaobaoWlbItemAddAPIRequest) GetPricingCat() string {
	return r._pricingCat
}

// Set is PackageMaterial Setter
// 商品包装材料类型
func (r *TaobaoWlbItemAddAPIRequest) SetPackageMaterial(_packageMaterial string) error {
	r._packageMaterial = _packageMaterial
	r.Set("package_material", _packageMaterial)
	return nil
}

// Get PackageMaterial Getter
func (r TaobaoWlbItemAddAPIRequest) GetPackageMaterial() string {
	return r._packageMaterial
}

// Set is Price Setter
// 商品价格，单位：分
func (r *TaobaoWlbItemAddAPIRequest) SetPrice(_price int64) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// Get Price Getter
func (r TaobaoWlbItemAddAPIRequest) GetPrice() int64 {
	return r._price
}

// Set is SupportBatch Setter
// 是否支持批次
func (r *TaobaoWlbItemAddAPIRequest) SetSupportBatch(_supportBatch bool) error {
	r._supportBatch = _supportBatch
	r.Set("support_batch", _supportBatch)
	return nil
}

// Get SupportBatch Getter
func (r TaobaoWlbItemAddAPIRequest) GetSupportBatch() bool {
	return r._supportBatch
}
