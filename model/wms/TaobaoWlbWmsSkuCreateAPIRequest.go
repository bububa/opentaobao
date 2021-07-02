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

// Set is ItemCode Setter
// 商家商品编码
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetItemCode(_itemCode string) error {
	r._itemCode = _itemCode
	r.Set("item_code", _itemCode)
	return nil
}

// Get ItemCode Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetItemCode() string {
	return r._itemCode
}

// Set is BarCode Setter
// 条形码，多条码请用”；”分隔；
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// Get BarCode Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetBarCode() string {
	return r._barCode
}

// Set is StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// Get StoreCode Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// Set is Name Setter
// 商品名称
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetName() string {
	return r._name
}

// Set is Title Setter
// 商品标题
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// Get Title Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetTitle() string {
	return r._title
}

// Set is Type Setter
// 商品类别NORMAL：普通商品、COMBINE：组合商品、DISTRIBUTION：分销商品、HAOCAI耗材、FUSHUPIN附属品、BAOCAI 包材、XUNI虚拟商品、QITA其他)
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetType() string {
	return r._type
}

// Set is Category Setter
// 商品类别编码（外部系统类别）
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetCategory(_category string) error {
	r._category = _category
	r.Set("category", _category)
	return nil
}

// Get Category Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetCategory() string {
	return r._category
}

// Set is CategoryName Setter
// 商品类别名称
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetCategoryName(_categoryName string) error {
	r._categoryName = _categoryName
	r.Set("category_name", _categoryName)
	return nil
}

// Get CategoryName Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetCategoryName() string {
	return r._categoryName
}

// Set is Brand Setter
// 品牌编码
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetBrand(_brand string) error {
	r._brand = _brand
	r.Set("brand", _brand)
	return nil
}

// Get Brand Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetBrand() string {
	return r._brand
}

// Set is BrandName Setter
// 品牌名称
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetBrandName(_brandName string) error {
	r._brandName = _brandName
	r.Set("brand_name", _brandName)
	return nil
}

// Get BrandName Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetBrandName() string {
	return r._brandName
}

// Set is Specification Setter
// 规格
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetSpecification(_specification string) error {
	r._specification = _specification
	r.Set("specification", _specification)
	return nil
}

// Get Specification Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetSpecification() string {
	return r._specification
}

// Set is Color Setter
// 颜色
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetColor(_color string) error {
	r._color = _color
	r.Set("color", _color)
	return nil
}

// Get Color Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetColor() string {
	return r._color
}

// Set is Size Setter
// 尺码
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetSize(_size string) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// Get Size Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetSize() string {
	return r._size
}

// Set is GrossWeight Setter
// 毛重，单位克
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetGrossWeight(_grossWeight int64) error {
	r._grossWeight = _grossWeight
	r.Set("gross_weight", _grossWeight)
	return nil
}

// Get GrossWeight Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetGrossWeight() int64 {
	return r._grossWeight
}

// Set is NetWeight Setter
// 净重，单位克
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetNetWeight(_netWeight int64) error {
	r._netWeight = _netWeight
	r.Set("net_weight", _netWeight)
	return nil
}

// Get NetWeight Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetNetWeight() int64 {
	return r._netWeight
}

// Set is Length Setter
// 长度，单位毫米
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetLength(_length int64) error {
	r._length = _length
	r.Set("length", _length)
	return nil
}

// Get Length Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetLength() int64 {
	return r._length
}

// Set is Width Setter
// 宽度，单位毫米
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetWidth(_width int64) error {
	r._width = _width
	r.Set("width", _width)
	return nil
}

// Get Width Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetWidth() int64 {
	return r._width
}

// Set is Height Setter
// 高度，单位毫米
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// Get Height Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetHeight() int64 {
	return r._height
}

// Set is Volume Setter
// 体积，单位立方厘米
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetVolume(_volume int64) error {
	r._volume = _volume
	r.Set("volume", _volume)
	return nil
}

// Get Volume Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetVolume() int64 {
	return r._volume
}

// Set is Pcs Setter
// 箱规
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetPcs(_pcs int64) error {
	r._pcs = _pcs
	r.Set("pcs", _pcs)
	return nil
}

// Get Pcs Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetPcs() int64 {
	return r._pcs
}

// Set is OriginAddress Setter
// 产地
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetOriginAddress(_originAddress int64) error {
	r._originAddress = _originAddress
	r.Set("origin_address", _originAddress)
	return nil
}

// Get OriginAddress Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetOriginAddress() int64 {
	return r._originAddress
}

// Set is ApprovalNumber Setter
// 批准文号
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetApprovalNumber(_approvalNumber string) error {
	r._approvalNumber = _approvalNumber
	r.Set("approval_number", _approvalNumber)
	return nil
}

// Get ApprovalNumber Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetApprovalNumber() string {
	return r._approvalNumber
}

// Set is IsShelflife Setter
// 是否启用保质期管理
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsShelflife(_isShelflife bool) error {
	r._isShelflife = _isShelflife
	r.Set("is_shelflife", _isShelflife)
	return nil
}

