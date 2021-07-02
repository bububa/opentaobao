package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsSkuCreateAPIRequest 商品同步 API请求
// taobao.wlb.wms.sku.create
//
// 商品同步
type TaobaoWlbWmsSkuCreateAPIRequest struct {
	model.Params
	// 商家商品编码
	_itemCode string
	// 条形码，多条码请用”；”分隔；
	_barCode string
	// 仓库编码
	_storeCode string
	// 商品名称
	_name string
	// 商品标题
	_title string
	// 商品类别NORMAL：普通商品、COMBINE：组合商品、DISTRIBUTION：分销商品、HAOCAI耗材、FUSHUPIN附属品、BAOCAI 包材、XUNI虚拟商品、QITA其他)
	_type string
	// 商品类别编码（外部系统类别）
	_category string
	// 商品类别名称
	_categoryName string
	// 品牌编码
	_brand string
	// 品牌名称
	_brandName string
	// 规格
	_specification string
	// 颜色
	_color string
	// 尺码
	_size string
	// 毛重，单位克
	_grossWeight int64
	// 净重，单位克
	_netWeight int64
	// 长度，单位毫米
	_length int64
	// 宽度，单位毫米
	_width int64
	// 高度，单位毫米
	_height int64
	// 体积，单位立方厘米
	_volume int64
	// 箱规
	_pcs int64
	// 产地
	_originAddress int64
	// 批准文号
	_approvalNumber string
	// 是否启用保质期管理
	_isShelflife bool
	// 商品保质期天数
	_lifecycle int64
	// 保质期禁收天数
	_rejectLifecycle int64
	// 保质期禁售天数
	_lockupLifecycle int64
	// 保质期预警天数
	_adventLifecycle int64
	// 是否启用序列号管理
	_isSnMgt bool
	// 是否易碎品
	_isHygroscopic bool
	// 是否危险品
	_isDanger bool
	// 吊牌价，单位分
	_tagPrice int64
	// 零售价，单位分
	_itemPrice int64
	// 成本价，单位分
	_costPrice int64
	// 是否启用批次管理
	_isBatchMgt bool
	// 启用标识
	_useYn bool
	// 拓展属性
	_extendFields string
	// 商家商品ID
	_itemId string
	// 是否区域销售
	_isAreaSale bool
}

// NewTaobaoWlbWmsSkuCreateRequest 初始化TaobaoWlbWmsSkuCreateAPIRequest对象
func NewTaobaoWlbWmsSkuCreateRequest() *TaobaoWlbWmsSkuCreateAPIRequest {
	return &TaobaoWlbWmsSkuCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.sku.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemCode is ItemCode Setter
// 商家商品编码
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetItemCode(_itemCode string) error {
	r._itemCode = _itemCode
	r.Set("item_code", _itemCode)
	return nil
}

// GetItemCode ItemCode Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetItemCode() string {
	return r._itemCode
}

// SetBarCode is BarCode Setter
// 条形码，多条码请用”；”分隔；
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// GetBarCode BarCode Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetBarCode() string {
	return r._barCode
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetName is Name Setter
// 商品名称
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetName() string {
	return r._name
}

// SetTitle is Title Setter
// 商品标题
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetTitle() string {
	return r._title
}

// SetType is Type Setter
// 商品类别NORMAL：普通商品、COMBINE：组合商品、DISTRIBUTION：分销商品、HAOCAI耗材、FUSHUPIN附属品、BAOCAI 包材、XUNI虚拟商品、QITA其他)
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetType() string {
	return r._type
}

// SetCategory is Category Setter
// 商品类别编码（外部系统类别）
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetCategory(_category string) error {
	r._category = _category
	r.Set("category", _category)
	return nil
}

// GetCategory Category Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetCategory() string {
	return r._category
}

// SetCategoryName is CategoryName Setter
// 商品类别名称
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetCategoryName(_categoryName string) error {
	r._categoryName = _categoryName
	r.Set("category_name", _categoryName)
	return nil
}

// GetCategoryName CategoryName Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetCategoryName() string {
	return r._categoryName
}

// SetBrand is Brand Setter
// 品牌编码
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetBrand(_brand string) error {
	r._brand = _brand
	r.Set("brand", _brand)
	return nil
}

// GetBrand Brand Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetBrand() string {
	return r._brand
}

