package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品同步 API请求
taobao.wlb.wms.sku.create

商品同步
*/
type TaobaoWlbWmsSkuCreateRequest struct {
    model.Params
    // 商家商品编码
    _itemCode   string
    // 条形码，多条码请用”；”分隔；
    _barCode   string
    // 仓库编码
    _storeCode   string
    // 商品名称
    _name   string
    // 商品标题
    _title   string
    // 商品类别NORMAL：普通商品、COMBINE：组合商品、DISTRIBUTION：分销商品、HAOCAI耗材、FUSHUPIN附属品、BAOCAI 包材、XUNI虚拟商品、QITA其他)
    _type   string
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
    // 商家商品ID
    _itemId   string
    // 是否区域销售
    _isAreaSale   bool
}

// 初始化TaobaoWlbWmsSkuCreateRequest对象
func NewTaobaoWlbWmsSkuCreateRequest() *TaobaoWlbWmsSkuCreateRequest{
    return &TaobaoWlbWmsSkuCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsSkuCreateRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.sku.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsSkuCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemCode Setter
// 商家商品编码
func (r *TaobaoWlbWmsSkuCreateRequest) SetItemCode(_itemCode string) error {
    r._itemCode = _itemCode
    r.Set("item_code", _itemCode)
    return nil
}

// ItemCode Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetItemCode() string {
    return r._itemCode
}
// BarCode Setter
// 条形码，多条码请用”；”分隔；
func (r *TaobaoWlbWmsSkuCreateRequest) SetBarCode(_barCode string) error {
    r._barCode = _barCode
    r.Set("bar_code", _barCode)
    return nil
}

// BarCode Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetBarCode() string {
    return r._barCode
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsSkuCreateRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetStoreCode() string {
    return r._storeCode
}
// Name Setter
// 商品名称
func (r *TaobaoWlbWmsSkuCreateRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetName() string {
    return r._name
}
// Title Setter
// 商品标题
func (r *TaobaoWlbWmsSkuCreateRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetTitle() string {
    return r._title
}
// Type Setter
// 商品类别NORMAL：普通商品、COMBINE：组合商品、DISTRIBUTION：分销商品、HAOCAI耗材、FUSHUPIN附属品、BAOCAI 包材、XUNI虚拟商品、QITA其他)
func (r *TaobaoWlbWmsSkuCreateRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetType() string {
    return r._type
}
// Category Setter
// 商品类别编码（外部系统类别）
func (r *TaobaoWlbWmsSkuCreateRequest) SetCategory(_category string) error {
    r._category = _category
    r.Set("category", _category)
    return nil
}

// Category Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetCategory() string {
    return r._category
}
// CategoryName Setter
// 商品类别名称
func (r *TaobaoWlbWmsSkuCreateRequest) SetCategoryName(_categoryName string) error {
    r._categoryName = _categoryName
    r.Set("category_name", _categoryName)
    return nil
}

// CategoryName Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetCategoryName() string {
    return r._categoryName
}
// Brand Setter
// 品牌编码
func (r *TaobaoWlbWmsSkuCreateRequest) SetBrand(_brand string) error {
    r._brand = _brand
    r.Set("brand", _brand)
    return nil
}

// Brand Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetBrand() string {
    return r._brand
}
// BrandName Setter
// 品牌名称
func (r *TaobaoWlbWmsSkuCreateRequest) SetBrandName(_brandName string) error {
    r._brandName = _brandName
    r.Set("brand_name", _brandName)
    return nil
}

// BrandName Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetBrandName() string {
    return r._brandName
}
// Specification Setter
// 规格
func (r *TaobaoWlbWmsSkuCreateRequest) SetSpecification(_specification string) error {
    r._specification = _specification
    r.Set("specification", _specification)
    return nil
}

// Specification Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetSpecification() string {
    return r._specification
}
// Color Setter
// 颜色
func (r *TaobaoWlbWmsSkuCreateRequest) SetColor(_color string) error {
    r._color = _color
    r.Set("color", _color)
    return nil
}

// Color Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetColor() string {
    return r._color
}
// Size Setter
// 尺码
func (r *TaobaoWlbWmsSkuCreateRequest) SetSize(_size string) error {
    r._size = _size
    r.Set("size", _size)
    return nil
}

