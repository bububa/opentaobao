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
    itemCode   string
    // 条形码，多条码请用”；”分隔；
    barCode   string
    // 仓库编码
    storeCode   string
    // 商品名称
    name   string
    // 商品标题
    title   string
    // 商品类别NORMAL：普通商品、COMBINE：组合商品、DISTRIBUTION：分销商品、HAOCAI耗材、FUSHUPIN附属品、BAOCAI 包材、XUNI虚拟商品、QITA其他)
    type   string
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
    // 商家商品ID
    itemId   string
    // 是否区域销售
    isAreaSale   bool
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
func (r *TaobaoWlbWmsSkuCreateRequest) SetItemCode(itemCode string) error {
    r.itemCode = itemCode
    r.Set("item_code", itemCode)
    return nil
}

// ItemCode Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetItemCode() string {
    return r.itemCode
}
// BarCode Setter
// 条形码，多条码请用”；”分隔；
func (r *TaobaoWlbWmsSkuCreateRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

// BarCode Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetBarCode() string {
    return r.barCode
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsSkuCreateRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetStoreCode() string {
    return r.storeCode
}
// Name Setter
// 商品名称
func (r *TaobaoWlbWmsSkuCreateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetName() string {
    return r.name
}
// Title Setter
// 商品标题
func (r *TaobaoWlbWmsSkuCreateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetTitle() string {
    return r.title
}
// Type Setter
// 商品类别NORMAL：普通商品、COMBINE：组合商品、DISTRIBUTION：分销商品、HAOCAI耗材、FUSHUPIN附属品、BAOCAI 包材、XUNI虚拟商品、QITA其他)
func (r *TaobaoWlbWmsSkuCreateRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetType() string {
    return r.type
}
// Category Setter
// 商品类别编码（外部系统类别）
func (r *TaobaoWlbWmsSkuCreateRequest) SetCategory(category string) error {
    r.category = category
    r.Set("category", category)
    return nil
}

// Category Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetCategory() string {
    return r.category
}
// CategoryName Setter
// 商品类别名称
func (r *TaobaoWlbWmsSkuCreateRequest) SetCategoryName(categoryName string) error {
    r.categoryName = categoryName
    r.Set("category_name", categoryName)
    return nil
}

// CategoryName Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetCategoryName() string {
    return r.categoryName
}
// Brand Setter
// 品牌编码
func (r *TaobaoWlbWmsSkuCreateRequest) SetBrand(brand string) error {
    r.brand = brand
    r.Set("brand", brand)
    return nil
}

// Brand Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetBrand() string {
    return r.brand
}
// BrandName Setter
// 品牌名称
func (r *TaobaoWlbWmsSkuCreateRequest) SetBrandName(brandName string) error {
    r.brandName = brandName
    r.Set("brand_name", brandName)
    return nil
}

// BrandName Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetBrandName() string {
    return r.brandName
}
// Specification Setter
// 规格
func (r *TaobaoWlbWmsSkuCreateRequest) SetSpecification(specification string) error {
    r.specification = specification
    r.Set("specification", specification)
    return nil
}

// Specification Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetSpecification() string {
    return r.specification
}
// Color Setter
// 颜色
func (r *TaobaoWlbWmsSkuCreateRequest) SetColor(color string) error {
    r.color = color
    r.Set("color", color)
    return nil
}

// Color Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetColor() string {
    return r.color
}
// Size Setter
// 尺码
func (r *TaobaoWlbWmsSkuCreateRequest) SetSize(size string) error {
    r.size = size
    r.Set("size", size)
    return nil
}

// Size Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetSize() string {
    return r.size
}
// GrossWeight Setter
// 毛重，单位克
func (r *TaobaoWlbWmsSkuCreateRequest) SetGrossWeight(grossWeight int64) error {
    r.grossWeight = grossWeight
    r.Set("gross_weight", grossWeight)
    return nil
}

