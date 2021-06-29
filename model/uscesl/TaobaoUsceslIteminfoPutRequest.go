package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子价签显示用商品信息写入 API请求
taobao.uscesl.iteminfo.put

用于电子价签上显示的商品信息的写入，包含价格及促销信息
*/
type TaobaoUsceslIteminfoPutRequest struct {
    model.Params
    // 型号
    modelNum   string
    // 价格单位
    priceUnit   string
    // 品牌名
    brandName   string
    // 销售规格
    saleSpec   string
    // 分类
    categoryName   string
    // 等级
    rank   string
    // 商品变更状态
    itemChangeStatus   string
    // 实际销售价格，单位：分
    acctionPrice   string
    // 能效
    energyEfficiency   string
    // 商品skuId
    skuId   string
    // 促销开始时间:yyyy-MM-dd HH:mm:ss
    promotionStart   string
    // 商品条码
    itemBarCode   string
    // 商品名称
    itemTitle   string
    // 促销文案
    promotionText   string
    // 扩展属性C
    customizeFeatureC   string
    // 扩展属性D
    customizeFeatureD   string
    // 扩展属性E
    customizeFeatureE   string
    // 扩展属性F
    customizeFeatureF   string
    // 扩展属性G
    customizeFeatureG   string
    // 扩展属性H
    customizeFeatureH   string
    // 扩展属性I
    customizeFeatureI   string
    // 扩展属性J
    customizeFeatureJ   string
    // 二维码
    itemQrCode   string
    // 商品状态0:在售 1:售罄
    itemStatus   int64
    // 促销状态0:非促销 1:促销
    ifPromotion   bool
    // 商品编码
    itemId   int64
    // 促销结束时间:yyyy-MM-dd HH:mm:ss
    promotionEnd   string
    // 促销原因
    promotionReason   string
    // 商品原价
    originalPrice   string
    // 商品简称
    shortTitle   string
    // 扩展属性B
    customizeFeatureB   string
    // 产地
    productionPlace   string
    // 扩展属性A
    customizeFeatureA   string
    // 门店ID
    shopId   int64
}

// 初始化TaobaoUsceslIteminfoPutRequest对象
func NewTaobaoUsceslIteminfoPutRequest() *TaobaoUsceslIteminfoPutRequest{
    return &TaobaoUsceslIteminfoPutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslIteminfoPutRequest) GetApiMethodName() string {
    return "taobao.uscesl.iteminfo.put"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslIteminfoPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ModelNum Setter
// 型号
func (r *TaobaoUsceslIteminfoPutRequest) SetModelNum(modelNum string) error {
    r.modelNum = modelNum
    r.Set("model_num", modelNum)
    return nil
}

// ModelNum Getter
func (r TaobaoUsceslIteminfoPutRequest) GetModelNum() string {
    return r.modelNum
}
// PriceUnit Setter
// 价格单位
func (r *TaobaoUsceslIteminfoPutRequest) SetPriceUnit(priceUnit string) error {
    r.priceUnit = priceUnit
    r.Set("price_unit", priceUnit)
    return nil
}

// PriceUnit Getter
func (r TaobaoUsceslIteminfoPutRequest) GetPriceUnit() string {
    return r.priceUnit
}
// BrandName Setter
// 品牌名
func (r *TaobaoUsceslIteminfoPutRequest) SetBrandName(brandName string) error {
    r.brandName = brandName
    r.Set("brand_name", brandName)
    return nil
}

// BrandName Getter
func (r TaobaoUsceslIteminfoPutRequest) GetBrandName() string {
    return r.brandName
}
// SaleSpec Setter
// 销售规格
func (r *TaobaoUsceslIteminfoPutRequest) SetSaleSpec(saleSpec string) error {
    r.saleSpec = saleSpec
    r.Set("sale_spec", saleSpec)
    return nil
}

// SaleSpec Getter
func (r TaobaoUsceslIteminfoPutRequest) GetSaleSpec() string {
    return r.saleSpec
}
// CategoryName Setter
// 分类
func (r *TaobaoUsceslIteminfoPutRequest) SetCategoryName(categoryName string) error {
    r.categoryName = categoryName
    r.Set("category_name", categoryName)
    return nil
}

// CategoryName Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCategoryName() string {
    return r.categoryName
}
// Rank Setter
// 等级
func (r *TaobaoUsceslIteminfoPutRequest) SetRank(rank string) error {
    r.rank = rank
    r.Set("rank", rank)
    return nil
}

// Rank Getter
func (r TaobaoUsceslIteminfoPutRequest) GetRank() string {
    return r.rank
}
// ItemChangeStatus Setter
// 商品变更状态
func (r *TaobaoUsceslIteminfoPutRequest) SetItemChangeStatus(itemChangeStatus string) error {
    r.itemChangeStatus = itemChangeStatus
    r.Set("item_change_status", itemChangeStatus)
    return nil
}

