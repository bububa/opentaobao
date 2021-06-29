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
    itemId   string
    // 仓库编码
    storeCode   string
    // 商品名称
    name   string
    // 商品标题
    title   string
    // 商品类别编码（外部系统类别）
    category   string
    // 商品类别名称
    categoryName   string
    // 品牌编码
    brand   string
    // 品牌名称
    brandName   string
    // 规格
    specification   string
    // 颜色
    color   string
    // 尺码
    size   string
    // 毛重，单位克
    grossWeight   int64
    // 净重，单位克
    netWeight   int64
    // 长度，单位毫米
    length   int64
    // 宽度，单位毫米
    width   int64
    // 高度，单位毫米
    height   int64
    // 体积，单位立方厘米
    volume   int64
    // 箱规
    pcs   int64
    // 产地
    originAddress   int64
    // 批准文号
    approvalNumber   string
    // 是否启用保质期管理
    isShelflife   bool
    // 商品保质期天数
    lifecycle   int64
    // 保质期禁收天数
    rejectLifecycle   int64
    // 保质期禁售天数
    lockupLifecycle   int64
    // 保质期预警天数
    adventLifecycle   int64
    // 是否启用序列号管理
    isSnMgt   bool
    // 是否易碎品
    isHygroscopic   bool
    // 是否危险品
    isDanger   bool
    // 吊牌价，单位分
    tagPrice   int64
    // 零售价，单位分
    itemPrice   int64
    // 成本价，单位分
    costPrice   int64
    // 是否启用批次管理
    isBatchMgt   bool
    // 启用标识
    useYn   bool
    // 拓展属性
    extendFields   string
    // 条形码，多条码请用”;”分隔；条码更新只技术增量更新，已有条码无法修改，只能在原条码基础上增加新的条码。例：原商品条码为：BAR001，要增加一条新条码BAR002时，此字段内容应填写为：BAR001;BAR002
    barCode   string
    // 商品属性编码
    attribute   string
    // 商品类型:NORMAL-普通商品、 COMBINE-组合商品、 DISTRIBUTION-分销商品
    type   string
    // 是否区域销售属性
    isAreaSale   bool
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
func (r *TaobaoWlbWmsSkuUpdateRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetItemId() string {
    return r.itemId
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsSkuUpdateRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetStoreCode() string {
    return r.storeCode
}
// Name Setter
// 商品名称
func (r *TaobaoWlbWmsSkuUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetName() string {
    return r.name
}
// Title Setter
// 商品标题
func (r *TaobaoWlbWmsSkuUpdateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetTitle() string {
    return r.title
}
// Category Setter
// 商品类别编码（外部系统类别）
func (r *TaobaoWlbWmsSkuUpdateRequest) SetCategory(category string) error {
    r.category = category
    r.Set("category", category)
    return nil
}

// Category Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetCategory() string {
    return r.category
}
// CategoryName Setter
// 商品类别名称
func (r *TaobaoWlbWmsSkuUpdateRequest) SetCategoryName(categoryName string) error {
    r.categoryName = categoryName
    r.Set("category_name", categoryName)
    return nil
}

// CategoryName Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetCategoryName() string {
    return r.categoryName
}
// Brand Setter
// 品牌编码
func (r *TaobaoWlbWmsSkuUpdateRequest) SetBrand(brand string) error {
    r.brand = brand
    r.Set("brand", brand)
    return nil
}

// Brand Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetBrand() string {
    return r.brand
}
// BrandName Setter
// 品牌名称
func (r *TaobaoWlbWmsSkuUpdateRequest) SetBrandName(brandName string) error {
    r.brandName = brandName
    r.Set("brand_name", brandName)
    return nil
}

// BrandName Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetBrandName() string {
    return r.brandName
}
// Specification Setter
// 规格
func (r *TaobaoWlbWmsSkuUpdateRequest) SetSpecification(specification string) error {
    r.specification = specification
    r.Set("specification", specification)
    return nil
}

// Specification Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetSpecification() string {
    return r.specification
}
// Color Setter
// 颜色
func (r *TaobaoWlbWmsSkuUpdateRequest) SetColor(color string) error {
    r.color = color
    r.Set("color", color)
    return nil
}

// Color Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetColor() string {
    return r.color
}
// Size Setter
// 尺码
func (r *TaobaoWlbWmsSkuUpdateRequest) SetSize(size string) error {
    r.size = size
    r.Set("size", size)
    return nil
}

