package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductaddAPIRequest 添加产品 API请求
// taobao.fenxiao.product.add
//
// 添加分销平台产品数据。业务逻辑与分销系统前台页面一致。&lt;br/&gt;&lt;br/&gt;    * 产品图片默认为空&lt;br/&gt;    * 产品发布后默认为下架状态
type TaobaofenxiaoproductaddAPIRequest struct {
	model.Params
	// 产品名称，长度不超过60个字节。
	_name string
	// 采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_standardPrice string
	// 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_standardRetailPrice string
	// 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_retailPriceLow string
	// 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
	_retailPriceHigh string
	// 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_costPrice string
	// 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_dealerCostPrice string
	// 商家编码，长度不能超过60个字节。
	_outerId string
	// 产品描述，长度为5到25000字符。
	_desc string
	// 所在地：省，例：“浙江”
	_prov string
	// 所在地：市，例：“杭州”
	_city string
	// 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。
	_postageType string
	// 平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
	_postageOrdinary string
	// 快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
	_postageFast string
	// ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。
	_postageEms string
	// 是否有发票，可选值：false（否）、true（是），默认false。
	_haveInvoice string
	// 是否有保修，可选值：false（否）、true（是），默认false。
	_haveQuarantee string
	// 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
	_tradeType string
	// 添加产品时，添加入参isAuthz:yes|no <br/>yes:需要授权 <br/>no:不需要授权 <br/>默认是需要授权
	_isAuthz string
	// 产品主图图片空间相对路径或绝对路径
	_picPath string
	// 产品属性，格式为pid:vid;pid:vid
	_properties string
	// 属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名）
	_propertyAlias string
	// 自定义属性。格式为pid:value;pid:value
	_inputProperties string
	// sku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
	_skuStandardPrices string
	// sku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
	_skuCostPrices string
	// sku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
	_skuOuterIds string
	// sku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
	_skuQuantitys string
	// sku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
	_skuProperties string
	// sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
	_skuDealerCostPrices string
	// 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
	_categoryId int64
	// 产品线ID
	_productcatId int64
	// 产品库存必须是1到999999。
	_quantity int64
	// 运费模板ID，参考taobao.postages.get。
	_postageId int64
	// 折扣ID
	_discountId int64
	// 产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片
	_image *model.File
	// 导入的商品ID
	_itemId int64
	// 产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传
	_spuId int64
}

// NewTaobaofenxiaoproductaddRequest 初始化TaobaofenxiaoproductaddAPIRequest对象
func NewTaobaofenxiaoproductaddRequest() *TaobaofenxiaoproductaddAPIRequest {
	return &TaobaofenxiaoproductaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductaddAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 产品名称，长度不超过60个字节。
func (r *TaobaofenxiaoproductaddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetName() string {
	return r._name
}

// SetStandardPrice is StandardPrice Setter
// 采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaofenxiaoproductaddAPIRequest) SetStandardPrice(_standardPrice string) error {
	r._standardPrice = _standardPrice
	r.Set("standard_price", _standardPrice)
	return nil
}

// GetStandardPrice StandardPrice Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetStandardPrice() string {
	return r._standardPrice
}

// SetStandardRetailPrice is StandardRetailPrice Setter
// 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaofenxiaoproductaddAPIRequest) SetStandardRetailPrice(_standardRetailPrice string) error {
	r._standardRetailPrice = _standardRetailPrice
	r.Set("standard_retail_price", _standardRetailPrice)
	return nil
}

// GetStandardRetailPrice StandardRetailPrice Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetStandardRetailPrice() string {
	return r._standardRetailPrice
}

// SetRetailPriceLow is RetailPriceLow Setter
// 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaofenxiaoproductaddAPIRequest) SetRetailPriceLow(_retailPriceLow string) error {
	r._retailPriceLow = _retailPriceLow
	r.Set("retail_price_low", _retailPriceLow)
	return nil
}

// GetRetailPriceLow RetailPriceLow Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetRetailPriceLow() string {
	return r._retailPriceLow
}

// SetRetailPriceHigh is RetailPriceHigh Setter
// 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
func (r *TaobaofenxiaoproductaddAPIRequest) SetRetailPriceHigh(_retailPriceHigh string) error {
	r._retailPriceHigh = _retailPriceHigh
	r.Set("retail_price_high", _retailPriceHigh)
	return nil
}

// GetRetailPriceHigh RetailPriceHigh Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetRetailPriceHigh() string {
	return r._retailPriceHigh
}

// SetCostPrice is CostPrice Setter
// 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaofenxiaoproductaddAPIRequest) SetCostPrice(_costPrice string) error {
	r._costPrice = _costPrice
	r.Set("cost_price", _costPrice)
	return nil
}

// GetCostPrice CostPrice Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetCostPrice() string {
	return r._costPrice
}

// SetDealerCostPrice is DealerCostPrice Setter
// 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaofenxiaoproductaddAPIRequest) SetDealerCostPrice(_dealerCostPrice string) error {
	r._dealerCostPrice = _dealerCostPrice
	r.Set("dealer_cost_price", _dealerCostPrice)
	return nil
}

