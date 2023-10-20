package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmsskuupdateAPIRequest 商品信息的更新 API请求
// taobao.wlb.wms.sku.update
//
// 商品信息的更新
type TaobaowlbwmsskuupdateAPIRequest struct {
	model.Params
	// 外部系统ID
	_itemId string
	// 仓库编码
	_storeCode string
	// 商品名称
	_name string
	// 商品标题
	_title string
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
	// 批准文号
	_approvalNumber string
	// 拓展属性
	_extendFields string
	// 条形码，多条码请用”;”分隔；条码更新只技术增量更新，已有条码无法修改，只能在原条码基础上增加新的条码。例：原商品条码为：BAR001，要增加一条新条码BAR002时，此字段内容应填写为：BAR001;BAR002
	_barCode string
	// 商品属性编码
	_attribute string
	// 商品类型:NORMAL-普通商品、 COMBINE-组合商品、 DISTRIBUTION-分销商品
	_type string
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
	// 商品保质期天数
	_lifecycle int64
	// 保质期禁收天数
	_rejectLifecycle int64
	// 保质期禁售天数
	_lockupLifecycle int64
	// 保质期预警天数
	_adventLifecycle int64
	// 吊牌价，单位分
	_tagPrice int64
	// 零售价，单位分
	_itemPrice int64
	// 成本价，单位分
	_costPrice int64
	// 是否启用保质期管理
	_isShelflife bool
	// 是否启用序列号管理
	_isSnMgt bool
	// 是否易碎品
	_isHygroscopic bool
	// 是否危险品
	_isDanger bool
	// 是否启用批次管理
	_isBatchMgt bool
	// 启用标识
	_useYn bool
	// 是否区域销售属性
	_isAreaSale bool
}

// NewTaobaowlbwmsskuupdateRequest 初始化TaobaowlbwmsskuupdateAPIRequest对象
func NewTaobaowlbwmsskuupdateRequest() *TaobaowlbwmsskuupdateAPIRequest {
	return &TaobaowlbwmsskuupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmsskuupdateAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.sku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmsskuupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmsskuupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 外部系统ID
func (r *TaobaowlbwmsskuupdateAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetItemId() string {
	return r._itemId
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaowlbwmsskuupdateAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetName is Name Setter
// 商品名称
func (r *TaobaowlbwmsskuupdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetName() string {
	return r._name
}

// SetTitle is Title Setter
// 商品标题
func (r *TaobaowlbwmsskuupdateAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetTitle() string {
	return r._title
}

// SetCategory is Category Setter
// 商品类别编码（外部系统类别）
func (r *TaobaowlbwmsskuupdateAPIRequest) SetCategory(_category string) error {
	r._category = _category
	r.Set("category", _category)
	return nil
}

// GetCategory Category Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetCategory() string {
	return r._category
}

// SetCategoryName is CategoryName Setter
// 商品类别名称
func (r *TaobaowlbwmsskuupdateAPIRequest) SetCategoryName(_categoryName string) error {
	r._categoryName = _categoryName
	r.Set("category_name", _categoryName)
	return nil
}

// GetCategoryName CategoryName Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetCategoryName() string {
	return r._categoryName
}

// SetBrand is Brand Setter
// 品牌编码
func (r *TaobaowlbwmsskuupdateAPIRequest) SetBrand(_brand string) error {
	r._brand = _brand
	r.Set("brand", _brand)
	return nil
}

// GetBrand Brand Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetBrand() string {
	return r._brand
}

// SetBrandName is BrandName Setter
// 品牌名称
func (r *TaobaowlbwmsskuupdateAPIRequest) SetBrandName(_brandName string) error {
	r._brandName = _brandName
	r.Set("brand_name", _brandName)
	return nil
}

// GetBrandName BrandName Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetBrandName() string {
	return r._brandName
}

// SetSpecification is Specification Setter
// 规格
func (r *TaobaowlbwmsskuupdateAPIRequest) SetSpecification(_specification string) error {
	r._specification = _specification
	r.Set("specification", _specification)
	return nil
}

// GetSpecification Specification Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetSpecification() string {
	return r._specification
}

// SetColor is Color Setter
// 颜色
func (r *TaobaowlbwmsskuupdateAPIRequest) SetColor(_color string) error {
	r._color = _color
	r.Set("color", _color)
	return nil
}

// GetColor Color Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetColor() string {
	return r._color
}

// SetSize is Size Setter
// 尺码
func (r *TaobaowlbwmsskuupdateAPIRequest) SetSize(_size string) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetSize() string {
	return r._size
}

// SetApprovalNumber is ApprovalNumber Setter
// 批准文号
func (r *TaobaowlbwmsskuupdateAPIRequest) SetApprovalNumber(_approvalNumber string) error {
	r._approvalNumber = _approvalNumber
	r.Set("approval_number", _approvalNumber)
	return nil
}

// GetApprovalNumber ApprovalNumber Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetApprovalNumber() string {
	return r._approvalNumber
}

