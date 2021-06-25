package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品同步 APIRequest
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

func NewTaobaoWlbWmsSkuCreateRequest() *TaobaoWlbWmsSkuCreateRequest{
    return &TaobaoWlbWmsSkuCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsSkuCreateRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.sku.create"
}

func (r TaobaoWlbWmsSkuCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsSkuCreateRequest) SetItemCode(itemCode string) error {
    r.itemCode = itemCode
    r.Set("item_code", itemCode)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetItemCode() string {
    return r.itemCode
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetBarCode() string {
    return r.barCode
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetName() string {
    return r.name
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetType() string {
    return r.type
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetCategory(category string) error {
    r.category = category
    r.Set("category", category)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetCategory() string {
    return r.category
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetCategoryName(categoryName string) error {
    r.categoryName = categoryName
    r.Set("category_name", categoryName)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetCategoryName() string {
    return r.categoryName
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetBrand(brand string) error {
    r.brand = brand
    r.Set("brand", brand)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetBrand() string {
    return r.brand
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetBrandName(brandName string) error {
    r.brandName = brandName
    r.Set("brand_name", brandName)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetBrandName() string {
    return r.brandName
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetSpecification(specification string) error {
    r.specification = specification
    r.Set("specification", specification)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetSpecification() string {
    return r.specification
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetColor(color string) error {
    r.color = color
    r.Set("color", color)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetColor() string {
    return r.color
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetSize(size string) error {
    r.size = size
    r.Set("size", size)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetSize() string {
    return r.size
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetGrossWeight(grossWeight int64) error {
    r.grossWeight = grossWeight
    r.Set("gross_weight", grossWeight)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetGrossWeight() int64 {
    return r.grossWeight
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetNetWeight(netWeight int64) error {
    r.netWeight = netWeight
    r.Set("net_weight", netWeight)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetNetWeight() int64 {
    return r.netWeight
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetLength(length int64) error {
    r.length = length
    r.Set("length", length)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetLength() int64 {
    return r.length
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetWidth(width int64) error {
    r.width = width
    r.Set("width", width)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetWidth() int64 {
    return r.width
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetHeight(height int64) error {
    r.height = height
    r.Set("height", height)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetHeight() int64 {
    return r.height
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetVolume(volume int64) error {
    r.volume = volume
    r.Set("volume", volume)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetVolume() int64 {
    return r.volume
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetPcs(pcs int64) error {
    r.pcs = pcs
    r.Set("pcs", pcs)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetPcs() int64 {
    return r.pcs
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetOriginAddress(originAddress int64) error {
    r.originAddress = originAddress
    r.Set("origin_address", originAddress)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetOriginAddress() int64 {
    return r.originAddress
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetApprovalNumber(approvalNumber string) error {
    r.approvalNumber = approvalNumber
    r.Set("approval_number", approvalNumber)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetApprovalNumber() string {
    return r.approvalNumber
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetIsShelflife(isShelflife bool) error {
    r.isShelflife = isShelflife
    r.Set("is_shelflife", isShelflife)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetIsShelflife() bool {
    return r.isShelflife
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetLifecycle(lifecycle int64) error {
    r.lifecycle = lifecycle
    r.Set("lifecycle", lifecycle)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetLifecycle() int64 {
    return r.lifecycle
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetRejectLifecycle(rejectLifecycle int64) error {
    r.rejectLifecycle = rejectLifecycle
    r.Set("reject_lifecycle", rejectLifecycle)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetRejectLifecycle() int64 {
    return r.rejectLifecycle
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetLockupLifecycle(lockupLifecycle int64) error {
    r.lockupLifecycle = lockupLifecycle
    r.Set("lockup_lifecycle", lockupLifecycle)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetLockupLifecycle() int64 {
    return r.lockupLifecycle
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetAdventLifecycle(adventLifecycle int64) error {
    r.adventLifecycle = adventLifecycle
    r.Set("advent_lifecycle", adventLifecycle)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetAdventLifecycle() int64 {
    return r.adventLifecycle
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetIsSnMgt(isSnMgt bool) error {
    r.isSnMgt = isSnMgt
    r.Set("is_sn_mgt", isSnMgt)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetIsSnMgt() bool {
    return r.isSnMgt
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetIsHygroscopic(isHygroscopic bool) error {
    r.isHygroscopic = isHygroscopic
    r.Set("is_hygroscopic", isHygroscopic)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetIsHygroscopic() bool {
    return r.isHygroscopic
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetIsDanger(isDanger bool) error {
    r.isDanger = isDanger
    r.Set("is_danger", isDanger)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetIsDanger() bool {
    return r.isDanger
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetTagPrice(tagPrice int64) error {
    r.tagPrice = tagPrice
    r.Set("tag_price", tagPrice)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetTagPrice() int64 {
    return r.tagPrice
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetItemPrice(itemPrice int64) error {
    r.itemPrice = itemPrice
    r.Set("item_price", itemPrice)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetItemPrice() int64 {
    return r.itemPrice
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetCostPrice(costPrice int64) error {
    r.costPrice = costPrice
    r.Set("cost_price", costPrice)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetCostPrice() int64 {
    return r.costPrice
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetIsBatchMgt(isBatchMgt bool) error {
    r.isBatchMgt = isBatchMgt
    r.Set("is_batch_mgt", isBatchMgt)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetIsBatchMgt() bool {
    return r.isBatchMgt
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetUseYn(useYn bool) error {
    r.useYn = useYn
    r.Set("use_yn", useYn)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetUseYn() bool {
    return r.useYn
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetExtendFields(extendFields string) error {
    r.extendFields = extendFields
    r.Set("extend_fields", extendFields)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetExtendFields() string {
    return r.extendFields
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetItemId() string {
    return r.itemId
}

func (r *TaobaoWlbWmsSkuCreateRequest) SetIsAreaSale(isAreaSale bool) error {
    r.isAreaSale = isAreaSale
    r.Set("is_area_sale", isAreaSale)
    return nil
}

func (r TaobaoWlbWmsSkuCreateRequest) GetIsAreaSale() bool {
    return r.isAreaSale
}

