package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouscesliteminfoputAPIRequest 电子价签显示用商品信息写入 API请求
// taobao.uscesl.iteminfo.put
//
// 用于电子价签上显示的商品信息的写入，包含价格及促销信息
type TaobaouscesliteminfoputAPIRequest struct {
	model.Params
	// 型号
	_modelNum string
	// 价格单位
	_priceUnit string
	// 品牌名
	_brandName string
	// 销售规格
	_saleSpec string
	// 分类
	_categoryName string
	// 等级
	_rank string
	// 商品变更状态
	_itemChangeStatus string
	// 实际销售价格，单位：分
	_acctionPrice string
	// 能效
	_energyEfficiency string
	// 商品skuId
	_skuId string
	// 促销开始时间:yyyy-MM-dd HH:mm:ss
	_promotionStart string
	// 商品条码
	_itemBarCode string
	// 商品名称
	_itemTitle string
	// 促销文案
	_promotionText string
	// 扩展属性C
	_customizeFeatureC string
	// 扩展属性D
	_customizeFeatureD string
	// 扩展属性E
	_customizeFeatureE string
	// 扩展属性F
	_customizeFeatureF string
	// 扩展属性G
	_customizeFeatureG string
	// 扩展属性H
	_customizeFeatureH string
	// 扩展属性I
	_customizeFeatureI string
	// 扩展属性J
	_customizeFeatureJ string
	// 二维码
	_itemQrCode string
	// 促销结束时间:yyyy-MM-dd HH:mm:ss
	_promotionEnd string
	// 促销原因
	_promotionReason string
	// 商品原价
	_originalPrice string
	// 商品简称
	_shortTitle string
	// 扩展属性B
	_customizeFeatureB string
	// 产地
	_productionPlace string
	// 扩展属性A
	_customizeFeatureA string
	// 商品状态0:在售 1:售罄
	_itemStatus int64
	// 商品编码
	_itemId int64
	// 门店ID
	_shopId int64
	// 促销状态0:非促销 1:促销
	_ifPromotion bool
}

// NewTaobaouscesliteminfoputRequest 初始化TaobaouscesliteminfoputAPIRequest对象
func NewTaobaouscesliteminfoputRequest() *TaobaouscesliteminfoputAPIRequest {
	return &TaobaouscesliteminfoputAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouscesliteminfoputAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.iteminfo.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouscesliteminfoputAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouscesliteminfoputAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModelNum is ModelNum Setter
// 型号
func (r *TaobaouscesliteminfoputAPIRequest) SetModelNum(_modelNum string) error {
	r._modelNum = _modelNum
	r.Set("model_num", _modelNum)
	return nil
}

// GetModelNum ModelNum Getter
func (r TaobaouscesliteminfoputAPIRequest) GetModelNum() string {
	return r._modelNum
}

// SetPriceUnit is PriceUnit Setter
// 价格单位
func (r *TaobaouscesliteminfoputAPIRequest) SetPriceUnit(_priceUnit string) error {
	r._priceUnit = _priceUnit
	r.Set("price_unit", _priceUnit)
	return nil
}

// GetPriceUnit PriceUnit Getter
func (r TaobaouscesliteminfoputAPIRequest) GetPriceUnit() string {
	return r._priceUnit
}

// SetBrandName is BrandName Setter
// 品牌名
func (r *TaobaouscesliteminfoputAPIRequest) SetBrandName(_brandName string) error {
	r._brandName = _brandName
	r.Set("brand_name", _brandName)
	return nil
}

// GetBrandName BrandName Getter
func (r TaobaouscesliteminfoputAPIRequest) GetBrandName() string {
	return r._brandName
}

// SetSaleSpec is SaleSpec Setter
// 销售规格
func (r *TaobaouscesliteminfoputAPIRequest) SetSaleSpec(_saleSpec string) error {
	r._saleSpec = _saleSpec
	r.Set("sale_spec", _saleSpec)
	return nil
}

// GetSaleSpec SaleSpec Getter
func (r TaobaouscesliteminfoputAPIRequest) GetSaleSpec() string {
	return r._saleSpec
}

// SetCategoryName is CategoryName Setter
// 分类
func (r *TaobaouscesliteminfoputAPIRequest) SetCategoryName(_categoryName string) error {
	r._categoryName = _categoryName
	r.Set("category_name", _categoryName)
	return nil
}

// GetCategoryName CategoryName Getter
func (r TaobaouscesliteminfoputAPIRequest) GetCategoryName() string {
	return r._categoryName
}

// SetRank is Rank Setter
// 等级
func (r *TaobaouscesliteminfoputAPIRequest) SetRank(_rank string) error {
	r._rank = _rank
	r.Set("rank", _rank)
	return nil
}

// GetRank Rank Getter
func (r TaobaouscesliteminfoputAPIRequest) GetRank() string {
	return r._rank
}

// SetItemChangeStatus is ItemChangeStatus Setter
// 商品变更状态
func (r *TaobaouscesliteminfoputAPIRequest) SetItemChangeStatus(_itemChangeStatus string) error {
	r._itemChangeStatus = _itemChangeStatus
	r.Set("item_change_status", _itemChangeStatus)
	return nil
}

// GetItemChangeStatus ItemChangeStatus Getter
func (r TaobaouscesliteminfoputAPIRequest) GetItemChangeStatus() string {
	return r._itemChangeStatus
}

// SetAcctionPrice is AcctionPrice Setter
// 实际销售价格，单位：分
func (r *TaobaouscesliteminfoputAPIRequest) SetAcctionPrice(_acctionPrice string) error {
	r._acctionPrice = _acctionPrice
	r.Set("acction_price", _acctionPrice)
	return nil
}

// GetAcctionPrice AcctionPrice Getter
func (r TaobaouscesliteminfoputAPIRequest) GetAcctionPrice() string {
	return r._acctionPrice
}

// SetEnergyEfficiency is EnergyEfficiency Setter
// 能效
func (r *TaobaouscesliteminfoputAPIRequest) SetEnergyEfficiency(_energyEfficiency string) error {
	r._energyEfficiency = _energyEfficiency
	r.Set("energy_efficiency", _energyEfficiency)
	return nil
}

// GetEnergyEfficiency EnergyEfficiency Getter
func (r TaobaouscesliteminfoputAPIRequest) GetEnergyEfficiency() string {
	return r._energyEfficiency
}

// SetSkuId is SkuId Setter
// 商品skuId
func (r *TaobaouscesliteminfoputAPIRequest) SetSkuId(_skuId string) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaouscesliteminfoputAPIRequest) GetSkuId() string {
	return r._skuId
}

