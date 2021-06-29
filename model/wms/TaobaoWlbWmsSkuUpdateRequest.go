package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品信息的更新 API请求
taobao.wlb.wms.sku.update

商品信息的更新
*/
type TaobaoWlbWmsSkuUpdateRequest struct {
    model.Params
    // 外部系统ID
    _itemId   string
    // 仓库编码
    _storeCode   string
    // 商品名称
    _name   string
    // 商品标题
    _title   string
    // 商品类别编码（外部系统类别）
    _category   string
    // 商品类别名称
    _categoryName   string
    // 品牌编码
    _brand   string
    // 品牌名称
    _brandName   string
    // 规格
    _specification   string
    // 颜色
    _color   string
    // 尺码
    _size   string
    // 毛重，单位克
    _grossWeight   int64
    // 净重，单位克
    _netWeight   int64
    // 长度，单位毫米
    _length   int64
    // 宽度，单位毫米
    _width   int64
    // 高度，单位毫米
    _height   int64
    // 体积，单位立方厘米
    _volume   int64
    // 箱规
    _pcs   int64
    // 产地
    _originAddress   int64
    // 批准文号
    _approvalNumber   string
    // 是否启用保质期管理
    _isShelflife   bool
    // 商品保质期天数
    _lifecycle   int64
    // 保质期禁收天数
    _rejectLifecycle   int64
    // 保质期禁售天数
    _lockupLifecycle   int64
    // 保质期预警天数
    _adventLifecycle   int64
    // 是否启用序列号管理
    _isSnMgt   bool
    // 是否易碎品
    _isHygroscopic   bool
    // 是否危险品
    _isDanger   bool
    // 吊牌价，单位分
    _tagPrice   int64
    // 零售价，单位分
    _itemPrice   int64
    // 成本价，单位分
    _costPrice   int64
    // 是否启用批次管理
    _isBatchMgt   bool
    // 启用标识
    _useYn   bool
    // 拓展属性
    _extendFields   string
    // 条形码，多条码请用”;”分隔；条码更新只技术增量更新，已有条码无法修改，只能在原条码基础上增加新的条码。例：原商品条码为：BAR001，要增加一条新条码BAR002时，此字段内容应填写为：BAR001;BAR002
    _barCode   string
    // 商品属性编码
    _attribute   string
    // 商品类型:NORMAL-普通商品、 COMBINE-组合商品、 DISTRIBUTION-分销商品
    _type   string
    // 是否区域销售属性
    _isAreaSale   bool
}

// 初始化TaobaoWlbWmsSkuUpdateRequest对象
func NewTaobaoWlbWmsSkuUpdateRequest() *TaobaoWlbWmsSkuUpdateRequest{
    return &TaobaoWlbWmsSkuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsSkuUpdateRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.sku.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsSkuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 外部系统ID
func (r *TaobaoWlbWmsSkuUpdateRequest) SetItemId(_itemId string) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetItemId() string {
    return r._itemId
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsSkuUpdateRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetStoreCode() string {
    return r._storeCode
}
// Name Setter
// 商品名称
func (r *TaobaoWlbWmsSkuUpdateRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetName() string {
    return r._name
}
// Title Setter
// 商品标题
func (r *TaobaoWlbWmsSkuUpdateRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetTitle() string {
    return r._title
}
// Category Setter
// 商品类别编码（外部系统类别）
func (r *TaobaoWlbWmsSkuUpdateRequest) SetCategory(_category string) error {
    r._category = _category
    r.Set("category", _category)
    return nil
}

// Category Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetCategory() string {
    return r._category
}
// CategoryName Setter
// 商品类别名称
func (r *TaobaoWlbWmsSkuUpdateRequest) SetCategoryName(_categoryName string) error {
    r._categoryName = _categoryName
    r.Set("category_name", _categoryName)
    return nil
}

// CategoryName Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetCategoryName() string {
    return r._categoryName
}
// Brand Setter
// 品牌编码
func (r *TaobaoWlbWmsSkuUpdateRequest) SetBrand(_brand string) error {
    r._brand = _brand
    r.Set("brand", _brand)
    return nil
}

// Brand Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetBrand() string {
    return r._brand
}
// BrandName Setter
// 品牌名称
func (r *TaobaoWlbWmsSkuUpdateRequest) SetBrandName(_brandName string) error {
    r._brandName = _brandName
    r.Set("brand_name", _brandName)
    return nil
}

// BrandName Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetBrandName() string {
    return r._brandName
}
// Specification Setter
// 规格
func (r *TaobaoWlbWmsSkuUpdateRequest) SetSpecification(_specification string) error {
    r._specification = _specification
    r.Set("specification", _specification)
    return nil
}

