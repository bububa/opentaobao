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
    _modelNum   string
    // 价格单位
    _priceUnit   string
    // 品牌名
    _brandName   string
    // 销售规格
    _saleSpec   string
    // 分类
    _categoryName   string
    // 等级
    _rank   string
    // 商品变更状态
    _itemChangeStatus   string
    // 实际销售价格，单位：分
    _acctionPrice   string
    // 能效
    _energyEfficiency   string
    // 商品skuId
    _skuId   string
    // 促销开始时间:yyyy-MM-dd HH:mm:ss
    _promotionStart   string
    // 商品条码
    _itemBarCode   string
    // 商品名称
    _itemTitle   string
    // 促销文案
    _promotionText   string
    // 扩展属性C
    _customizeFeatureC   string
    // 扩展属性D
    _customizeFeatureD   string
    // 扩展属性E
    _customizeFeatureE   string
    // 扩展属性F
    _customizeFeatureF   string
    // 扩展属性G
    _customizeFeatureG   string
    // 扩展属性H
    _customizeFeatureH   string
    // 扩展属性I
    _customizeFeatureI   string
    // 扩展属性J
    _customizeFeatureJ   string
    // 二维码
    _itemQrCode   string
    // 商品状态0:在售 1:售罄
    _itemStatus   int64
    // 促销状态0:非促销 1:促销
    _ifPromotion   bool
    // 商品编码
    _itemId   int64
    // 促销结束时间:yyyy-MM-dd HH:mm:ss
    _promotionEnd   string
    // 促销原因
    _promotionReason   string
    // 商品原价
    _originalPrice   string
    // 商品简称
    _shortTitle   string
    // 扩展属性B
    _customizeFeatureB   string
    // 产地
    _productionPlace   string
    // 扩展属性A
    _customizeFeatureA   string
    // 门店ID
    _shopId   int64
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
func (r *TaobaoUsceslIteminfoPutRequest) SetModelNum(_modelNum string) error {
    r._modelNum = _modelNum
    r.Set("model_num", _modelNum)
    return nil
}

// ModelNum Getter
func (r TaobaoUsceslIteminfoPutRequest) GetModelNum() string {
    return r._modelNum
}
// PriceUnit Setter
// 价格单位
func (r *TaobaoUsceslIteminfoPutRequest) SetPriceUnit(_priceUnit string) error {
    r._priceUnit = _priceUnit
    r.Set("price_unit", _priceUnit)
    return nil
}

// PriceUnit Getter
func (r TaobaoUsceslIteminfoPutRequest) GetPriceUnit() string {
    return r._priceUnit
}
// BrandName Setter
// 品牌名
func (r *TaobaoUsceslIteminfoPutRequest) SetBrandName(_brandName string) error {
    r._brandName = _brandName
    r.Set("brand_name", _brandName)
    return nil
}

// BrandName Getter
func (r TaobaoUsceslIteminfoPutRequest) GetBrandName() string {
    return r._brandName
}
// SaleSpec Setter
// 销售规格
func (r *TaobaoUsceslIteminfoPutRequest) SetSaleSpec(_saleSpec string) error {
    r._saleSpec = _saleSpec
    r.Set("sale_spec", _saleSpec)
    return nil
}

// SaleSpec Getter
func (r TaobaoUsceslIteminfoPutRequest) GetSaleSpec() string {
    return r._saleSpec
}
// CategoryName Setter
// 分类
func (r *TaobaoUsceslIteminfoPutRequest) SetCategoryName(_categoryName string) error {
    r._categoryName = _categoryName
    r.Set("category_name", _categoryName)
    return nil
}

// CategoryName Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCategoryName() string {
    return r._categoryName
}
// Rank Setter
// 等级
func (r *TaobaoUsceslIteminfoPutRequest) SetRank(_rank string) error {
    r._rank = _rank
    r.Set("rank", _rank)
    return nil
}

// Rank Getter
func (r TaobaoUsceslIteminfoPutRequest) GetRank() string {
    return r._rank
}
// ItemChangeStatus Setter
// 商品变更状态
func (r *TaobaoUsceslIteminfoPutRequest) SetItemChangeStatus(_itemChangeStatus string) error {
    r._itemChangeStatus = _itemChangeStatus
    r.Set("item_change_status", _itemChangeStatus)
    return nil
}

// ItemChangeStatus Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemChangeStatus() string {
    return r._itemChangeStatus
}
// AcctionPrice Setter
// 实际销售价格，单位：分
func (r *TaobaoUsceslIteminfoPutRequest) SetAcctionPrice(_acctionPrice string) error {
    r._acctionPrice = _acctionPrice
    r.Set("acction_price", _acctionPrice)
    return nil
}