// SetExtendFields is ExtendFields Setter
// 拓展属性
func (r *TaobaowlbwmsskuupdateAPIRequest) SetExtendFields(_extendFields string) error {
	r._extendFields = _extendFields
	r.Set("extend_fields", _extendFields)
	return nil
}

// GetExtendFields ExtendFields Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetExtendFields() string {
	return r._extendFields
}

// SetBarCode is BarCode Setter
// 条形码，多条码请用”;”分隔；条码更新只技术增量更新，已有条码无法修改，只能在原条码基础上增加新的条码。例：原商品条码为：BAR001，要增加一条新条码BAR002时，此字段内容应填写为：BAR001;BAR002
func (r *TaobaowlbwmsskuupdateAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// GetBarCode BarCode Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetBarCode() string {
	return r._barCode
}

// SetAttribute is Attribute Setter
// 商品属性编码
func (r *TaobaowlbwmsskuupdateAPIRequest) SetAttribute(_attribute string) error {
	r._attribute = _attribute
	r.Set("attribute", _attribute)
	return nil
}

// GetAttribute Attribute Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetAttribute() string {
	return r._attribute
}

// SetType is Type Setter
// 商品类型:NORMAL-普通商品、 COMBINE-组合商品、 DISTRIBUTION-分销商品
func (r *TaobaowlbwmsskuupdateAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetType() string {
	return r._type
}

// SetGrossWeight is GrossWeight Setter
// 毛重，单位克
func (r *TaobaowlbwmsskuupdateAPIRequest) SetGrossWeight(_grossWeight int64) error {
	r._grossWeight = _grossWeight
	r.Set("gross_weight", _grossWeight)
	return nil
}

// GetGrossWeight GrossWeight Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetGrossWeight() int64 {
	return r._grossWeight
}

// SetNetWeight is NetWeight Setter
// 净重，单位克
func (r *TaobaowlbwmsskuupdateAPIRequest) SetNetWeight(_netWeight int64) error {
	r._netWeight = _netWeight
	r.Set("net_weight", _netWeight)
	return nil
}

// GetNetWeight NetWeight Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetNetWeight() int64 {
	return r._netWeight
}

// SetLength is Length Setter
// 长度，单位毫米
func (r *TaobaowlbwmsskuupdateAPIRequest) SetLength(_length int64) error {
	r._length = _length
	r.Set("length", _length)
	return nil
}

// GetLength Length Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetLength() int64 {
	return r._length
}

// SetWidth is Width Setter
// 宽度，单位毫米
func (r *TaobaowlbwmsskuupdateAPIRequest) SetWidth(_width int64) error {
	r._width = _width
	r.Set("width", _width)
	return nil
}

// GetWidth Width Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetWidth() int64 {
	return r._width
}

// SetHeight is Height Setter
// 高度，单位毫米
func (r *TaobaowlbwmsskuupdateAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// GetHeight Height Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetHeight() int64 {
	return r._height
}

// SetVolume is Volume Setter
// 体积，单位立方厘米
func (r *TaobaowlbwmsskuupdateAPIRequest) SetVolume(_volume int64) error {
	r._volume = _volume
	r.Set("volume", _volume)
	return nil
}

// GetVolume Volume Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetVolume() int64 {
	return r._volume
}

// SetPcs is Pcs Setter
// 箱规
func (r *TaobaowlbwmsskuupdateAPIRequest) SetPcs(_pcs int64) error {
	r._pcs = _pcs
	r.Set("pcs", _pcs)
	return nil
}

// GetPcs Pcs Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetPcs() int64 {
	return r._pcs
}

// SetOriginAddress is OriginAddress Setter
// 产地
func (r *TaobaowlbwmsskuupdateAPIRequest) SetOriginAddress(_originAddress int64) error {
	r._originAddress = _originAddress
	r.Set("origin_address", _originAddress)
	return nil
}

// GetOriginAddress OriginAddress Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetOriginAddress() int64 {
	return r._originAddress
}

// SetLifecycle is Lifecycle Setter
// 商品保质期天数
func (r *TaobaowlbwmsskuupdateAPIRequest) SetLifecycle(_lifecycle int64) error {
	r._lifecycle = _lifecycle
	r.Set("lifecycle", _lifecycle)
	return nil
}

// GetLifecycle Lifecycle Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetLifecycle() int64 {
	return r._lifecycle
}

// SetRejectLifecycle is RejectLifecycle Setter
// 保质期禁收天数
func (r *TaobaowlbwmsskuupdateAPIRequest) SetRejectLifecycle(_rejectLifecycle int64) error {
	r._rejectLifecycle = _rejectLifecycle
	r.Set("reject_lifecycle", _rejectLifecycle)
	return nil
}

// GetRejectLifecycle RejectLifecycle Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetRejectLifecycle() int64 {
	return r._rejectLifecycle
}

