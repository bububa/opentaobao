package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductAddAPIRequest 添加产品 API请求
// taobao.fenxiao.product.add
//
// 添加分销平台产品数据。业务逻辑与分销系统前台页面一致。<br/><br/>    * 产品图片默认为空<br/>    * 产品发布后默认为下架状态
type TaobaoFenxiaoProductAddAPIRequest struct {
	model.Params
	// 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。
	_postageType string
	// 是否有发票，可选值：false（否）、true（是），默认false。
	_haveInvoice string
	// 是否有保修，可选值：false（否）、true（是），默认false。
	_haveQuarantee string
	// 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
	_tradeType string
	// 产品名称，长度不超过60个字节。
	_name string
	// 采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_standardPrice string
	// 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_costPrice string
	// 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_retailPriceLow string
	// 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
	_retailPriceHigh string
	// 商家编码，长度不能超过60个字节。
	_outerId string
	// 产品描述，长度为5到25000字符。
	_desc string
	// 所在地：省，例：“浙江”
	_prov string
	// 所在地：市，例：“杭州”
	_city string
	// 平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
	_postageOrdinary string
	// 快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
	_postageFast string
	// ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。
	_postageEms string
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
	// 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_dealerCostPrice string
	// sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
	_skuDealerCostPrices string
	// 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_standardRetailPrice string
	// 产品线ID
	_productcatId int64
	// 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
	_categoryId int64
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

// NewTaobaoFenxiaoProductAddRequest 初始化TaobaoFenxiaoProductAddAPIRequest对象
func NewTaobaoFenxiaoProductAddRequest() *TaobaoFenxiaoProductAddAPIRequest {
	return &TaobaoFenxiaoProductAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductAddAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPostageType is PostageType Setter
// 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetPostageType(_postageType string) error {
	r._postageType = _postageType
	r.Set("postage_type", _postageType)
	return nil
}

// GetPostageType PostageType Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetPostageType() string {
	return r._postageType
}

// SetHaveInvoice is HaveInvoice Setter
// 是否有发票，可选值：false（否）、true（是），默认false。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetHaveInvoice(_haveInvoice string) error {
	r._haveInvoice = _haveInvoice
	r.Set("have_invoice", _haveInvoice)
	return nil
}

// GetHaveInvoice HaveInvoice Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetHaveInvoice() string {
	return r._haveInvoice
}

// SetHaveQuarantee is HaveQuarantee Setter
// 是否有保修，可选值：false（否）、true（是），默认false。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetHaveQuarantee(_haveQuarantee string) error {
	r._haveQuarantee = _haveQuarantee
	r.Set("have_quarantee", _haveQuarantee)
	return nil
}

// GetHaveQuarantee HaveQuarantee Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetHaveQuarantee() string {
	return r._haveQuarantee
}

// SetTradeType is TradeType Setter
// 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
func (r *TaobaoFenxiaoProductAddAPIRequest) SetTradeType(_tradeType string) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetTradeType() string {
	return r._tradeType
}

// SetName is Name Setter
// 产品名称，长度不超过60个字节。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetName() string {
	return r._name
}

// SetStandardPrice is StandardPrice Setter
// 采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetStandardPrice(_standardPrice string) error {
	r._standardPrice = _standardPrice
	r.Set("standard_price", _standardPrice)
	return nil
}

// GetStandardPrice StandardPrice Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetStandardPrice() string {
	return r._standardPrice
}

// SetCostPrice is CostPrice Setter
// 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetCostPrice(_costPrice string) error {
	r._costPrice = _costPrice
	r.Set("cost_price", _costPrice)
	return nil
}

// GetCostPrice CostPrice Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetCostPrice() string {
	return r._costPrice
}

// SetRetailPriceLow is RetailPriceLow Setter
// 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetRetailPriceLow(_retailPriceLow string) error {
	r._retailPriceLow = _retailPriceLow
	r.Set("retail_price_low", _retailPriceLow)
	return nil
}

// GetRetailPriceLow RetailPriceLow Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetRetailPriceLow() string {
	return r._retailPriceLow
}

// SetRetailPriceHigh is RetailPriceHigh Setter
// 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetRetailPriceHigh(_retailPriceHigh string) error {
	r._retailPriceHigh = _retailPriceHigh
	r.Set("retail_price_high", _retailPriceHigh)
	return nil
}