// Specification Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetSpecification() string {
    return r._specification
}
// Color Setter
// 颜色
func (r *TaobaoWlbWmsSkuUpdateRequest) SetColor(_color string) error {
    r._color = _color
    r.Set("color", _color)
    return nil
}

// Color Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetColor() string {
    return r._color
}
// Size Setter
// 尺码
func (r *TaobaoWlbWmsSkuUpdateRequest) SetSize(_size string) error {
    r._size = _size
    r.Set("size", _size)
    return nil
}

// Size Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetSize() string {
    return r._size
}
// GrossWeight Setter
// 毛重，单位克
func (r *TaobaoWlbWmsSkuUpdateRequest) SetGrossWeight(_grossWeight int64) error {
    r._grossWeight = _grossWeight
    r.Set("gross_weight", _grossWeight)
    return nil
}

// GrossWeight Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetGrossWeight() int64 {
    return r._grossWeight
}
// NetWeight Setter
// 净重，单位克
func (r *TaobaoWlbWmsSkuUpdateRequest) SetNetWeight(_netWeight int64) error {
    r._netWeight = _netWeight
    r.Set("net_weight", _netWeight)
    return nil
}

// NetWeight Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetNetWeight() int64 {
    return r._netWeight
}
// Length Setter
// 长度，单位毫米
func (r *TaobaoWlbWmsSkuUpdateRequest) SetLength(_length int64) error {
    r._length = _length
    r.Set("length", _length)
    return nil
}

// Length Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetLength() int64 {
    return r._length
}
// Width Setter
// 宽度，单位毫米
func (r *TaobaoWlbWmsSkuUpdateRequest) SetWidth(_width int64) error {
    r._width = _width
    r.Set("width", _width)
    return nil
}

// Width Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetWidth() int64 {
    return r._width
}
// Height Setter
// 高度，单位毫米
func (r *TaobaoWlbWmsSkuUpdateRequest) SetHeight(_height int64) error {
    r._height = _height
    r.Set("height", _height)
    return nil
}

// Height Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetHeight() int64 {
    return r._height
}
// Volume Setter
// 体积，单位立方厘米
func (r *TaobaoWlbWmsSkuUpdateRequest) SetVolume(_volume int64) error {
    r._volume = _volume
    r.Set("volume", _volume)
    return nil
}

// Volume Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetVolume() int64 {
    return r._volume
}
// Pcs Setter
// 箱规
func (r *TaobaoWlbWmsSkuUpdateRequest) SetPcs(_pcs int64) error {
    r._pcs = _pcs
    r.Set("pcs", _pcs)
    return nil
}

// Pcs Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetPcs() int64 {
    return r._pcs
}
// OriginAddress Setter
// 产地
func (r *TaobaoWlbWmsSkuUpdateRequest) SetOriginAddress(_originAddress int64) error {
    r._originAddress = _originAddress
    r.Set("origin_address", _originAddress)
    return nil
}

// OriginAddress Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetOriginAddress() int64 {
    return r._originAddress
}
// ApprovalNumber Setter
// 批准文号
func (r *TaobaoWlbWmsSkuUpdateRequest) SetApprovalNumber(_approvalNumber string) error {
    r._approvalNumber = _approvalNumber
    r.Set("approval_number", _approvalNumber)
    return nil
}

// ApprovalNumber Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetApprovalNumber() string {
    return r._approvalNumber
}
// IsShelflife Setter
// 是否启用保质期管理
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsShelflife(_isShelflife bool) error {
    r._isShelflife = _isShelflife
    r.Set("is_shelflife", _isShelflife)
    return nil
}

// IsShelflife Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsShelflife() bool {
    return r._isShelflife
}
// Lifecycle Setter
// 商品保质期天数
func (r *TaobaoWlbWmsSkuUpdateRequest) SetLifecycle(_lifecycle int64) error {
    r._lifecycle = _lifecycle
    r.Set("lifecycle", _lifecycle)
    return nil
}

// Lifecycle Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetLifecycle() int64 {
    return r._lifecycle
}
// RejectLifecycle Setter
// 保质期禁收天数
func (r *TaobaoWlbWmsSkuUpdateRequest) SetRejectLifecycle(_rejectLifecycle int64) error {
    r._rejectLifecycle = _rejectLifecycle
    r.Set("reject_lifecycle", _rejectLifecycle)
    return nil
}

// RejectLifecycle Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetRejectLifecycle() int64 {
    return r._rejectLifecycle
}
// LockupLifecycle Setter
// 保质期禁售天数
func (r *TaobaoWlbWmsSkuUpdateRequest) SetLockupLifecycle(_lockupLifecycle int64) error {
    r._lockupLifecycle = _lockupLifecycle
    r.Set("lockup_lifecycle", _lockupLifecycle)
    return nil
}