// SetLockupLifecycle is LockupLifecycle Setter
// 保质期禁售天数
func (r *TaobaowlbwmsskuupdateAPIRequest) SetLockupLifecycle(_lockupLifecycle int64) error {
	r._lockupLifecycle = _lockupLifecycle
	r.Set("lockup_lifecycle", _lockupLifecycle)
	return nil
}

// GetLockupLifecycle LockupLifecycle Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetLockupLifecycle() int64 {
	return r._lockupLifecycle
}

// SetAdventLifecycle is AdventLifecycle Setter
// 保质期预警天数
func (r *TaobaowlbwmsskuupdateAPIRequest) SetAdventLifecycle(_adventLifecycle int64) error {
	r._adventLifecycle = _adventLifecycle
	r.Set("advent_lifecycle", _adventLifecycle)
	return nil
}

// GetAdventLifecycle AdventLifecycle Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetAdventLifecycle() int64 {
	return r._adventLifecycle
}

// SetTagPrice is TagPrice Setter
// 吊牌价，单位分
func (r *TaobaowlbwmsskuupdateAPIRequest) SetTagPrice(_tagPrice int64) error {
	r._tagPrice = _tagPrice
	r.Set("tag_price", _tagPrice)
	return nil
}

// GetTagPrice TagPrice Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetTagPrice() int64 {
	return r._tagPrice
}

// SetItemPrice is ItemPrice Setter
// 零售价，单位分
func (r *TaobaowlbwmsskuupdateAPIRequest) SetItemPrice(_itemPrice int64) error {
	r._itemPrice = _itemPrice
	r.Set("item_price", _itemPrice)
	return nil
}

// GetItemPrice ItemPrice Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetItemPrice() int64 {
	return r._itemPrice
}

// SetCostPrice is CostPrice Setter
// 成本价，单位分
func (r *TaobaowlbwmsskuupdateAPIRequest) SetCostPrice(_costPrice int64) error {
	r._costPrice = _costPrice
	r.Set("cost_price", _costPrice)
	return nil
}

// GetCostPrice CostPrice Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetCostPrice() int64 {
	return r._costPrice
}

// SetIsShelflife is IsShelflife Setter
// 是否启用保质期管理
func (r *TaobaowlbwmsskuupdateAPIRequest) SetIsShelflife(_isShelflife bool) error {
	r._isShelflife = _isShelflife
	r.Set("is_shelflife", _isShelflife)
	return nil
}

// GetIsShelflife IsShelflife Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetIsShelflife() bool {
	return r._isShelflife
}

// SetIsSnMgt is IsSnMgt Setter
// 是否启用序列号管理
func (r *TaobaowlbwmsskuupdateAPIRequest) SetIsSnMgt(_isSnMgt bool) error {
	r._isSnMgt = _isSnMgt
	r.Set("is_sn_mgt", _isSnMgt)
	return nil
}

// GetIsSnMgt IsSnMgt Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetIsSnMgt() bool {
	return r._isSnMgt
}

// SetIsHygroscopic is IsHygroscopic Setter
// 是否易碎品
func (r *TaobaowlbwmsskuupdateAPIRequest) SetIsHygroscopic(_isHygroscopic bool) error {
	r._isHygroscopic = _isHygroscopic
	r.Set("is_hygroscopic", _isHygroscopic)
	return nil
}

// GetIsHygroscopic IsHygroscopic Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetIsHygroscopic() bool {
	return r._isHygroscopic
}

// SetIsDanger is IsDanger Setter
// 是否危险品
func (r *TaobaowlbwmsskuupdateAPIRequest) SetIsDanger(_isDanger bool) error {
	r._isDanger = _isDanger
	r.Set("is_danger", _isDanger)
	return nil
}

// GetIsDanger IsDanger Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetIsDanger() bool {
	return r._isDanger
}

// SetIsBatchMgt is IsBatchMgt Setter
// 是否启用批次管理
func (r *TaobaowlbwmsskuupdateAPIRequest) SetIsBatchMgt(_isBatchMgt bool) error {
	r._isBatchMgt = _isBatchMgt
	r.Set("is_batch_mgt", _isBatchMgt)
	return nil
}

// GetIsBatchMgt IsBatchMgt Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetIsBatchMgt() bool {
	return r._isBatchMgt
}

// SetUseYn is UseYn Setter
// 启用标识
func (r *TaobaowlbwmsskuupdateAPIRequest) SetUseYn(_useYn bool) error {
	r._useYn = _useYn
	r.Set("use_yn", _useYn)
	return nil
}

// GetUseYn UseYn Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetUseYn() bool {
	return r._useYn
}

// SetIsAreaSale is IsAreaSale Setter
// 是否区域销售属性
func (r *TaobaowlbwmsskuupdateAPIRequest) SetIsAreaSale(_isAreaSale bool) error {
	r._isAreaSale = _isAreaSale
	r.Set("is_area_sale", _isAreaSale)
	return nil
}

// GetIsAreaSale IsAreaSale Getter
func (r TaobaowlbwmsskuupdateAPIRequest) GetIsAreaSale() bool {
	return r._isAreaSale
}