// ItemChangeStatus Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemChangeStatus() string {
    return r.itemChangeStatus
}
// AcctionPrice Setter
// 实际销售价格，单位：分
func (r *TaobaoUsceslIteminfoPutRequest) SetAcctionPrice(acctionPrice string) error {
    r.acctionPrice = acctionPrice
    r.Set("acction_price", acctionPrice)
    return nil
}

// AcctionPrice Getter
func (r TaobaoUsceslIteminfoPutRequest) GetAcctionPrice() string {
    return r.acctionPrice
}
// EnergyEfficiency Setter
// 能效
func (r *TaobaoUsceslIteminfoPutRequest) SetEnergyEfficiency(energyEfficiency string) error {
    r.energyEfficiency = energyEfficiency
    r.Set("energy_efficiency", energyEfficiency)
    return nil
}

// EnergyEfficiency Getter
func (r TaobaoUsceslIteminfoPutRequest) GetEnergyEfficiency() string {
    return r.energyEfficiency
}
// SkuId Setter
// 商品skuId
func (r *TaobaoUsceslIteminfoPutRequest) SetSkuId(skuId string) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

// SkuId Getter
func (r TaobaoUsceslIteminfoPutRequest) GetSkuId() string {
    return r.skuId
}
// PromotionStart Setter
// 促销开始时间:yyyy-MM-dd HH:mm:ss
func (r *TaobaoUsceslIteminfoPutRequest) SetPromotionStart(promotionStart string) error {
    r.promotionStart = promotionStart
    r.Set("promotion_start", promotionStart)
    return nil
}

// PromotionStart Getter
func (r TaobaoUsceslIteminfoPutRequest) GetPromotionStart() string {
    return r.promotionStart
}
// ItemBarCode Setter
// 商品条码
func (r *TaobaoUsceslIteminfoPutRequest) SetItemBarCode(itemBarCode string) error {
    r.itemBarCode = itemBarCode
    r.Set("item_bar_code", itemBarCode)
    return nil
}

// ItemBarCode Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemBarCode() string {
    return r.itemBarCode
}
// ItemTitle Setter
// 商品名称
func (r *TaobaoUsceslIteminfoPutRequest) SetItemTitle(itemTitle string) error {
    r.itemTitle = itemTitle
    r.Set("item_title", itemTitle)
    return nil
}

// ItemTitle Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemTitle() string {
    return r.itemTitle
}
// PromotionText Setter
// 促销文案
func (r *TaobaoUsceslIteminfoPutRequest) SetPromotionText(promotionText string) error {
    r.promotionText = promotionText
    r.Set("promotion_text", promotionText)
    return nil
}

// PromotionText Getter
func (r TaobaoUsceslIteminfoPutRequest) GetPromotionText() string {
    return r.promotionText
}
// CustomizeFeatureC Setter
// 扩展属性C
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureC(customizeFeatureC string) error {
    r.customizeFeatureC = customizeFeatureC
    r.Set("customize_feature_c", customizeFeatureC)
    return nil
}

// CustomizeFeatureC Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureC() string {
    return r.customizeFeatureC
}
// CustomizeFeatureD Setter
// 扩展属性D
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureD(customizeFeatureD string) error {
    r.customizeFeatureD = customizeFeatureD
    r.Set("customize_feature_d", customizeFeatureD)
    return nil
}

// CustomizeFeatureD Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureD() string {
    return r.customizeFeatureD
}
// CustomizeFeatureE Setter
// 扩展属性E
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureE(customizeFeatureE string) error {
    r.customizeFeatureE = customizeFeatureE
    r.Set("customize_feature_e", customizeFeatureE)
    return nil
}

// CustomizeFeatureE Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureE() string {
    return r.customizeFeatureE
}
// CustomizeFeatureF Setter
// 扩展属性F
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureF(customizeFeatureF string) error {
    r.customizeFeatureF = customizeFeatureF
    r.Set("customize_feature_f", customizeFeatureF)
    return nil
}

// CustomizeFeatureF Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureF() string {
    return r.customizeFeatureF
}
// CustomizeFeatureG Setter
// 扩展属性G
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureG(customizeFeatureG string) error {
    r.customizeFeatureG = customizeFeatureG
    r.Set("customize_feature_g", customizeFeatureG)
    return nil
}

// CustomizeFeatureG Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureG() string {
    return r.customizeFeatureG
}
// CustomizeFeatureH Setter
// 扩展属性H
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureH(customizeFeatureH string) error {
    r.customizeFeatureH = customizeFeatureH
    r.Set("customize_feature_h", customizeFeatureH)
    return nil
}