// LockupLifecycle Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetLockupLifecycle() int64 {
    return r._lockupLifecycle
}
// AdventLifecycle Setter
// 保质期预警天数
func (r *TaobaoWlbWmsSkuUpdateRequest) SetAdventLifecycle(_adventLifecycle int64) error {
    r._adventLifecycle = _adventLifecycle
    r.Set("advent_lifecycle", _adventLifecycle)
    return nil
}

// AdventLifecycle Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetAdventLifecycle() int64 {
    return r._adventLifecycle
}
// IsSnMgt Setter
// 是否启用序列号管理
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsSnMgt(_isSnMgt bool) error {
    r._isSnMgt = _isSnMgt
    r.Set("is_sn_mgt", _isSnMgt)
    return nil
}

// IsSnMgt Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsSnMgt() bool {
    return r._isSnMgt
}
// IsHygroscopic Setter
// 是否易碎品
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsHygroscopic(_isHygroscopic bool) error {
    r._isHygroscopic = _isHygroscopic
    r.Set("is_hygroscopic", _isHygroscopic)
    return nil
}

// IsHygroscopic Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsHygroscopic() bool {
    return r._isHygroscopic
}
// IsDanger Setter
// 是否危险品
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsDanger(_isDanger bool) error {
    r._isDanger = _isDanger
    r.Set("is_danger", _isDanger)
    return nil
}

// IsDanger Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsDanger() bool {
    return r._isDanger
}
// TagPrice Setter
// 吊牌价，单位分
func (r *TaobaoWlbWmsSkuUpdateRequest) SetTagPrice(_tagPrice int64) error {
    r._tagPrice = _tagPrice
    r.Set("tag_price", _tagPrice)
    return nil
}

// TagPrice Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetTagPrice() int64 {
    return r._tagPrice
}
// ItemPrice Setter
// 零售价，单位分
func (r *TaobaoWlbWmsSkuUpdateRequest) SetItemPrice(_itemPrice int64) error {
    r._itemPrice = _itemPrice
    r.Set("item_price", _itemPrice)
    return nil
}

// ItemPrice Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetItemPrice() int64 {
    return r._itemPrice
}
// CostPrice Setter
// 成本价，单位分
func (r *TaobaoWlbWmsSkuUpdateRequest) SetCostPrice(_costPrice int64) error {
    r._costPrice = _costPrice
    r.Set("cost_price", _costPrice)
    return nil
}

// CostPrice Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetCostPrice() int64 {
    return r._costPrice
}
// IsBatchMgt Setter
// 是否启用批次管理
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsBatchMgt(_isBatchMgt bool) error {
    r._isBatchMgt = _isBatchMgt
    r.Set("is_batch_mgt", _isBatchMgt)
    return nil
}

// IsBatchMgt Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsBatchMgt() bool {
    return r._isBatchMgt
}
// UseYn Setter
// 启用标识
func (r *TaobaoWlbWmsSkuUpdateRequest) SetUseYn(_useYn bool) error {
    r._useYn = _useYn
    r.Set("use_yn", _useYn)
    return nil
}

// UseYn Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetUseYn() bool {
    return r._useYn
}
// ExtendFields Setter
// 拓展属性
func (r *TaobaoWlbWmsSkuUpdateRequest) SetExtendFields(_extendFields string) error {
    r._extendFields = _extendFields
    r.Set("extend_fields", _extendFields)
    return nil
}

// ExtendFields Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetExtendFields() string {
    return r._extendFields
}
// BarCode Setter
// 条形码，多条码请用”;”分隔；条码更新只技术增量更新，已有条码无法修改，只能在原条码基础上增加新的条码。例：原商品条码为：BAR001，要增加一条新条码BAR002时，此字段内容应填写为：BAR001;BAR002
func (r *TaobaoWlbWmsSkuUpdateRequest) SetBarCode(_barCode string) error {
    r._barCode = _barCode
    r.Set("bar_code", _barCode)
    return nil
}

// BarCode Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetBarCode() string {
    return r._barCode
}
// Attribute Setter
// 商品属性编码
func (r *TaobaoWlbWmsSkuUpdateRequest) SetAttribute(_attribute string) error {
    r._attribute = _attribute
    r.Set("attribute", _attribute)
    return nil
}

// Attribute Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetAttribute() string {
    return r._attribute
}
// Type Setter
// 商品类型:NORMAL-普通商品、 COMBINE-组合商品、 DISTRIBUTION-分销商品
func (r *TaobaoWlbWmsSkuUpdateRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetType() string {
    return r._type
}
// IsAreaSale Setter
// 是否区域销售属性
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsAreaSale(_isAreaSale bool) error {
    r._isAreaSale = _isAreaSale
    r.Set("is_area_sale", _isAreaSale)
    return nil
}

// IsAreaSale Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsAreaSale() bool {
    return r._isAreaSale
}
