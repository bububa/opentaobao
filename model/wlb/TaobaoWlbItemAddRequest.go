package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加单个物流宝商品 API请求
taobao.wlb.item.add

添加物流宝商品，支持物流宝子商品和属性添加
*/
type TaobaoWlbItemAddRequest struct {
    model.Params
    // 商品名称
    _name   string
    // 商品标题
    _title   string
    // 商品编码
    _itemCode   string
    // 商品备注
    _remark   string
    // NORMAL--普通商品<br/>COMBINE--组合商品<br/>DISTRIBUTION--分销
    _type   string
    // 是否sku
    _isSku   string
    // 属性名列表,目前支持的属性：<br/>毛重:GWeight	<br/>净重:Nweight<br/>皮重: Tweight<br/>自定义属性：<br/>paramkey1<br/>paramkey2<br/>paramkey3<br/>paramkey4
    _proNameList   string
    // 属性值列表：<br/>10,8
    _proValueList   string
    // 是否易碎品
    _isFriable   bool
    // 是否危险品
    _isDangerous   bool
    // 商品颜色
    _color   string
    // 商品重量，单位G
    _weight   int64
    // 商品长度，单位毫米
    _length   int64
    // 商品宽度，单位毫米
    _width   int64
    // 商品高度，单位毫米
    _height   int64
    // 商品体积，单位立方厘米
    _volume   int64
    // 货类
    _goodsCat   string
    // 计价货类
    _pricingCat   string
    // 商品包装材料类型
    _packageMaterial   string
    // 商品价格，单位：分
    _price   int64
    // 是否支持批次
    _supportBatch   bool
}