// SetPromotionStart is PromotionStart Setter
// 促销开始时间:yyyy-MM-dd HH:mm:ss
func (r *TaobaouscesliteminfoputAPIRequest) SetPromotionStart(_promotionStart string) error {
	r._promotionStart = _promotionStart
	r.Set("promotion_start", _promotionStart)
	return nil
}

// GetPromotionStart PromotionStart Getter
func (r TaobaouscesliteminfoputAPIRequest) GetPromotionStart() string {
	return r._promotionStart
}

// SetItemBarCode is ItemBarCode Setter
// 商品条码
func (r *TaobaouscesliteminfoputAPIRequest) SetItemBarCode(_itemBarCode string) error {
	r._itemBarCode = _itemBarCode
	r.Set("item_bar_code", _itemBarCode)
	return nil
}

// GetItemBarCode ItemBarCode Getter
func (r TaobaouscesliteminfoputAPIRequest) GetItemBarCode() string {
	return r._itemBarCode
}

// SetItemTitle is ItemTitle Setter
// 商品名称
func (r *TaobaouscesliteminfoputAPIRequest) SetItemTitle(_itemTitle string) error {
	r._itemTitle = _itemTitle
	r.Set("item_title", _itemTitle)
	return nil
}

// GetItemTitle ItemTitle Getter
func (r TaobaouscesliteminfoputAPIRequest) GetItemTitle() string {
	return r._itemTitle
}

// SetPromotionText is PromotionText Setter
// 促销文案
func (r *TaobaouscesliteminfoputAPIRequest) SetPromotionText(_promotionText string) error {
	r._promotionText = _promotionText
	r.Set("promotion_text", _promotionText)
	return nil
}

// GetPromotionText PromotionText Getter
func (r TaobaouscesliteminfoputAPIRequest) GetPromotionText() string {
	return r._promotionText
}

// SetCustomizeFeatureC is CustomizeFeatureC Setter
// 扩展属性C
func (r *TaobaouscesliteminfoputAPIRequest) SetCustomizeFeatureC(_customizeFeatureC string) error {
	r._customizeFeatureC = _customizeFeatureC
	r.Set("customize_feature_c", _customizeFeatureC)
	return nil
}

// GetCustomizeFeatureC CustomizeFeatureC Getter
func (r TaobaouscesliteminfoputAPIRequest) GetCustomizeFeatureC() string {
	return r._customizeFeatureC
}