// GetDealerCostPrice DealerCostPrice Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetDealerCostPrice() string {
	return r._dealerCostPrice
}

// SetOuterId is OuterId Setter
// 商家编码，长度不能超过60个字节。
func (r *TaobaofenxiaoproductaddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetDesc is Desc Setter
// 产品描述，长度为5到25000字符。
func (r *TaobaofenxiaoproductaddAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetDesc() string {
	return r._desc
}

// SetProv is Prov Setter
// 所在地：省，例：“浙江”
func (r *TaobaofenxiaoproductaddAPIRequest) SetProv(_prov string) error {
	r._prov = _prov
	r.Set("prov", _prov)
	return nil
}

// GetProv Prov Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetProv() string {
	return r._prov
}

// SetCity is City Setter
// 所在地：市，例：“杭州”
func (r *TaobaofenxiaoproductaddAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetCity() string {
	return r._city
}

// SetPostageType is PostageType Setter
// 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。
func (r *TaobaofenxiaoproductaddAPIRequest) SetPostageType(_postageType string) error {
	r._postageType = _postageType
	r.Set("postage_type", _postageType)
	return nil
}

// GetPostageType PostageType Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetPostageType() string {
	return r._postageType
}

// SetPostageOrdinary is PostageOrdinary Setter
// 平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
func (r *TaobaofenxiaoproductaddAPIRequest) SetPostageOrdinary(_postageOrdinary string) error {
	r._postageOrdinary = _postageOrdinary
	r.Set("postage_ordinary", _postageOrdinary)
	return nil
}

// GetPostageOrdinary PostageOrdinary Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetPostageOrdinary() string {
	return r._postageOrdinary
}

// SetPostageFast is PostageFast Setter
// 快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
func (r *TaobaofenxiaoproductaddAPIRequest) SetPostageFast(_postageFast string) error {
	r._postageFast = _postageFast
	r.Set("postage_fast", _postageFast)
	return nil
}

// GetPostageFast PostageFast Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetPostageFast() string {
	return r._postageFast
}

// SetPostageEms is PostageEms Setter
// ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。
func (r *TaobaofenxiaoproductaddAPIRequest) SetPostageEms(_postageEms string) error {
	r._postageEms = _postageEms
	r.Set("postage_ems", _postageEms)
	return nil
}

// GetPostageEms PostageEms Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetPostageEms() string {
	return r._postageEms
}

// SetHaveInvoice is HaveInvoice Setter
// 是否有发票，可选值：false（否）、true（是），默认false。
func (r *TaobaofenxiaoproductaddAPIRequest) SetHaveInvoice(_haveInvoice string) error {
	r._haveInvoice = _haveInvoice
	r.Set("have_invoice", _haveInvoice)
	return nil
}

// GetHaveInvoice HaveInvoice Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetHaveInvoice() string {
	return r._haveInvoice
}

// SetHaveQuarantee is HaveQuarantee Setter
// 是否有保修，可选值：false（否）、true（是），默认false。
func (r *TaobaofenxiaoproductaddAPIRequest) SetHaveQuarantee(_haveQuarantee string) error {
	r._haveQuarantee = _haveQuarantee
	r.Set("have_quarantee", _haveQuarantee)
	return nil
}

// GetHaveQuarantee HaveQuarantee Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetHaveQuarantee() string {
	return r._haveQuarantee
}

// SetTradeType is TradeType Setter
// 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
func (r *TaobaofenxiaoproductaddAPIRequest) SetTradeType(_tradeType string) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetTradeType() string {
	return r._tradeType
}

// SetIsAuthz is IsAuthz Setter
// 添加产品时，添加入参isAuthz:yes|no &lt;br/&gt;yes:需要授权 &lt;br/&gt;no:不需要授权 &lt;br/&gt;默认是需要授权
func (r *TaobaofenxiaoproductaddAPIRequest) SetIsAuthz(_isAuthz string) error {
	r._isAuthz = _isAuthz
	r.Set("is_authz", _isAuthz)
	return nil
}

// GetIsAuthz IsAuthz Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetIsAuthz() string {
	return r._isAuthz
}

// SetPicPath is PicPath Setter
// 产品主图图片空间相对路径或绝对路径
func (r *TaobaofenxiaoproductaddAPIRequest) SetPicPath(_picPath string) error {
	r._picPath = _picPath
	r.Set("pic_path", _picPath)
	return nil
}

// GetPicPath PicPath Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetPicPath() string {
	return r._picPath
}

// SetProperties is Properties Setter
// 产品属性，格式为pid:vid;pid:vid
func (r *TaobaofenxiaoproductaddAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetProperties() string {
	return r._properties
}

// SetPropertyAlias is PropertyAlias Setter
// 属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名）
func (r *TaobaofenxiaoproductaddAPIRequest) SetPropertyAlias(_propertyAlias string) error {
	r._propertyAlias = _propertyAlias
	r.Set("property_alias", _propertyAlias)
	return nil
}

// GetPropertyAlias PropertyAlias Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetPropertyAlias() string {
	return r._propertyAlias
}