// Size Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetSize() string {
    return r._size
}
// GrossWeight Setter
// 毛重，单位克
func (r *TaobaoWlbWmsSkuCreateRequest) SetGrossWeight(_grossWeight int64) error {
    r._grossWeight = _grossWeight
    r.Set("gross_weight", _grossWeight)
    return nil
}

// GrossWeight Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetGrossWeight() int64 {
    return r._grossWeight
}
// NetWeight Setter
// 净重，单位克
func (r *TaobaoWlbWmsSkuCreateRequest) SetNetWeight(_netWeight int64) error {
    r._netWeight = _netWeight
    r.Set("net_weight", _netWeight)
    return nil
}

// NetWeight Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetNetWeight() int64 {
    return r._netWeight
}
// Length Setter
// 长度，单位毫米
func (r *TaobaoWlbWmsSkuCreateRequest) SetLength(_length int64) error {
    r._length = _length
    r.Set("length", _length)
    return nil
}

// Length Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetLength() int64 {
    return r._length
}
// Width Setter
// 宽度，单位毫米
func (r *TaobaoWlbWmsSkuCreateRequest) SetWidth(_width int64) error {
    r._width = _width
    r.Set("width", _width)
    return nil
}

// Width Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetWidth() int64 {
    return r._width
}
// Height Setter
// 高度，单位毫米
func (r *TaobaoWlbWmsSkuCreateRequest) SetHeight(_height int64) error {
    r._height = _height
    r.Set("height", _height)
    return nil
}

// Height Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetHeight() int64 {
    return r._height
}
// Volume Setter
// 体积，单位立方厘米
func (r *TaobaoWlbWmsSkuCreateRequest) SetVolume(_volume int64) error {
    r._volume = _volume
    r.Set("volume", _volume)
    return nil
}

// Volume Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetVolume() int64 {
    return r._volume
}
// Pcs Setter
// 箱规
func (r *TaobaoWlbWmsSkuCreateRequest) SetPcs(_pcs int64) error {
    r._pcs = _pcs
    r.Set("pcs", _pcs)
    return nil
}

// Pcs Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetPcs() int64 {
    return r._pcs
}
// OriginAddress Setter
// 产地
func (r *TaobaoWlbWmsSkuCreateRequest) SetOriginAddress(_originAddress int64) error {
    r._originAddress = _originAddress
    r.Set("origin_address", _originAddress)
    return nil
}

// OriginAddress Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetOriginAddress() int64 {
    return r._originAddress
}
// ApprovalNumber Setter
// 批准文号
func (r *TaobaoWlbWmsSkuCreateRequest) SetApprovalNumber(_approvalNumber string) error {
    r._approvalNumber = _approvalNumber
    r.Set("approval_number", _approvalNumber)
    return nil
}

// ApprovalNumber Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetApprovalNumber() string {
    return r._approvalNumber
}
// IsShelflife Setter
// 是否启用保质期管理
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsShelflife(_isShelflife bool) error {
    r._isShelflife = _isShelflife
    r.Set("is_shelflife", _isShelflife)
    return nil
}

// IsShelflife Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsShelflife() bool {
    return r._isShelflife
}
// Lifecycle Setter
// 商品保质期天数
func (r *TaobaoWlbWmsSkuCreateRequest) SetLifecycle(_lifecycle int64) error {
    r._lifecycle = _lifecycle
    r.Set("lifecycle", _lifecycle)
    return nil
}

// Lifecycle Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetLifecycle() int64 {
    return r._lifecycle
}
// RejectLifecycle Setter
// 保质期禁收天数
func (r *TaobaoWlbWmsSkuCreateRequest) SetRejectLifecycle(_rejectLifecycle int64) error {
    r._rejectLifecycle = _rejectLifecycle
    r.Set("reject_lifecycle", _rejectLifecycle)
    return nil
}

// RejectLifecycle Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetRejectLifecycle() int64 {
    return r._rejectLifecycle
}
// LockupLifecycle Setter
// 保质期禁售天数
func (r *TaobaoWlbWmsSkuCreateRequest) SetLockupLifecycle(_lockupLifecycle int64) error {
    r._lockupLifecycle = _lockupLifecycle
    r.Set("lockup_lifecycle", _lockupLifecycle)
    return nil
}