// SetBrandName is BrandName Setter
// 品牌名称
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetBrandName(_brandName string) error {
	r._brandName = _brandName
	r.Set("brand_name", _brandName)
	return nil
}

// GetBrandName BrandName Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetBrandName() string {
	return r._brandName
}

// SetSpecification is Specification Setter
// 规格
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetSpecification(_specification string) error {
	r._specification = _specification
	r.Set("specification", _specification)
	return nil
}

// GetSpecification Specification Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetSpecification() string {
	return r._specification
}

// SetColor is Color Setter
// 颜色
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetColor(_color string) error {
	r._color = _color
	r.Set("color", _color)
	return nil
}

// GetColor Color Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetColor() string {
	return r._color
}

// SetSize is Size Setter
// 尺码
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetSize(_size string) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetSize() string {
	return r._size
}

// SetGrossWeight is GrossWeight Setter
// 毛重，单位克
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetGrossWeight(_grossWeight int64) error {
	r._grossWeight = _grossWeight
	r.Set("gross_weight", _grossWeight)
	return nil
}

// GetGrossWeight GrossWeight Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetGrossWeight() int64 {
	return r._grossWeight
}

// SetNetWeight is NetWeight Setter
// 净重，单位克
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetNetWeight(_netWeight int64) error {
	r._netWeight = _netWeight
	r.Set("net_weight", _netWeight)
	return nil
}

// GetNetWeight NetWeight Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetNetWeight() int64 {
	return r._netWeight
}

// SetLength is Length Setter
// 长度，单位毫米
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetLength(_length int64) error {
	r._length = _length
	r.Set("length", _length)
	return nil
}

// GetLength Length Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetLength() int64 {
	return r._length
}

// SetWidth is Width Setter
// 宽度，单位毫米
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetWidth(_width int64) error {
	r._width = _width
	r.Set("width", _width)
	return nil
}

// GetWidth Width Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetWidth() int64 {
	return r._width
}

// SetHeight is Height Setter
// 高度，单位毫米
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// GetHeight Height Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetHeight() int64 {
	return r._height
}

// SetVolume is Volume Setter
// 体积，单位立方厘米
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetVolume(_volume int64) error {
	r._volume = _volume
	r.Set("volume", _volume)
	return nil
}

// GetVolume Volume Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetVolume() int64 {
	return r._volume
}

// SetPcs is Pcs Setter
// 箱规
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetPcs(_pcs int64) error {
	r._pcs = _pcs
	r.Set("pcs", _pcs)
	return nil
}

// GetPcs Pcs Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetPcs() int64 {
	return r._pcs
}

// SetOriginAddress is OriginAddress Setter
// 产地
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetOriginAddress(_originAddress int64) error {
	r._originAddress = _originAddress
	r.Set("origin_address", _originAddress)
	return nil
}

// GetOriginAddress OriginAddress Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetOriginAddress() int64 {
	return r._originAddress
}

// SetApprovalNumber is ApprovalNumber Setter
// 批准文号
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetApprovalNumber(_approvalNumber string) error {
	r._approvalNumber = _approvalNumber
	r.Set("approval_number", _approvalNumber)
	return nil
}

// GetApprovalNumber ApprovalNumber Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetApprovalNumber() string {
	return r._approvalNumber
}

// SetIsShelflife is IsShelflife Setter
// 是否启用保质期管理
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsShelflife(_isShelflife bool) error {
	r._isShelflife = _isShelflife
	r.Set("is_shelflife", _isShelflife)
	return nil
}

// GetIsShelflife IsShelflife Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsShelflife() bool {
	return r._isShelflife
}

// SetLifecycle is Lifecycle Setter
// 商品保质期天数
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetLifecycle(_lifecycle int64) error {
	r._lifecycle = _lifecycle
	r.Set("lifecycle", _lifecycle)
	return nil
}

// GetLifecycle Lifecycle Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetLifecycle() int64 {
	return r._lifecycle
}

// SetRejectLifecycle is RejectLifecycle Setter
// 保质期禁收天数
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetRejectLifecycle(_rejectLifecycle int64) error {
	r._rejectLifecycle = _rejectLifecycle
	r.Set("reject_lifecycle", _rejectLifecycle)
	return nil
}

// GetRejectLifecycle RejectLifecycle Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetRejectLifecycle() int64 {
	return r._rejectLifecycle
}