// AcctionPrice Getter
func (r TaobaoUsceslIteminfoPutRequest) GetAcctionPrice() string {
    return r._acctionPrice
}
// EnergyEfficiency Setter
// 能效
func (r *TaobaoUsceslIteminfoPutRequest) SetEnergyEfficiency(_energyEfficiency string) error {
    r._energyEfficiency = _energyEfficiency
    r.Set("energy_efficiency", _energyEfficiency)
    return nil
}

// EnergyEfficiency Getter
func (r TaobaoUsceslIteminfoPutRequest) GetEnergyEfficiency() string {
    return r._energyEfficiency
}
// SkuId Setter
// 商品skuId
func (r *TaobaoUsceslIteminfoPutRequest) SetSkuId(_skuId string) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoUsceslIteminfoPutRequest) GetSkuId() string {
    return r._skuId
}
// PromotionStart Setter
// 促销开始时间:yyyy-MM-dd HH:mm:ss
func (r *TaobaoUsceslIteminfoPutRequest) SetPromotionStart(_promotionStart string) error {
    r._promotionStart = _promotionStart
    r.Set("promotion_start", _promotionStart)
    return nil
}

// PromotionStart Getter
func (r TaobaoUsceslIteminfoPutRequest) GetPromotionStart() string {
    return r._promotionStart
}
// ItemBarCode Setter
// 商品条码
func (r *TaobaoUsceslIteminfoPutRequest) SetItemBarCode(_itemBarCode string) error {
    r._itemBarCode = _itemBarCode
    r.Set("item_bar_code", _itemBarCode)
    return nil
}

// ItemBarCode Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemBarCode() string {
    return r._itemBarCode
}
// ItemTitle Setter
// 商品名称
func (r *TaobaoUsceslIteminfoPutRequest) SetItemTitle(_itemTitle string) error {
    r._itemTitle = _itemTitle
    r.Set("item_title", _itemTitle)
    return nil
}

// ItemTitle Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemTitle() string {
    return r._itemTitle
}
// PromotionText Setter
// 促销文案
func (r *TaobaoUsceslIteminfoPutRequest) SetPromotionText(_promotionText string) error {
    r._promotionText = _promotionText
    r.Set("promotion_text", _promotionText)
    return nil
}

// PromotionText Getter
func (r TaobaoUsceslIteminfoPutRequest) GetPromotionText() string {
    return r._promotionText
}
// CustomizeFeatureC Setter
// 扩展属性C
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureC(_customizeFeatureC string) error {
    r._customizeFeatureC = _customizeFeatureC
    r.Set("customize_feature_c", _customizeFeatureC)
    return nil
}

// CustomizeFeatureC Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureC() string {
    return r._customizeFeatureC
}
// CustomizeFeatureD Setter
// 扩展属性D
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureD(_customizeFeatureD string) error {
    r._customizeFeatureD = _customizeFeatureD
    r.Set("customize_feature_d", _customizeFeatureD)
    return nil
}

// CustomizeFeatureD Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureD() string {
    return r._customizeFeatureD
}
// CustomizeFeatureE Setter
// 扩展属性E
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureE(_customizeFeatureE string) error {
    r._customizeFeatureE = _customizeFeatureE
    r.Set("customize_feature_e", _customizeFeatureE)
    return nil
}

// CustomizeFeatureE Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureE() string {
    return r._customizeFeatureE
}
// CustomizeFeatureF Setter
// 扩展属性F
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureF(_customizeFeatureF string) error {
    r._customizeFeatureF = _customizeFeatureF
    r.Set("customize_feature_f", _customizeFeatureF)
    return nil
}

// CustomizeFeatureF Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureF() string {
    return r._customizeFeatureF
}
// CustomizeFeatureG Setter
// 扩展属性G
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureG(_customizeFeatureG string) error {
    r._customizeFeatureG = _customizeFeatureG
    r.Set("customize_feature_g", _customizeFeatureG)
    return nil
}

// CustomizeFeatureG Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureG() string {
    return r._customizeFeatureG
}
// CustomizeFeatureH Setter
// 扩展属性H
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureH(_customizeFeatureH string) error {
    r._customizeFeatureH = _customizeFeatureH
    r.Set("customize_feature_h", _customizeFeatureH)
    return nil
}

// CustomizeFeatureH Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureH() string {
    return r._customizeFeatureH
}
// CustomizeFeatureI Setter
// 扩展属性I
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureI(_customizeFeatureI string) error {
    r._customizeFeatureI = _customizeFeatureI
    r.Set("customize_feature_i", _customizeFeatureI)
    return nil
}

// CustomizeFeatureI Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureI() string {
    return r._customizeFeatureI
}
// CustomizeFeatureJ Setter
// 扩展属性J
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureJ(_customizeFeatureJ string) error {
    r._customizeFeatureJ = _customizeFeatureJ
    r.Set("customize_feature_j", _customizeFeatureJ)
    return nil
}