// LockupLifecycle Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetLockupLifecycle() int64 {
    return r._lockupLifecycle
}
// AdventLifecycle Setter
// 保质期预警天数
func (r *TaobaoWlbWmsSkuCreateRequest) SetAdventLifecycle(_adventLifecycle int64) error {
    r._adventLifecycle = _adventLifecycle
    r.Set("advent_lifecycle", _adventLifecycle)
    return nil
}

// AdventLifecycle Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetAdventLifecycle() int64 {
    return r._adventLifecycle
}
// IsSnMgt Setter
// 是否启用序列号管理
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsSnMgt(_isSnMgt bool) error {
    r._isSnMgt = _isSnMgt
    r.Set("is_sn_mgt", _isSnMgt)
    return nil
}

// IsSnMgt Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsSnMgt() bool {
    return r._isSnMgt
}
// IsHygroscopic Setter
// 是否易碎品
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsHygroscopic(_isHygroscopic bool) error {
    r._isHygroscopic = _isHygroscopic
    r.Set("is_hygroscopic", _isHygroscopic)
    return nil
}

// IsHygroscopic Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsHygroscopic() bool {
    return r._isHygroscopic
}
// IsDanger Setter
// 是否危险品
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsDanger(_isDanger bool) error {
    r._isDanger = _isDanger
    r.Set("is_danger", _isDanger)
    return nil
}

// IsDanger Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsDanger() bool {
    return r._isDanger
}
// TagPrice Setter
// 吊牌价，单位分
func (r *TaobaoWlbWmsSkuCreateRequest) SetTagPrice(_tagPrice int64) error {
    r._tagPrice = _tagPrice
    r.Set("tag_price", _tagPrice)
    return nil
}

// TagPrice Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetTagPrice() int64 {
    return r._tagPrice
}
// ItemPrice Setter
// 零售价，单位分
func (r *TaobaoWlbWmsSkuCreateRequest) SetItemPrice(_itemPrice int64) error {
    r._itemPrice = _itemPrice
    r.Set("item_price", _itemPrice)
    return nil
}

// ItemPrice Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetItemPrice() int64 {
    return r._itemPrice
}
// CostPrice Setter
// 成本价，单位分
func (r *TaobaoWlbWmsSkuCreateRequest) SetCostPrice(_costPrice int64) error {
    r._costPrice = _costPrice
    r.Set("cost_price", _costPrice)
    return nil
}

// CostPrice Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetCostPrice() int64 {
    return r._costPrice
}
// IsBatchMgt Setter
// 是否启用批次管理
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsBatchMgt(_isBatchMgt bool) error {
    r._isBatchMgt = _isBatchMgt
    r.Set("is_batch_mgt", _isBatchMgt)
    return nil
}

// IsBatchMgt Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsBatchMgt() bool {
    return r._isBatchMgt
}
// UseYn Setter
// 启用标识
func (r *TaobaoWlbWmsSkuCreateRequest) SetUseYn(_useYn bool) error {
    r._useYn = _useYn
    r.Set("use_yn", _useYn)
    return nil
}

// UseYn Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetUseYn() bool {
    return r._useYn
}
// ExtendFields Setter
// 拓展属性
func (r *TaobaoWlbWmsSkuCreateRequest) SetExtendFields(_extendFields string) error {
    r._extendFields = _extendFields
    r.Set("extend_fields", _extendFields)
    return nil
}

// ExtendFields Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetExtendFields() string {
    return r._extendFields
}
// ItemId Setter
// 商家商品ID
func (r *TaobaoWlbWmsSkuCreateRequest) SetItemId(_itemId string) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetItemId() string {
    return r._itemId
}
// IsAreaSale Setter
// 是否区域销售
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsAreaSale(_isAreaSale bool) error {
    r._isAreaSale = _isAreaSale
    r.Set("is_area_sale", _isAreaSale)
    return nil
}

// IsAreaSale Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsAreaSale() bool {
    return r._isAreaSale
}