// GetRetailPriceHigh RetailPriceHigh Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetRetailPriceHigh() string {
	return r._retailPriceHigh
}

// SetOuterId is OuterId Setter
// 商家编码，长度不能超过60个字节。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetDesc is Desc Setter
// 产品描述，长度为5到25000字符。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetDesc() string {
	return r._desc
}

// SetProv is Prov Setter
// 所在地：省，例：“浙江”
func (r *TaobaoFenxiaoProductAddAPIRequest) SetProv(_prov string) error {
	r._prov = _prov
	r.Set("prov", _prov)
	return nil
}

// GetProv Prov Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetProv() string {
	return r._prov
}

// SetCity is City Setter
// 所在地：市，例：“杭州”
func (r *TaobaoFenxiaoProductAddAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetCity() string {
	return r._city
}

// SetPostageOrdinary is PostageOrdinary Setter
// 平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetPostageOrdinary(_postageOrdinary string) error {
	r._postageOrdinary = _postageOrdinary
	r.Set("postage_ordinary", _postageOrdinary)
	return nil
}

// GetPostageOrdinary PostageOrdinary Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetPostageOrdinary() string {
	return r._postageOrdinary
}

// SetPostageFast is PostageFast Setter
// 快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetPostageFast(_postageFast string) error {
	r._postageFast = _postageFast
	r.Set("postage_fast", _postageFast)
	return nil
}

// GetPostageFast PostageFast Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetPostageFast() string {
	return r._postageFast
}

// SetPostageEms is PostageEms Setter
// ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetPostageEms(_postageEms string) error {
	r._postageEms = _postageEms
	r.Set("postage_ems", _postageEms)
	return nil
}

// GetPostageEms PostageEms Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetPostageEms() string {
	return r._postageEms
}

// SetIsAuthz is IsAuthz Setter
// 添加产品时，添加入参isAuthz:yes|no <br/>yes:需要授权 <br/>no:不需要授权 <br/>默认是需要授权
func (r *TaobaoFenxiaoProductAddAPIRequest) SetIsAuthz(_isAuthz string) error {
	r._isAuthz = _isAuthz
	r.Set("is_authz", _isAuthz)
	return nil
}

// GetIsAuthz IsAuthz Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetIsAuthz() string {
	return r._isAuthz
}

// SetPicPath is PicPath Setter
// 产品主图图片空间相对路径或绝对路径
func (r *TaobaoFenxiaoProductAddAPIRequest) SetPicPath(_picPath string) error {
	r._picPath = _picPath
	r.Set("pic_path", _picPath)
	return nil
}

// GetPicPath PicPath Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetPicPath() string {
	return r._picPath
}

// SetProperties is Properties Setter
// 产品属性，格式为pid:vid;pid:vid
func (r *TaobaoFenxiaoProductAddAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetProperties() string {
	return r._properties
}

// SetPropertyAlias is PropertyAlias Setter
// 属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名）
func (r *TaobaoFenxiaoProductAddAPIRequest) SetPropertyAlias(_propertyAlias string) error {
	r._propertyAlias = _propertyAlias
	r.Set("property_alias", _propertyAlias)
	return nil
}

// GetPropertyAlias PropertyAlias Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetPropertyAlias() string {
	return r._propertyAlias
}

// SetInputProperties is InputProperties Setter
// 自定义属性。格式为pid:value;pid:value
func (r *TaobaoFenxiaoProductAddAPIRequest) SetInputProperties(_inputProperties string) error {
	r._inputProperties = _inputProperties
	r.Set("input_properties", _inputProperties)
	return nil
}

// GetInputProperties InputProperties Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetInputProperties() string {
	return r._inputProperties
}

// SetSkuStandardPrices is SkuStandardPrices Setter
// sku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddAPIRequest) SetSkuStandardPrices(_skuStandardPrices string) error {
	r._skuStandardPrices = _skuStandardPrices
	r.Set("sku_standard_prices", _skuStandardPrices)
	return nil
}

// GetSkuStandardPrices SkuStandardPrices Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetSkuStandardPrices() string {
	return r._skuStandardPrices
}

// SetSkuCostPrices is SkuCostPrices Setter
// sku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddAPIRequest) SetSkuCostPrices(_skuCostPrices string) error {
	r._skuCostPrices = _skuCostPrices
	r.Set("sku_cost_prices", _skuCostPrices)
	return nil
}

// GetSkuCostPrices SkuCostPrices Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetSkuCostPrices() string {
	return r._skuCostPrices
}