// Get IsShelflife Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsShelflife() bool {
	return r._isShelflife
}

// Set is Lifecycle Setter
// 商品保质期天数
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetLifecycle(_lifecycle int64) error {
	r._lifecycle = _lifecycle
	r.Set("lifecycle", _lifecycle)
	return nil
}

// Get Lifecycle Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetLifecycle() int64 {
	return r._lifecycle
}

// Set is RejectLifecycle Setter
// 保质期禁收天数
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetRejectLifecycle(_rejectLifecycle int64) error {
	r._rejectLifecycle = _rejectLifecycle
	r.Set("reject_lifecycle", _rejectLifecycle)
	return nil
}

// Get RejectLifecycle Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetRejectLifecycle() int64 {
	return r._rejectLifecycle
}

// Set is LockupLifecycle Setter
// 保质期禁售天数
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetLockupLifecycle(_lockupLifecycle int64) error {
	r._lockupLifecycle = _lockupLifecycle
	r.Set("lockup_lifecycle", _lockupLifecycle)
	return nil
}

// Get LockupLifecycle Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetLockupLifecycle() int64 {
	return r._lockupLifecycle
}

// Set is AdventLifecycle Setter
// 保质期预警天数
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetAdventLifecycle(_adventLifecycle int64) error {
	r._adventLifecycle = _adventLifecycle
	r.Set("advent_lifecycle", _adventLifecycle)
	return nil
}

// Get AdventLifecycle Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetAdventLifecycle() int64 {
	return r._adventLifecycle
}

// Set is IsSnMgt Setter
// 是否启用序列号管理
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsSnMgt(_isSnMgt bool) error {
	r._isSnMgt = _isSnMgt
	r.Set("is_sn_mgt", _isSnMgt)
	return nil
}

// Get IsSnMgt Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsSnMgt() bool {
	return r._isSnMgt
}

// Set is IsHygroscopic Setter
// 是否易碎品
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsHygroscopic(_isHygroscopic bool) error {
	r._isHygroscopic = _isHygroscopic
	r.Set("is_hygroscopic", _isHygroscopic)
	return nil
}

// Get IsHygroscopic Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsHygroscopic() bool {
	return r._isHygroscopic
}

// Set is IsDanger Setter
// 是否危险品
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsDanger(_isDanger bool) error {
	r._isDanger = _isDanger
	r.Set("is_danger", _isDanger)
	return nil
}

// Get IsDanger Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsDanger() bool {
	return r._isDanger
}

// Set is TagPrice Setter
// 吊牌价，单位分
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetTagPrice(_tagPrice int64) error {
	r._tagPrice = _tagPrice
	r.Set("tag_price", _tagPrice)
	return nil
}

// Get TagPrice Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetTagPrice() int64 {
	return r._tagPrice
}

// Set is ItemPrice Setter
// 零售价，单位分
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetItemPrice(_itemPrice int64) error {
	r._itemPrice = _itemPrice
	r.Set("item_price", _itemPrice)
	return nil
}

// Get ItemPrice Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetItemPrice() int64 {
	return r._itemPrice
}

// Set is CostPrice Setter
// 成本价，单位分
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetCostPrice(_costPrice int64) error {
	r._costPrice = _costPrice
	r.Set("cost_price", _costPrice)
	return nil
}

// Get CostPrice Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetCostPrice() int64 {
	return r._costPrice
}

// Set is IsBatchMgt Setter
// 是否启用批次管理
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsBatchMgt(_isBatchMgt bool) error {
	r._isBatchMgt = _isBatchMgt
	r.Set("is_batch_mgt", _isBatchMgt)
	return nil
}

// Get IsBatchMgt Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsBatchMgt() bool {
	return r._isBatchMgt
}

// Set is UseYn Setter
// 启用标识
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetUseYn(_useYn bool) error {
	r._useYn = _useYn
	r.Set("use_yn", _useYn)
	return nil
}

// Get UseYn Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetUseYn() bool {
	return r._useYn
}

// Set is ExtendFields Setter
// 拓展属性
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetExtendFields(_extendFields string) error {
	r._extendFields = _extendFields
	r.Set("extend_fields", _extendFields)
	return nil
}

// Get ExtendFields Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetExtendFields() string {
	return r._extendFields
}

// Set is ItemId Setter
// 商家商品ID
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetItemId() string {
	return r._itemId
}

// Set is IsAreaSale Setter
// 是否区域销售
func (r *TaobaoWlbWmsSkuCreateAPIRequest) SetIsAreaSale(_isAreaSale bool) error {
	r._isAreaSale = _isAreaSale
	r.Set("is_area_sale", _isAreaSale)
	return nil
}

// Get IsAreaSale Getter
func (r TaobaoWlbWmsSkuCreateAPIRequest) GetIsAreaSale() bool {
	return r._isAreaSale
}