// 初始化TaobaoWlbItemAddRequest对象
func NewTaobaoWlbItemAddRequest() *TaobaoWlbItemAddRequest{
    return &TaobaoWlbItemAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemAddRequest) GetApiMethodName() string {
    return "taobao.wlb.item.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 商品名称
func (r *TaobaoWlbItemAddRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoWlbItemAddRequest) GetName() string {
    return r._name
}
// Title Setter
// 商品标题
func (r *TaobaoWlbItemAddRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoWlbItemAddRequest) GetTitle() string {
    return r._title
}
// ItemCode Setter
// 商品编码
func (r *TaobaoWlbItemAddRequest) SetItemCode(_itemCode string) error {
    r._itemCode = _itemCode
    r.Set("item_code", _itemCode)
    return nil
}

// ItemCode Getter
func (r TaobaoWlbItemAddRequest) GetItemCode() string {
    return r._itemCode
}
// Remark Setter
// 商品备注
func (r *TaobaoWlbItemAddRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoWlbItemAddRequest) GetRemark() string {
    return r._remark
}
// Type Setter
// NORMAL--普通商品<br/>COMBINE--组合商品<br/>DISTRIBUTION--分销
func (r *TaobaoWlbItemAddRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoWlbItemAddRequest) GetType() string {
    return r._type
}
// IsSku Setter
// 是否sku
func (r *TaobaoWlbItemAddRequest) SetIsSku(_isSku string) error {
    r._isSku = _isSku
    r.Set("is_sku", _isSku)
    return nil
}

// IsSku Getter
func (r TaobaoWlbItemAddRequest) GetIsSku() string {
    return r._isSku
}
// ProNameList Setter
// 属性名列表,目前支持的属性：<br/>毛重:GWeight	<br/>净重:Nweight<br/>皮重: Tweight<br/>自定义属性：<br/>paramkey1<br/>paramkey2<br/>paramkey3<br/>paramkey4
func (r *TaobaoWlbItemAddRequest) SetProNameList(_proNameList string) error {
    r._proNameList = _proNameList
    r.Set("pro_name_list", _proNameList)
    return nil
}

// ProNameList Getter
func (r TaobaoWlbItemAddRequest) GetProNameList() string {
    return r._proNameList
}
// ProValueList Setter
// 属性值列表：<br/>10,8
func (r *TaobaoWlbItemAddRequest) SetProValueList(_proValueList string) error {
    r._proValueList = _proValueList
    r.Set("pro_value_list", _proValueList)
    return nil
}

// ProValueList Getter
func (r TaobaoWlbItemAddRequest) GetProValueList() string {
    return r._proValueList
}
// IsFriable Setter
// 是否易碎品
func (r *TaobaoWlbItemAddRequest) SetIsFriable(_isFriable bool) error {
    r._isFriable = _isFriable
    r.Set("is_friable", _isFriable)
    return nil
}

// IsFriable Getter
func (r TaobaoWlbItemAddRequest) GetIsFriable() bool {
    return r._isFriable
}
// IsDangerous Setter
// 是否危险品
func (r *TaobaoWlbItemAddRequest) SetIsDangerous(_isDangerous bool) error {
    r._isDangerous = _isDangerous
    r.Set("is_dangerous", _isDangerous)
    return nil
}

// IsDangerous Getter
func (r TaobaoWlbItemAddRequest) GetIsDangerous() bool {
    return r._isDangerous
}
// Color Setter
// 商品颜色
func (r *TaobaoWlbItemAddRequest) SetColor(_color string) error {
    r._color = _color
    r.Set("color", _color)
    return nil
}

// Color Getter
func (r TaobaoWlbItemAddRequest) GetColor() string {
    return r._color
}
// Weight Setter
// 商品重量，单位G
func (r *TaobaoWlbItemAddRequest) SetWeight(_weight int64) error {
    r._weight = _weight
    r.Set("weight", _weight)
    return nil
}

// Weight Getter
func (r TaobaoWlbItemAddRequest) GetWeight() int64 {
    return r._weight
}
// Length Setter
// 商品长度，单位毫米
func (r *TaobaoWlbItemAddRequest) SetLength(_length int64) error {
    r._length = _length
    r.Set("length", _length)
    return nil
}

// Length Getter
func (r TaobaoWlbItemAddRequest) GetLength() int64 {
    return r._length
}
// Width Setter
// 商品宽度，单位毫米
func (r *TaobaoWlbItemAddRequest) SetWidth(_width int64) error {
    r._width = _width
    r.Set("width", _width)
    return nil
}

// Width Getter
func (r TaobaoWlbItemAddRequest) GetWidth() int64 {
    return r._width
}
// Height Setter
// 商品高度，单位毫米
func (r *TaobaoWlbItemAddRequest) SetHeight(_height int64) error {
    r._height = _height
    r.Set("height", _height)
    return nil
}

// Height Getter
func (r TaobaoWlbItemAddRequest) GetHeight() int64 {
    return r._height
}
// Volume Setter
// 商品体积，单位立方厘米
func (r *TaobaoWlbItemAddRequest) SetVolume(_volume int64) error {
    r._volume = _volume
    r.Set("volume", _volume)
    return nil
}

// Volume Getter
func (r TaobaoWlbItemAddRequest) GetVolume() int64 {
    return r._volume
}
// GoodsCat Setter
// 货类
func (r *TaobaoWlbItemAddRequest) SetGoodsCat(_goodsCat string) error {
    r._goodsCat = _goodsCat
    r.Set("goods_cat", _goodsCat)
    return nil
}

// GoodsCat Getter
func (r TaobaoWlbItemAddRequest) GetGoodsCat() string {
    return r._goodsCat
}
// PricingCat Setter
// 计价货类
func (r *TaobaoWlbItemAddRequest) SetPricingCat(_pricingCat string) error {
    r._pricingCat = _pricingCat
    r.Set("pricing_cat", _pricingCat)
    return nil
}

// PricingCat Getter
func (r TaobaoWlbItemAddRequest) GetPricingCat() string {
    return r._pricingCat
}
// PackageMaterial Setter
// 商品包装材料类型
func (r *TaobaoWlbItemAddRequest) SetPackageMaterial(_packageMaterial string) error {
    r._packageMaterial = _packageMaterial
    r.Set("package_material", _packageMaterial)
    return nil
}

// PackageMaterial Getter
func (r TaobaoWlbItemAddRequest) GetPackageMaterial() string {
    return r._packageMaterial
}
// Price Setter
// 商品价格，单位：分
func (r *TaobaoWlbItemAddRequest) SetPrice(_price int64) error {
    r._price = _price
    r.Set("price", _price)
    return nil
}

// Price Getter
func (r TaobaoWlbItemAddRequest) GetPrice() int64 {
    return r._price
}
// SupportBatch Setter
// 是否支持批次
func (r *TaobaoWlbItemAddRequest) SetSupportBatch(_supportBatch bool) error {
    r._supportBatch = _supportBatch
    r.Set("support_batch", _supportBatch)
    return nil
}

// SupportBatch Getter
func (r TaobaoWlbItemAddRequest) GetSupportBatch() bool {
    return r._supportBatch
}