// GrossWeight Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetGrossWeight() int64 {
    return r.grossWeight
}
// NetWeight Setter
// 净重，单位克
func (r *TaobaoWlbWmsSkuCreateRequest) SetNetWeight(netWeight int64) error {
    r.netWeight = netWeight
    r.Set("net_weight", netWeight)
    return nil
}

// NetWeight Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetNetWeight() int64 {
    return r.netWeight
}
// Length Setter
// 长度，单位毫米
func (r *TaobaoWlbWmsSkuCreateRequest) SetLength(length int64) error {
    r.length = length
    r.Set("length", length)
    return nil
}

// Length Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetLength() int64 {
    return r.length
}
// Width Setter
// 宽度，单位毫米
func (r *TaobaoWlbWmsSkuCreateRequest) SetWidth(width int64) error {
    r.width = width
    r.Set("width", width)
    return nil
}

// Width Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetWidth() int64 {
    return r.width
}
// Height Setter
// 高度，单位毫米
func (r *TaobaoWlbWmsSkuCreateRequest) SetHeight(height int64) error {
    r.height = height
    r.Set("height", height)
    return nil
}

// Height Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetHeight() int64 {
    return r.height
}
// Volume Setter
// 体积，单位立方厘米
func (r *TaobaoWlbWmsSkuCreateRequest) SetVolume(volume int64) error {
    r.volume = volume
    r.Set("volume", volume)
    return nil
}

// Volume Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetVolume() int64 {
    return r.volume
}
// Pcs Setter
// 箱规
func (r *TaobaoWlbWmsSkuCreateRequest) SetPcs(pcs int64) error {
    r.pcs = pcs
    r.Set("pcs", pcs)
    return nil
}

// Pcs Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetPcs() int64 {
    return r.pcs
}
// OriginAddress Setter
// 产地
func (r *TaobaoWlbWmsSkuCreateRequest) SetOriginAddress(originAddress int64) error {
    r.originAddress = originAddress
    r.Set("origin_address", originAddress)
    return nil
}

// OriginAddress Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetOriginAddress() int64 {
    return r.originAddress
}
// ApprovalNumber Setter
// 批准文号
func (r *TaobaoWlbWmsSkuCreateRequest) SetApprovalNumber(approvalNumber string) error {
    r.approvalNumber = approvalNumber
    r.Set("approval_number", approvalNumber)
    return nil
}

// ApprovalNumber Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetApprovalNumber() string {
    return r.approvalNumber
}
// IsShelflife Setter
// 是否启用保质期管理
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsShelflife(isShelflife bool) error {
    r.isShelflife = isShelflife
    r.Set("is_shelflife", isShelflife)
    return nil
}

// IsShelflife Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsShelflife() bool {
    return r.isShelflife
}
// Lifecycle Setter
// 商品保质期天数
func (r *TaobaoWlbWmsSkuCreateRequest) SetLifecycle(lifecycle int64) error {
    r.lifecycle = lifecycle
    r.Set("lifecycle", lifecycle)
    return nil
}

// Lifecycle Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetLifecycle() int64 {
    return r.lifecycle
}
// RejectLifecycle Setter
// 保质期禁收天数
func (r *TaobaoWlbWmsSkuCreateRequest) SetRejectLifecycle(rejectLifecycle int64) error {
    r.rejectLifecycle = rejectLifecycle
    r.Set("reject_lifecycle", rejectLifecycle)
    return nil
}

// RejectLifecycle Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetRejectLifecycle() int64 {
    return r.rejectLifecycle
}
// LockupLifecycle Setter
// 保质期禁售天数
func (r *TaobaoWlbWmsSkuCreateRequest) SetLockupLifecycle(lockupLifecycle int64) error {
    r.lockupLifecycle = lockupLifecycle
    r.Set("lockup_lifecycle", lockupLifecycle)
    return nil
}