// Size Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetSize() string {
    return r.size
}
// GrossWeight Setter
// 毛重，单位克
func (r *TaobaoWlbWmsSkuUpdateRequest) SetGrossWeight(grossWeight int64) error {
    r.grossWeight = grossWeight
    r.Set("gross_weight", grossWeight)
    return nil
}

// GrossWeight Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetGrossWeight() int64 {
    return r.grossWeight
}
// NetWeight Setter
// 净重，单位克
func (r *TaobaoWlbWmsSkuUpdateRequest) SetNetWeight(netWeight int64) error {
    r.netWeight = netWeight
    r.Set("net_weight", netWeight)
    return nil
}

// NetWeight Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetNetWeight() int64 {
    return r.netWeight
}
// Length Setter
// 长度，单位毫米
func (r *TaobaoWlbWmsSkuUpdateRequest) SetLength(length int64) error {
    r.length = length
    r.Set("length", length)
    return nil
}

// Length Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetLength() int64 {
    return r.length
}
// Width Setter
// 宽度，单位毫米
func (r *TaobaoWlbWmsSkuUpdateRequest) SetWidth(width int64) error {
    r.width = width
    r.Set("width", width)
    return nil
}

// Width Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetWidth() int64 {
    return r.width
}
// Height Setter
// 高度，单位毫米
func (r *TaobaoWlbWmsSkuUpdateRequest) SetHeight(height int64) error {
    r.height = height
    r.Set("height", height)
    return nil
}

// Height Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetHeight() int64 {
    return r.height
}
// Volume Setter
// 体积，单位立方厘米
func (r *TaobaoWlbWmsSkuUpdateRequest) SetVolume(volume int64) error {
    r.volume = volume
    r.Set("volume", volume)
    return nil
}

// Volume Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetVolume() int64 {
    return r.volume
}
// Pcs Setter
// 箱规
func (r *TaobaoWlbWmsSkuUpdateRequest) SetPcs(pcs int64) error {
    r.pcs = pcs
    r.Set("pcs", pcs)
    return nil
}

// Pcs Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetPcs() int64 {
    return r.pcs
}
// OriginAddress Setter
// 产地
func (r *TaobaoWlbWmsSkuUpdateRequest) SetOriginAddress(originAddress int64) error {
    r.originAddress = originAddress
    r.Set("origin_address", originAddress)
    return nil
}

// OriginAddress Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetOriginAddress() int64 {
    return r.originAddress
}
// ApprovalNumber Setter
// 批准文号
func (r *TaobaoWlbWmsSkuUpdateRequest) SetApprovalNumber(approvalNumber string) error {
    r.approvalNumber = approvalNumber
    r.Set("approval_number", approvalNumber)
    return nil
}

// ApprovalNumber Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetApprovalNumber() string {
    return r.approvalNumber
}
// IsShelflife Setter
// 是否启用保质期管理
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsShelflife(isShelflife bool) error {
    r.isShelflife = isShelflife
    r.Set("is_shelflife", isShelflife)
    return nil
}

// IsShelflife Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsShelflife() bool {
    return r.isShelflife
}
// Lifecycle Setter
// 商品保质期天数
func (r *TaobaoWlbWmsSkuUpdateRequest) SetLifecycle(lifecycle int64) error {
    r.lifecycle = lifecycle
    r.Set("lifecycle", lifecycle)
    return nil
}

// Lifecycle Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetLifecycle() int64 {
    return r.lifecycle
}
// RejectLifecycle Setter
// 保质期禁收天数
func (r *TaobaoWlbWmsSkuUpdateRequest) SetRejectLifecycle(rejectLifecycle int64) error {
    r.rejectLifecycle = rejectLifecycle
    r.Set("reject_lifecycle", rejectLifecycle)
    return nil
}

// RejectLifecycle Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetRejectLifecycle() int64 {
    return r.rejectLifecycle
}
// LockupLifecycle Setter
// 保质期禁售天数
func (r *TaobaoWlbWmsSkuUpdateRequest) SetLockupLifecycle(lockupLifecycle int64) error {
    r.lockupLifecycle = lockupLifecycle
    r.Set("lockup_lifecycle", lockupLifecycle)
    return nil
}

// LockupLifecycle Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetLockupLifecycle() int64 {
    return r.lockupLifecycle
}
// AdventLifecycle Setter
// 保质期预警天数
func (r *TaobaoWlbWmsSkuUpdateRequest) SetAdventLifecycle(adventLifecycle int64) error {
    r.adventLifecycle = adventLifecycle
    r.Set("advent_lifecycle", adventLifecycle)
    return nil
}