// SetCustomizeFeatureD is CustomizeFeatureD Setter
// 扩展属性D
func (r *TaobaouscesliteminfoputAPIRequest) SetCustomizeFeatureD(_customizeFeatureD string) error {
	r._customizeFeatureD = _customizeFeatureD
	r.Set("customize_feature_d", _customizeFeatureD)
	return nil
}

// GetCustomizeFeatureD CustomizeFeatureD Getter
func (r TaobaouscesliteminfoputAPIRequest) GetCustomizeFeatureD() string {
	return r._customizeFeatureD
}

// SetCustomizeFeatureE is CustomizeFeatureE Setter
// 扩展属性E
func (r *TaobaouscesliteminfoputAPIRequest) SetCustomizeFeatureE(_customizeFeatureE string) error {
	r._customizeFeatureE = _customizeFeatureE
	r.Set("customize_feature_e", _customizeFeatureE)
	return nil
}

// GetCustomizeFeatureE CustomizeFeatureE Getter
func (r TaobaouscesliteminfoputAPIRequest) GetCustomizeFeatureE() string {
	return r._customizeFeatureE
}

// SetCustomizeFeatureF is CustomizeFeatureF Setter
// 扩展属性F
func (r *TaobaouscesliteminfoputAPIRequest) SetCustomizeFeatureF(_customizeFeatureF string) error {
	r._customizeFeatureF = _customizeFeatureF
	r.Set("customize_feature_f", _customizeFeatureF)
	return nil
}

// GetCustomizeFeatureF CustomizeFeatureF Getter
func (r TaobaouscesliteminfoputAPIRequest) GetCustomizeFeatureF() string {
	return r._customizeFeatureF
}

// SetCustomizeFeatureG is CustomizeFeatureG Setter
// 扩展属性G
func (r *TaobaouscesliteminfoputAPIRequest) SetCustomizeFeatureG(_customizeFeatureG string) error {
	r._customizeFeatureG = _customizeFeatureG
	r.Set("customize_feature_g", _customizeFeatureG)
	return nil
}

// GetCustomizeFeatureG CustomizeFeatureG Getter
func (r TaobaouscesliteminfoputAPIRequest) GetCustomizeFeatureG() string {
	return r._customizeFeatureG
}

// SetCustomizeFeatureH is CustomizeFeatureH Setter
// 扩展属性H
func (r *TaobaouscesliteminfoputAPIRequest) SetCustomizeFeatureH(_customizeFeatureH string) error {
	r._customizeFeatureH = _customizeFeatureH
	r.Set("customize_feature_h", _customizeFeatureH)
	return nil
}

// GetCustomizeFeatureH CustomizeFeatureH Getter
func (r TaobaouscesliteminfoputAPIRequest) GetCustomizeFeatureH() string {
	return r._customizeFeatureH
}

// SetCustomizeFeatureI is CustomizeFeatureI Setter
// 扩展属性I
func (r *TaobaouscesliteminfoputAPIRequest) SetCustomizeFeatureI(_customizeFeatureI string) error {
	r._customizeFeatureI = _customizeFeatureI
	r.Set("customize_feature_i", _customizeFeatureI)
	return nil
}

// GetCustomizeFeatureI CustomizeFeatureI Getter
func (r TaobaouscesliteminfoputAPIRequest) GetCustomizeFeatureI() string {
	return r._customizeFeatureI
}

// SetCustomizeFeatureJ is CustomizeFeatureJ Setter
// 扩展属性J
func (r *TaobaouscesliteminfoputAPIRequest) SetCustomizeFeatureJ(_customizeFeatureJ string) error {
	r._customizeFeatureJ = _customizeFeatureJ
	r.Set("customize_feature_j", _customizeFeatureJ)
	return nil
}

// GetCustomizeFeatureJ CustomizeFeatureJ Getter
func (r TaobaouscesliteminfoputAPIRequest) GetCustomizeFeatureJ() string {
	return r._customizeFeatureJ
}

// SetItemQrCode is ItemQrCode Setter
// 二维码
func (r *TaobaouscesliteminfoputAPIRequest) SetItemQrCode(_itemQrCode string) error {
	r._itemQrCode = _itemQrCode
	r.Set("item_qr_code", _itemQrCode)
	return nil
}

// GetItemQrCode ItemQrCode Getter
func (r TaobaouscesliteminfoputAPIRequest) GetItemQrCode() string {
	return r._itemQrCode
}

// SetPromotionEnd is PromotionEnd Setter
// 促销结束时间:yyyy-MM-dd HH:mm:ss
func (r *TaobaouscesliteminfoputAPIRequest) SetPromotionEnd(_promotionEnd string) error {
	r._promotionEnd = _promotionEnd
	r.Set("promotion_end", _promotionEnd)
	return nil
}