// SetLockupLifecycle is LockupLifecycle Setter
// 保质期禁售天数
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetLockupLifecycle(_lockupLifecycle int64) error {
	r._lockupLifecycle = _lockupLifecycle
	r.Set("lockup_lifecycle", _lockupLifecycle)
	return nil
}

// GetLockupLifecycle LockupLifecycle Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetLockupLifecycle() int64 {
	return r._lockupLifecycle
}

// SetAdventLifecycle is AdventLifecycle Setter
// 保质期预警天数
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetAdventLifecycle(_adventLifecycle int64) error {
	r._adventLifecycle = _adventLifecycle
	r.Set("advent_lifecycle", _adventLifecycle)
	return nil
}

// GetAdventLifecycle AdventLifecycle Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetAdventLifecycle() int64 {
	return r._adventLifecycle
}

// SetIsSnMgt is IsSnMgt Setter
// 是否启用序列号管理
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsSnMgt(_isSnMgt bool) error {
	r._isSnMgt = _isSnMgt
	r.Set("is_sn_mgt", _isSnMgt)
	return nil
}

// GetIsSnMgt IsSnMgt Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsSnMgt() bool {
	return r._isSnMgt
}

// SetIsHygroscopic is IsHygroscopic Setter
// 是否易碎品
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsHygroscopic(_isHygroscopic bool) error {
	r._isHygroscopic = _isHygroscopic
	r.Set("is_hygroscopic", _isHygroscopic)
	return nil
}

// GetIsHygroscopic IsHygroscopic Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsHygroscopic() bool {
	return r._isHygroscopic
}

// SetIsDanger is IsDanger Setter
// 是否危险品
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsDanger(_isDanger bool) error {
	r._isDanger = _isDanger
	r.Set("is_danger", _isDanger)
	return nil
}

// GetIsDanger IsDanger Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsDanger() bool {
	return r._isDanger
}

// SetTagPrice is TagPrice Setter
// 吊牌价，单位分
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetTagPrice(_tagPrice int64) error {
	r._tagPrice = _tagPrice
	r.Set("tag_price", _tagPrice)
	return nil
}

// GetTagPrice TagPrice Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetTagPrice() int64 {
	return r._tagPrice
}

// SetItemPrice is ItemPrice Setter
// 零售价，单位分
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetItemPrice(_itemPrice int64) error {
	r._itemPrice = _itemPrice
	r.Set("item_price", _itemPrice)
	return nil
}

// GetItemPrice ItemPrice Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetItemPrice() int64 {
	return r._itemPrice
}

// SetCostPrice is CostPrice Setter
// 成本价，单位分
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetCostPrice(_costPrice int64) error {
	r._costPrice = _costPrice
	r.Set("cost_price", _costPrice)
	return nil
}

// GetCostPrice CostPrice Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetCostPrice() int64 {
	return r._costPrice
}

// SetIsBatchMgt is IsBatchMgt Setter
// 是否启用批次管理
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsBatchMgt(_isBatchMgt bool) error {
	r._isBatchMgt = _isBatchMgt
	r.Set("is_batch_mgt", _isBatchMgt)
	return nil
}

// GetIsBatchMgt IsBatchMgt Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsBatchMgt() bool {
	return r._isBatchMgt
}

// SetUseYn is UseYn Setter
// 启用标识
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetUseYn(_useYn bool) error {
	r._useYn = _useYn
	r.Set("use_yn", _useYn)
	return nil
}

// GetUseYn UseYn Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetUseYn() bool {
	return r._useYn
}

// SetExtendFields is ExtendFields Setter
// 拓展属性
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetExtendFields(_extendFields string) error {
	r._extendFields = _extendFields
	r.Set("extend_fields", _extendFields)
	return nil
}

// GetExtendFields ExtendFields Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetExtendFields() string {
	return r._extendFields
}

// SetItemId is ItemId Setter
// 商家商品ID
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetItemId() string {
	return r._itemId
}

// SetIsAreaSale is IsAreaSale Setter
// 是否区域销售
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsAreaSale(_isAreaSale bool) error {
	r._isAreaSale = _isAreaSale
	r.Set("is_area_sale", _isAreaSale)
	return nil
}

// GetIsAreaSale IsAreaSale Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsAreaSale() bool {
	return r._isAreaSale
}