// CustomizeFeatureJ Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureJ() string {
    return r._customizeFeatureJ
}
// ItemQrCode Setter
// 二维码
func (r *TaobaoUsceslIteminfoPutRequest) SetItemQrCode(_itemQrCode string) error {
    r._itemQrCode = _itemQrCode
    r.Set("item_qr_code", _itemQrCode)
    return nil
}

// ItemQrCode Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemQrCode() string {
    return r._itemQrCode
}
// ItemStatus Setter
// 商品状态0:在售 1:售罄
func (r *TaobaoUsceslIteminfoPutRequest) SetItemStatus(_itemStatus int64) error {
    r._itemStatus = _itemStatus
    r.Set("item_status", _itemStatus)
    return nil
}

// ItemStatus Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemStatus() int64 {
    return r._itemStatus
}
// IfPromotion Setter
// 促销状态0:非促销 1:促销
func (r *TaobaoUsceslIteminfoPutRequest) SetIfPromotion(_ifPromotion bool) error {
    r._ifPromotion = _ifPromotion
    r.Set("if_promotion", _ifPromotion)
    return nil
}

// IfPromotion Getter
func (r TaobaoUsceslIteminfoPutRequest) GetIfPromotion() bool {
    return r._ifPromotion
}
// ItemId Setter
// 商品编码
func (r *TaobaoUsceslIteminfoPutRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoUsceslIteminfoPutRequest) GetItemId() int64 {
    return r._itemId
}
// PromotionEnd Setter
// 促销结束时间:yyyy-MM-dd HH:mm:ss
func (r *TaobaoUsceslIteminfoPutRequest) SetPromotionEnd(_promotionEnd string) error {
    r._promotionEnd = _promotionEnd
    r.Set("promotion_end", _promotionEnd)
    return nil
}

// PromotionEnd Getter
func (r TaobaoUsceslIteminfoPutRequest) GetPromotionEnd() string {
    return r._promotionEnd
}
// PromotionReason Setter
// 促销原因
func (r *TaobaoUsceslIteminfoPutRequest) SetPromotionReason(_promotionReason string) error {
    r._promotionReason = _promotionReason
    r.Set("promotion_reason", _promotionReason)
    return nil
}

// PromotionReason Getter
func (r TaobaoUsceslIteminfoPutRequest) GetPromotionReason() string {
    return r._promotionReason
}
// OriginalPrice Setter
// 商品原价
func (r *TaobaoUsceslIteminfoPutRequest) SetOriginalPrice(_originalPrice string) error {
    r._originalPrice = _originalPrice
    r.Set("original_price", _originalPrice)
    return nil
}

// OriginalPrice Getter
func (r TaobaoUsceslIteminfoPutRequest) GetOriginalPrice() string {
    return r._originalPrice
}
// ShortTitle Setter
// 商品简称
func (r *TaobaoUsceslIteminfoPutRequest) SetShortTitle(_shortTitle string) error {
    r._shortTitle = _shortTitle
    r.Set("short_title", _shortTitle)
    return nil
}

// ShortTitle Getter
func (r TaobaoUsceslIteminfoPutRequest) GetShortTitle() string {
    return r._shortTitle
}
// CustomizeFeatureB Setter
// 扩展属性B
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureB(_customizeFeatureB string) error {
    r._customizeFeatureB = _customizeFeatureB
    r.Set("customize_feature_b", _customizeFeatureB)
    return nil
}

// CustomizeFeatureB Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureB() string {
    return r._customizeFeatureB
}
// ProductionPlace Setter
// 产地
func (r *TaobaoUsceslIteminfoPutRequest) SetProductionPlace(_productionPlace string) error {
    r._productionPlace = _productionPlace
    r.Set("production_place", _productionPlace)
    return nil
}

// ProductionPlace Getter
func (r TaobaoUsceslIteminfoPutRequest) GetProductionPlace() string {
    return r._productionPlace
}
// CustomizeFeatureA Setter
// 扩展属性A
func (r *TaobaoUsceslIteminfoPutRequest) SetCustomizeFeatureA(_customizeFeatureA string) error {
    r._customizeFeatureA = _customizeFeatureA
    r.Set("customize_feature_a", _customizeFeatureA)
    return nil
}

// CustomizeFeatureA Getter
func (r TaobaoUsceslIteminfoPutRequest) GetCustomizeFeatureA() string {
    return r._customizeFeatureA
}
// ShopId Setter
// 门店ID
func (r *TaobaoUsceslIteminfoPutRequest) SetShopId(_shopId int64) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r TaobaoUsceslIteminfoPutRequest) GetShopId() int64 {
    return r._shopId
}