// GetPromotionEnd PromotionEnd Getter
func (r TaobaouscesliteminfoputAPIRequest) GetPromotionEnd() string {
	return r._promotionEnd
}

// SetPromotionReason is PromotionReason Setter
// 促销原因
func (r *TaobaouscesliteminfoputAPIRequest) SetPromotionReason(_promotionReason string) error {
	r._promotionReason = _promotionReason
	r.Set("promotion_reason", _promotionReason)
	return nil
}

// GetPromotionReason PromotionReason Getter
func (r TaobaouscesliteminfoputAPIRequest) GetPromotionReason() string {
	return r._promotionReason
}

// SetOriginalPrice is OriginalPrice Setter
// 商品原价
func (r *TaobaouscesliteminfoputAPIRequest) SetOriginalPrice(_originalPrice string) error {
	r._originalPrice = _originalPrice
	r.Set("original_price", _originalPrice)
	return nil
}

// GetOriginalPrice OriginalPrice Getter
func (r TaobaouscesliteminfoputAPIRequest) GetOriginalPrice() string {
	return r._originalPrice
}

// SetShortTitle is ShortTitle Setter
// 商品简称
func (r *TaobaouscesliteminfoputAPIRequest) SetShortTitle(_shortTitle string) error {
	r._shortTitle = _shortTitle
	r.Set("short_title", _shortTitle)
	return nil
}

// GetShortTitle ShortTitle Getter
func (r TaobaouscesliteminfoputAPIRequest) GetShortTitle() string {
	return r._shortTitle
}

// SetCustomizeFeatureB is CustomizeFeatureB Setter
// 扩展属性B
func (r *TaobaouscesliteminfoputAPIRequest) SetCustomizeFeatureB(_customizeFeatureB string) error {
	r._customizeFeatureB = _customizeFeatureB
	r.Set("customize_feature_b", _customizeFeatureB)
	return nil
}

// GetCustomizeFeatureB CustomizeFeatureB Getter
func (r TaobaouscesliteminfoputAPIRequest) GetCustomizeFeatureB() string {
	return r._customizeFeatureB
}

// SetProductionPlace is ProductionPlace Setter
// 产地
func (r *TaobaouscesliteminfoputAPIRequest) SetProductionPlace(_productionPlace string) error {
	r._productionPlace = _productionPlace
	r.Set("production_place", _productionPlace)
	return nil
}

// GetProductionPlace ProductionPlace Getter
func (r TaobaouscesliteminfoputAPIRequest) GetProductionPlace() string {
	return r._productionPlace
}

// SetCustomizeFeatureA is CustomizeFeatureA Setter
// 扩展属性A
func (r *TaobaouscesliteminfoputAPIRequest) SetCustomizeFeatureA(_customizeFeatureA string) error {
	r._customizeFeatureA = _customizeFeatureA
	r.Set("customize_feature_a", _customizeFeatureA)
	return nil
}

// GetCustomizeFeatureA CustomizeFeatureA Getter
func (r TaobaouscesliteminfoputAPIRequest) GetCustomizeFeatureA() string {
	return r._customizeFeatureA
}

// SetItemStatus is ItemStatus Setter
// 商品状态0:在售 1:售罄
func (r *TaobaouscesliteminfoputAPIRequest) SetItemStatus(_itemStatus int64) error {
	r._itemStatus = _itemStatus
	r.Set("item_status", _itemStatus)
	return nil
}

// GetItemStatus ItemStatus Getter
func (r TaobaouscesliteminfoputAPIRequest) GetItemStatus() int64 {
	return r._itemStatus
}

// SetItemId is ItemId Setter
// 商品编码
func (r *TaobaouscesliteminfoputAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaouscesliteminfoputAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetShopId is ShopId Setter
// 门店ID
func (r *TaobaouscesliteminfoputAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TaobaouscesliteminfoputAPIRequest) GetShopId() int64 {
	return r._shopId
}

// SetIfPromotion is IfPromotion Setter
// 促销状态0:非促销 1:促销
func (r *TaobaouscesliteminfoputAPIRequest) SetIfPromotion(_ifPromotion bool) error {
	r._ifPromotion = _ifPromotion
	r.Set("if_promotion", _ifPromotion)
	return nil
}

// GetIfPromotion IfPromotion Getter
func (r TaobaouscesliteminfoputAPIRequest) GetIfPromotion() bool {
	return r._ifPromotion
}