// LockupLifecycle Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetLockupLifecycle() int64 {
    return r.lockupLifecycle
}
// AdventLifecycle Setter
// 保质期预警天数
func (r *TaobaoWlbWmsSkuCreateRequest) SetAdventLifecycle(adventLifecycle int64) error {
    r.adventLifecycle = adventLifecycle
    r.Set("advent_lifecycle", adventLifecycle)
    return nil
}

// AdventLifecycle Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetAdventLifecycle() int64 {
    return r.adventLifecycle
}
// IsSnMgt Setter
// 是否启用序列号管理
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsSnMgt(isSnMgt bool) error {
    r.isSnMgt = isSnMgt
    r.Set("is_sn_mgt", isSnMgt)
    return nil
}

// IsSnMgt Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsSnMgt() bool {
    return r.isSnMgt
}
// IsHygroscopic Setter
// 是否易碎品
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsHygroscopic(isHygroscopic bool) error {
    r.isHygroscopic = isHygroscopic
    r.Set("is_hygroscopic", isHygroscopic)
    return nil
}

// IsHygroscopic Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsHygroscopic() bool {
    return r.isHygroscopic
}
// IsDanger Setter
// 是否危险品
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsDanger(isDanger bool) error {
    r.isDanger = isDanger
    r.Set("is_danger", isDanger)
    return nil
}

// IsDanger Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsDanger() bool {
    return r.isDanger
}
// TagPrice Setter
// 吊牌价，单位分
func (r *TaobaoWlbWmsSkuCreateRequest) SetTagPrice(tagPrice int64) error {
    r.tagPrice = tagPrice
    r.Set("tag_price", tagPrice)
    return nil
}

// TagPrice Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetTagPrice() int64 {
    return r.tagPrice
}
// ItemPrice Setter
// 零售价，单位分
func (r *TaobaoWlbWmsSkuCreateRequest) SetItemPrice(itemPrice int64) error {
    r.itemPrice = itemPrice
    r.Set("item_price", itemPrice)
    return nil
}

// ItemPrice Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetItemPrice() int64 {
    return r.itemPrice
}
// CostPrice Setter
// 成本价，单位分
func (r *TaobaoWlbWmsSkuCreateRequest) SetCostPrice(costPrice int64) error {
    r.costPrice = costPrice
    r.Set("cost_price", costPrice)
    return nil
}

// CostPrice Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetCostPrice() int64 {
    return r.costPrice
}
// IsBatchMgt Setter
// 是否启用批次管理
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsBatchMgt(isBatchMgt bool) error {
    r.isBatchMgt = isBatchMgt
    r.Set("is_batch_mgt", isBatchMgt)
    return nil
}

// IsBatchMgt Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsBatchMgt() bool {
    return r.isBatchMgt
}
// UseYn Setter
// 启用标识
func (r *TaobaoWlbWmsSkuCreateRequest) SetUseYn(useYn bool) error {
    r.useYn = useYn
    r.Set("use_yn", useYn)
    return nil
}

// UseYn Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetUseYn() bool {
    return r.useYn
}
// ExtendFields Setter
// 拓展属性
func (r *TaobaoWlbWmsSkuCreateRequest) SetExtendFields(extendFields string) error {
    r.extendFields = extendFields
    r.Set("extend_fields", extendFields)
    return nil
}

// ExtendFields Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetExtendFields() string {
    return r.extendFields
}
// ItemId Setter
// 商家商品ID
func (r *TaobaoWlbWmsSkuCreateRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetItemId() string {
    return r.itemId
}
// IsAreaSale Setter
// 是否区域销售
func (r *TaobaoWlbWmsSkuCreateRequest) SetIsAreaSale(isAreaSale bool) error {
    r.isAreaSale = isAreaSale
    r.Set("is_area_sale", isAreaSale)
    return nil
}

// IsAreaSale Getter
func (r TaobaoWlbWmsSkuCreateRequest) GetIsAreaSale() bool {
    return r.isAreaSale
}