// SetInputProperties is InputProperties Setter
// 自定义属性。格式为pid:value;pid:value
func (r *TaobaofenxiaoproductaddAPIRequest) SetInputProperties(_inputProperties string) error {
	r._inputProperties = _inputProperties
	r.Set("input_properties", _inputProperties)
	return nil
}

// GetInputProperties InputProperties Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetInputProperties() string {
	return r._inputProperties
}

// SetSkuStandardPrices is SkuStandardPrices Setter
// sku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaofenxiaoproductaddAPIRequest) SetSkuStandardPrices(_skuStandardPrices string) error {
	r._skuStandardPrices = _skuStandardPrices
	r.Set("sku_standard_prices", _skuStandardPrices)
	return nil
}

// GetSkuStandardPrices SkuStandardPrices Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetSkuStandardPrices() string {
	return r._skuStandardPrices
}

// SetSkuCostPrices is SkuCostPrices Setter
// sku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaofenxiaoproductaddAPIRequest) SetSkuCostPrices(_skuCostPrices string) error {
	r._skuCostPrices = _skuCostPrices
	r.Set("sku_cost_prices", _skuCostPrices)
	return nil
}

// GetSkuCostPrices SkuCostPrices Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetSkuCostPrices() string {
	return r._skuCostPrices
}

// SetSkuOuterIds is SkuOuterIds Setter
// sku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaofenxiaoproductaddAPIRequest) SetSkuOuterIds(_skuOuterIds string) error {
	r._skuOuterIds = _skuOuterIds
	r.Set("sku_outer_ids", _skuOuterIds)
	return nil
}

// GetSkuOuterIds SkuOuterIds Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetSkuOuterIds() string {
	return r._skuOuterIds
}

// SetSkuQuantitys is SkuQuantitys Setter
// sku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaofenxiaoproductaddAPIRequest) SetSkuQuantitys(_skuQuantitys string) error {
	r._skuQuantitys = _skuQuantitys
	r.Set("sku_quantitys", _skuQuantitys)
	return nil
}

// GetSkuQuantitys SkuQuantitys Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetSkuQuantitys() string {
	return r._skuQuantitys
}

// SetSkuProperties is SkuProperties Setter
// sku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaofenxiaoproductaddAPIRequest) SetSkuProperties(_skuProperties string) error {
	r._skuProperties = _skuProperties
	r.Set("sku_properties", _skuProperties)
	return nil
}

// GetSkuProperties SkuProperties Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetSkuProperties() string {
	return r._skuProperties
}

// SetSkuDealerCostPrices is SkuDealerCostPrices Setter
// sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
func (r *TaobaofenxiaoproductaddAPIRequest) SetSkuDealerCostPrices(_skuDealerCostPrices string) error {
	r._skuDealerCostPrices = _skuDealerCostPrices
	r.Set("sku_dealer_cost_prices", _skuDealerCostPrices)
	return nil
}

// GetSkuDealerCostPrices SkuDealerCostPrices Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetSkuDealerCostPrices() string {
	return r._skuDealerCostPrices
}

// SetCategoryId is CategoryId Setter
// 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
func (r *TaobaofenxiaoproductaddAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetProductcatId is ProductcatId Setter
// 产品线ID
func (r *TaobaofenxiaoproductaddAPIRequest) SetProductcatId(_productcatId int64) error {
	r._productcatId = _productcatId
	r.Set("productcat_id", _productcatId)
	return nil
}

// GetProductcatId ProductcatId Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetProductcatId() int64 {
	return r._productcatId
}

// SetQuantity is Quantity Setter
// 产品库存必须是1到999999。
func (r *TaobaofenxiaoproductaddAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetQuantity() int64 {
	return r._quantity
}

// SetPostageId is PostageId Setter
// 运费模板ID，参考taobao.postages.get。
func (r *TaobaofenxiaoproductaddAPIRequest) SetPostageId(_postageId int64) error {
	r._postageId = _postageId
	r.Set("postage_id", _postageId)
	return nil
}

// GetPostageId PostageId Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetPostageId() int64 {
	return r._postageId
}

// SetDiscountId is DiscountId Setter
// 折扣ID
func (r *TaobaofenxiaoproductaddAPIRequest) SetDiscountId(_discountId int64) error {
	r._discountId = _discountId
	r.Set("discount_id", _discountId)
	return nil
}

// GetDiscountId DiscountId Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetDiscountId() int64 {
	return r._discountId
}

// SetImage is Image Setter
// 产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片
func (r *TaobaofenxiaoproductaddAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetImage() *model.File {
	return r._image
}

// SetItemId is ItemId Setter
// 导入的商品ID
func (r *TaobaofenxiaoproductaddAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSpuId is SpuId Setter
// 产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传
func (r *TaobaofenxiaoproductaddAPIRequest) SetSpuId(_spuId int64) error {
	r._spuId = _spuId
	r.Set("spu_id", _spuId)
	return nil
}

// GetSpuId SpuId Getter
func (r TaobaofenxiaoproductaddAPIRequest) GetSpuId() int64 {
	return r._spuId
}