// CustomizeFeatureH Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureH() string {
    return r.customizeFeatureH
}
// CustomizeFeatureI Setter
// 扩展属性I
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureI(customizeFeatureI string) error {
    r.customizeFeatureI = customizeFeatureI
    r.Set("customize_feature_i", customizeFeatureI)
    return nil
}

// CustomizeFeatureI Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureI() string {
    return r.customizeFeatureI
}
// CustomizeFeatureJ Setter
// 扩展属性J
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureJ(customizeFeatureJ string) error {
    r.customizeFeatureJ = customizeFeatureJ
    r.Set("customize_feature_j", customizeFeatureJ)
    return nil
}

// CustomizeFeatureJ Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureJ() string {
    return r.customizeFeatureJ
}
// ItemQrCode Setter
// 二维码
func (r *TaobaoUsceslIteminfoPutRequest) SetItemQrCode(itemQrCode string) error {
    r.itemQrCode = itemQrCode
    r.Set("item_qr_code", itemQrCode)
    return nil
}

// ItemQrCode Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemQrCode() string {
    return r.itemQrCode
}
// ItemStatus Setter
// 商品状态0:在售 1:售罄
func (r *TaobaoUsceslIteminfoPutRequest) SetItemStatus(itemStatus int64) error {
    r.itemStatus = itemStatus
    r.Set("item_status", itemStatus)
    return nil
}

// ItemStatus Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemStatus() int64 {
    return r.itemStatus
}
// IfPromotion Setter
// 促销状态0:非促销 1:促销
func (r *TaobaoUsceslIteminfoPutRequest) SetIfPromotion(ifPromotion bool) error {
    r.ifPromotion = ifPromotion
    r.Set("if_promotion", ifPromotion)
    return nil
}

// IfPromotion Getter
func (r TaobaoUsceslIteminfoPutRequest) GetIfPromotion() bool {
    return r.ifPromotion
}
// ItemId Setter
// 商品编码
func (r *TaobaoUsceslIteminfoPutRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemId() int64 {
    return r.itemId
}
// PromotionEnd Setter
// 促销结束时间:yyyy-MM-dd HH:mm:ss
func (r *TaobaoUsceslIteminfoPutRequest) SetPromotionEnd(promotionEnd string) error {
    r.promotionEnd = promotionEnd
    r.Set("promotion_end", promotionEnd)
    return nil
}

// PromotionEnd Getter
func (r TaobaoUsceslIteminfoPutRequest) GetPromotionEnd() string {
    return r.promotionEnd
}
// PromotionReason Setter
// 促销原因
func (r *TaobaoUsceslIteminfoPutRequest) SetPromotionReason(promotionReason string) error {
    r.promotionReason = promotionReason
    r.Set("promotion_reason", promotionReason)
    return nil
}

// PromotionReason Getter
func (r TaobaoUsceslIteminfoPutRequest) GetPromotionReason() string {
    return r.promotionReason
}
// OriginalPrice Setter
// 商品原价
func (r *TaobaoUsceslIteminfoPutRequest) SetOriginalPrice(originalPrice string) error {
    r.originalPrice = originalPrice
    r.Set("original_price", originalPrice)
    return nil
}

// OriginalPrice Getter
func (r TaobaoUsceslIteminfoPutRequest) GetOriginalPrice() string {
    return r.originalPrice
}
// ShortTitle Setter
// 商品简称
func (r *TaobaoUsceslIteminfoPutRequest) SetShortTitle(shortTitle string) error {
    r.shortTitle = shortTitle
    r.Set("short_title", shortTitle)
    return nil
}

// ShortTitle Getter
func (r TaobaoUsceslIteminfoPutRequest) GetShortTitle() string {
    return r.shortTitle
}
// CustomizeFeatureB Setter
// 扩展属性B
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureB(customizeFeatureB string) error {
    r.customizeFeatureB = customizeFeatureB
    r.Set("customize_feature_b", customizeFeatureB)
    return nil
}

// CustomizeFeatureB Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureB() string {
    return r.customizeFeatureB
}
// ProductionPlace Setter
// 产地
func (r *TaobaoUsceslIteminfoPutRequest) SetProductionPlace(productionPlace string) error {
    r.productionPlace = productionPlace
    r.Set("production_place", productionPlace)
    return nil
}

// ProductionPlace Getter
func (r TaobaoUsceslIteminfoPutRequest) GetProductionPlace() string {
    return r.productionPlace
}
// CustomizeFeatureA Setter
// 扩展属性A
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureA(customizeFeatureA string) error {
    r.customizeFeatureA = customizeFeatureA
    r.Set("customize_feature_a", customizeFeatureA)
    return nil
}

// CustomizeFeatureA Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureA() string {
    return r.customizeFeatureA
}
// ShopId Setter
// 门店ID
func (r *TaobaoUsceslIteminfoPutRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r TaobaoUsceslIteminfoPutRequest) GetShopId() int64 {
    return r.shopId
}