// AdventLifecycle Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetAdventLifecycle() int64 {
    return r.adventLifecycle
}
// IsSnMgt Setter
// 是否启用序列号管理
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsSnMgt(isSnMgt bool) error {
    r.isSnMgt = isSnMgt
    r.Set("is_sn_mgt", isSnMgt)
    return nil
}

// IsSnMgt Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsSnMgt() bool {
    return r.isSnMgt
}
// IsHygroscopic Setter
// 是否易碎品
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsHygroscopic(isHygroscopic bool) error {
    r.isHygroscopic = isHygroscopic
    r.Set("is_hygroscopic", isHygroscopic)
    return nil
}

// IsHygroscopic Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsHygroscopic() bool {
    return r.isHygroscopic
}
// IsDanger Setter
// 是否危险品
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsDanger(isDanger bool) error {
    r.isDanger = isDanger
    r.Set("is_danger", isDanger)
    return nil
}

// IsDanger Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsDanger() bool {
    return r.isDanger
}
// TagPrice Setter
// 吊牌价，单位分
func (r *TaobaoWlbWmsSkuUpdateRequest) SetTagPrice(tagPrice int64) error {
    r.tagPrice = tagPrice
    r.Set("tag_price", tagPrice)
    return nil
}

// TagPrice Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetTagPrice() int64 {
    return r.tagPrice
}
// ItemPrice Setter
// 零售价，单位分
func (r *TaobaoWlbWmsSkuUpdateRequest) SetItemPrice(itemPrice int64) error {
    r.itemPrice = itemPrice
    r.Set("item_price", itemPrice)
    return nil
}

// ItemPrice Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetItemPrice() int64 {
    return r.itemPrice
}
// CostPrice Setter
// 成本价，单位分
func (r *TaobaoWlbWmsSkuUpdateRequest) SetCostPrice(costPrice int64) error {
    r.costPrice = costPrice
    r.Set("cost_price", costPrice)
    return nil
}

// CostPrice Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetCostPrice() int64 {
    return r.costPrice
}
// IsBatchMgt Setter
// 是否启用批次管理
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsBatchMgt(isBatchMgt bool) error {
    r.isBatchMgt = isBatchMgt
    r.Set("is_batch_mgt", isBatchMgt)
    return nil
}

// IsBatchMgt Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsBatchMgt() bool {
    return r.isBatchMgt
}
// UseYn Setter
// 启用标识
func (r *TaobaoWlbWmsSkuUpdateRequest) SetUseYn(useYn bool) error {
    r.useYn = useYn
    r.Set("use_yn", useYn)
    return nil
}

// UseYn Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetUseYn() bool {
    return r.useYn
}
// ExtendFields Setter
// 拓展属性
func (r *TaobaoWlbWmsSkuUpdateRequest) SetExtendFields(extendFields string) error {
    r.extendFields = extendFields
    r.Set("extend_fields", extendFields)
    return nil
}

// ExtendFields Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetExtendFields() string {
    return r.extendFields
}
// BarCode Setter
// 条形码，多条码请用”;”分隔；条码更新只技术增量更新，已有条码无法修改，只能在原条码基础上增加新的条码。例：原商品条码为：BAR001，要增加一条新条码BAR002时，此字段内容应填写为：BAR001;BAR002
func (r *TaobaoWlbWmsSkuUpdateRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

// BarCode Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetBarCode() string {
    return r.barCode
}
// Attribute Setter
// 商品属性编码
func (r *TaobaoWlbWmsSkuUpdateRequest) SetAttribute(attribute string) error {
    r.attribute = attribute
    r.Set("attribute", attribute)
    return nil
}

// Attribute Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetAttribute() string {
    return r.attribute
}
// Type Setter
// 商品类型:NORMAL-普通商品、 COMBINE-组合商品、 DISTRIBUTION-分销商品
func (r *TaobaoWlbWmsSkuUpdateRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetType() string {
    return r.type
}
// IsAreaSale Setter
// 是否区域销售属性
func (r *TaobaoWlbWmsSkuUpdateRequest) SetIsAreaSale(isAreaSale bool) error {
    r.isAreaSale = isAreaSale
    r.Set("is_area_sale", isAreaSale)
    return nil
}

// IsAreaSale Getter
func (r TaobaoWlbWmsSkuUpdateRequest) GetIsAreaSale() bool {
    return r.isAreaSale
}