// SetSkuOuterIds is SkuOuterIds Setter
// sku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddAPIRequest) SetSkuOuterIds(_skuOuterIds string) error {
	r._skuOuterIds = _skuOuterIds
	r.Set("sku_outer_ids", _skuOuterIds)
	return nil
}

// GetSkuOuterIds SkuOuterIds Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetSkuOuterIds() string {
	return r._skuOuterIds
}

// SetSkuQuantitys is SkuQuantitys Setter
// sku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddAPIRequest) SetSkuQuantitys(_skuQuantitys string) error {
	r._skuQuantitys = _skuQuantitys
	r.Set("sku_quantitys", _skuQuantitys)
	return nil
}

// GetSkuQuantitys SkuQuantitys Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetSkuQuantitys() string {
	return r._skuQuantitys
}

// SetSkuProperties is SkuProperties Setter
// sku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddAPIRequest) SetSkuProperties(_skuProperties string) error {
	r._skuProperties = _skuProperties
	r.Set("sku_properties", _skuProperties)
	return nil
}

// GetSkuProperties SkuProperties Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetSkuProperties() string {
	return r._skuProperties
}

// SetDealerCostPrice is DealerCostPrice Setter
// 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetDealerCostPrice(_dealerCostPrice string) error {
	r._dealerCostPrice = _dealerCostPrice
	r.Set("dealer_cost_price", _dealerCostPrice)
	return nil
}

// GetDealerCostPrice DealerCostPrice Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetDealerCostPrice() string {
	return r._dealerCostPrice
}

// SetSkuDealerCostPrices is SkuDealerCostPrices Setter
// sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetSkuDealerCostPrices(_skuDealerCostPrices string) error {
	r._skuDealerCostPrices = _skuDealerCostPrices
	r.Set("sku_dealer_cost_prices", _skuDealerCostPrices)
	return nil
}

// GetSkuDealerCostPrices SkuDealerCostPrices Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetSkuDealerCostPrices() string {
	return r._skuDealerCostPrices
}

// SetStandardRetailPrice is StandardRetailPrice Setter
// 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetStandardRetailPrice(_standardRetailPrice string) error {
	r._standardRetailPrice = _standardRetailPrice
	r.Set("standard_retail_price", _standardRetailPrice)
	return nil
}

// GetStandardRetailPrice StandardRetailPrice Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetStandardRetailPrice() string {
	return r._standardRetailPrice
}

// SetProductcatId is ProductcatId Setter
// 产品线ID
func (r *TaobaoFenxiaoProductAddAPIRequest) SetProductcatId(_productcatId int64) error {
	r._productcatId = _productcatId
	r.Set("productcat_id", _productcatId)
	return nil
}

// GetProductcatId ProductcatId Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetProductcatId() int64 {
	return r._productcatId
}

// SetCategoryId is CategoryId Setter
// 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetQuantity is Quantity Setter
// 产品库存必须是1到999999。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetQuantity() int64 {
	return r._quantity
}

// SetPostageId is PostageId Setter
// 运费模板ID，参考taobao.postages.get。
func (r *TaobaoFenxiaoProductAddAPIRequest) SetPostageId(_postageId int64) error {
	r._postageId = _postageId
	r.Set("postage_id", _postageId)
	return nil
}

// GetPostageId PostageId Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetPostageId() int64 {
	return r._postageId
}

// SetDiscountId is DiscountId Setter
// 折扣ID
func (r *TaobaoFenxiaoProductAddAPIRequest) SetDiscountId(_discountId int64) error {
	r._discountId = _discountId
	r.Set("discount_id", _discountId)
	return nil
}

// GetDiscountId DiscountId Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetDiscountId() int64 {
	return r._discountId
}

// SetImage is Image Setter
// 产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片
func (r *TaobaoFenxiaoProductAddAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetImage() *model.File {
	return r._image
}

// SetItemId is ItemId Setter
// 导入的商品ID
func (r *TaobaoFenxiaoProductAddAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSpuId is SpuId Setter
// 产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传
func (r *TaobaoFenxiaoProductAddAPIRequest) SetSpuId(_spuId int64) error {
	r._spuId = _spuId
	r.Set("spu_id", _spuId)
	return nil
}

// GetSpuId SpuId Getter
func (r TaobaoFenxiaoProductAddAPIRequest) GetSpuId() int64 {
	return r._spuId
}
