package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加产品 API请求
taobao.fenxiao.product.add

添加分销平台产品数据。业务逻辑与分销系统前台页面一致。<br/><br/>    * 产品图片默认为空<br/>    * 产品发布后默认为下架状态
*/
type TaobaoFenxiaoProductAddRequest struct {
    model.Params
    // 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。
    _postageType   string
    // 是否有发票，可选值：false（否）、true（是），默认false。
    _haveInvoice   string
    // 是否有保修，可选值：false（否）、true（是），默认false。
    _haveQuarantee   string
    // 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
    _tradeType   string
    // 产品名称，长度不超过60个字节。
    _name   string
    // 产品线ID
    _productcatId   int64
    // 采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    _standardPrice   string
    // 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    _costPrice   string
    // 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    _retailPriceLow   string
    // 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
    _retailPriceHigh   string
    // 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
    _categoryId   int64
    // 商家编码，长度不能超过60个字节。
    _outerId   string
    // 产品库存必须是1到999999。
    _quantity   int64
    // 产品描述，长度为5到25000字符。
    _desc   string
    // 所在地：省，例：“浙江”
    _prov   string
    // 所在地：市，例：“杭州”
    _city   string
    // 运费模板ID，参考taobao.postages.get。
    _postageId   int64
    // 平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
    _postageOrdinary   string
    // 快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
    _postageFast   string
    // ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。
    _postageEms   string
    // 折扣ID
    _discountId   int64
    // 添加产品时，添加入参isAuthz:yes|no <br/>yes:需要授权 <br/>no:不需要授权 <br/>默认是需要授权
    _isAuthz   string
    // 产品主图图片空间相对路径或绝对路径
    _picPath   string
    // 产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片
    _image   []*model.File
    // 产品属性，格式为pid:vid;pid:vid
    _properties   string
    // 属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名）
    _propertyAlias   string
    // 自定义属性。格式为pid:value;pid:value
    _inputProperties   string
    // sku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    _skuStandardPrices   string
    // sku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    _skuCostPrices   string
    // sku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    _skuOuterIds   string
    // sku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    _skuQuantitys   string
    // sku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    _skuProperties   string
    // 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    _dealerCostPrice   string
    // sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
    _skuDealerCostPrices   string
    // 导入的商品ID
    _itemId   int64
    // 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    _standardRetailPrice   string
    // 产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传
    _spuId   int64
}

// 初始化TaobaoFenxiaoProductAddRequest对象
func NewTaobaoFenxiaoProductAddRequest() *TaobaoFenxiaoProductAddRequest{
    return &TaobaoFenxiaoProductAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductAddRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PostageType Setter
// 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。
func (r *TaobaoFenxiaoProductAddRequest) SetPostageType(_postageType string) error {
    r._postageType = _postageType
    r.Set("postage_type", _postageType)
    return nil
}

// PostageType Getter
func (r TaobaoFenxiaoProductAddRequest) GetPostageType() string {
    return r._postageType
}
// HaveInvoice Setter
// 是否有发票，可选值：false（否）、true（是），默认false。
func (r *TaobaoFenxiaoProductAddRequest) SetHaveInvoice(_haveInvoice string) error {
    r._haveInvoice = _haveInvoice
    r.Set("have_invoice", _haveInvoice)
    return nil
}

// HaveInvoice Getter
func (r TaobaoFenxiaoProductAddRequest) GetHaveInvoice() string {
    return r._haveInvoice
}
// HaveQuarantee Setter
// 是否有保修，可选值：false（否）、true（是），默认false。
func (r *TaobaoFenxiaoProductAddRequest) SetHaveQuarantee(_haveQuarantee string) error {
    r._haveQuarantee = _haveQuarantee
    r.Set("have_quarantee", _haveQuarantee)
    return nil
}

// HaveQuarantee Getter
func (r TaobaoFenxiaoProductAddRequest) GetHaveQuarantee() string {
    return r._haveQuarantee
}
// TradeType Setter
// 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
func (r *TaobaoFenxiaoProductAddRequest) SetTradeType(_tradeType string) error {
    r._tradeType = _tradeType
    r.Set("trade_type", _tradeType)
    return nil
}

// TradeType Getter
func (r TaobaoFenxiaoProductAddRequest) GetTradeType() string {
    return r._tradeType
}
// Name Setter
// 产品名称，长度不超过60个字节。
func (r *TaobaoFenxiaoProductAddRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoFenxiaoProductAddRequest) GetName() string {
    return r._name
}
// ProductcatId Setter
// 产品线ID
func (r *TaobaoFenxiaoProductAddRequest) SetProductcatId(_productcatId int64) error {
    r._productcatId = _productcatId
    r.Set("productcat_id", _productcatId)
    return nil
}

// ProductcatId Getter
func (r TaobaoFenxiaoProductAddRequest) GetProductcatId() int64 {
    return r._productcatId
}
// StandardPrice Setter
// 采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetStandardPrice(_standardPrice string) error {
    r._standardPrice = _standardPrice
    r.Set("standard_price", _standardPrice)
    return nil
}

// StandardPrice Getter
func (r TaobaoFenxiaoProductAddRequest) GetStandardPrice() string {
    return r._standardPrice
}
// CostPrice Setter
// 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetCostPrice(_costPrice string) error {
    r._costPrice = _costPrice
    r.Set("cost_price", _costPrice)
    return nil
}

// CostPrice Getter
func (r TaobaoFenxiaoProductAddRequest) GetCostPrice() string {
    return r._costPrice
}
// RetailPriceLow Setter
// 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetRetailPriceLow(_retailPriceLow string) error {
    r._retailPriceLow = _retailPriceLow
    r.Set("retail_price_low", _retailPriceLow)
    return nil
}

// RetailPriceLow Getter
func (r TaobaoFenxiaoProductAddRequest) GetRetailPriceLow() string {
    return r._retailPriceLow
}
// RetailPriceHigh Setter
// 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
func (r *TaobaoFenxiaoProductAddRequest) SetRetailPriceHigh(_retailPriceHigh string) error {
    r._retailPriceHigh = _retailPriceHigh
    r.Set("retail_price_high", _retailPriceHigh)
    return nil
}

// RetailPriceHigh Getter
func (r TaobaoFenxiaoProductAddRequest) GetRetailPriceHigh() string {
    return r._retailPriceHigh
}
// CategoryId Setter
// 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
func (r *TaobaoFenxiaoProductAddRequest) SetCategoryId(_categoryId int64) error {
    r._categoryId = _categoryId
    r.Set("category_id", _categoryId)
    return nil
}

// CategoryId Getter
func (r TaobaoFenxiaoProductAddRequest) GetCategoryId() int64 {
    return r._categoryId
}
// OuterId Setter
// 商家编码，长度不能超过60个字节。
func (r *TaobaoFenxiaoProductAddRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoFenxiaoProductAddRequest) GetOuterId() string {
    return r._outerId
}
// Quantity Setter
// 产品库存必须是1到999999。
func (r *TaobaoFenxiaoProductAddRequest) SetQuantity(_quantity int64) error {
    r._quantity = _quantity
    r.Set("quantity", _quantity)
    return nil
}

// Quantity Getter
func (r TaobaoFenxiaoProductAddRequest) GetQuantity() int64 {
    return r._quantity
}
// Desc Setter
// 产品描述，长度为5到25000字符。
func (r *TaobaoFenxiaoProductAddRequest) SetDesc(_desc string) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r TaobaoFenxiaoProductAddRequest) GetDesc() string {
    return r._desc
}
// Prov Setter
// 所在地：省，例：“浙江”
func (r *TaobaoFenxiaoProductAddRequest) SetProv(_prov string) error {
    r._prov = _prov
    r.Set("prov", _prov)
    return nil
}

// Prov Getter
func (r TaobaoFenxiaoProductAddRequest) GetProv() string {
    return r._prov
}
// City Setter
// 所在地：市，例：“杭州”
func (r *TaobaoFenxiaoProductAddRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r TaobaoFenxiaoProductAddRequest) GetCity() string {
    return r._city
}
// PostageId Setter
// 运费模板ID，参考taobao.postages.get。
func (r *TaobaoFenxiaoProductAddRequest) SetPostageId(_postageId int64) error {
    r._postageId = _postageId
    r.Set("postage_id", _postageId)
    return nil
}

// PostageId Getter
func (r TaobaoFenxiaoProductAddRequest) GetPostageId() int64 {
    return r._postageId
}
// PostageOrdinary Setter
// 平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetPostageOrdinary(_postageOrdinary string) error {
    r._postageOrdinary = _postageOrdinary
    r.Set("postage_ordinary", _postageOrdinary)
    return nil
}

// PostageOrdinary Getter
func (r TaobaoFenxiaoProductAddRequest) GetPostageOrdinary() string {
    return r._postageOrdinary
}
// PostageFast Setter
// 快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetPostageFast(_postageFast string) error {
    r._postageFast = _postageFast
    r.Set("postage_fast", _postageFast)
    return nil
}

// PostageFast Getter
func (r TaobaoFenxiaoProductAddRequest) GetPostageFast() string {
    return r._postageFast
}
// PostageEms Setter
// ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetPostageEms(_postageEms string) error {
    r._postageEms = _postageEms
    r.Set("postage_ems", _postageEms)
    return nil
}

// PostageEms Getter
func (r TaobaoFenxiaoProductAddRequest) GetPostageEms() string {
    return r._postageEms
}
// DiscountId Setter
// 折扣ID
func (r *TaobaoFenxiaoProductAddRequest) SetDiscountId(_discountId int64) error {
    r._discountId = _discountId
    r.Set("discount_id", _discountId)
    return nil
}

// DiscountId Getter
func (r TaobaoFenxiaoProductAddRequest) GetDiscountId() int64 {
    return r._discountId
}
// IsAuthz Setter
// 添加产品时，添加入参isAuthz:yes|no <br/>yes:需要授权 <br/>no:不需要授权 <br/>默认是需要授权
func (r *TaobaoFenxiaoProductAddRequest) SetIsAuthz(_isAuthz string) error {
    r._isAuthz = _isAuthz
    r.Set("is_authz", _isAuthz)
    return nil
}

// IsAuthz Getter
func (r TaobaoFenxiaoProductAddRequest) GetIsAuthz() string {
    return r._isAuthz
}
// PicPath Setter
// 产品主图图片空间相对路径或绝对路径
func (r *TaobaoFenxiaoProductAddRequest) SetPicPath(_picPath string) error {
    r._picPath = _picPath
    r.Set("pic_path", _picPath)
    return nil
}

// PicPath Getter
func (r TaobaoFenxiaoProductAddRequest) GetPicPath() string {
    return r._picPath
}
// Image Setter
// 产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片
func (r *TaobaoFenxiaoProductAddRequest) SetImage(_image []*model.File) error {
    r._image = _image
    r.Set("image", _image)
    return nil
}

// Image Getter
func (r TaobaoFenxiaoProductAddRequest) GetImage() []*model.File {
    return r._image
}
// Properties Setter
// 产品属性，格式为pid:vid;pid:vid
func (r *TaobaoFenxiaoProductAddRequest) SetProperties(_properties string) error {
    r._properties = _properties
    r.Set("properties", _properties)
    return nil
}

// Properties Getter
func (r TaobaoFenxiaoProductAddRequest) GetProperties() string {
    return r._properties
}
// PropertyAlias Setter
// 属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名）
func (r *TaobaoFenxiaoProductAddRequest) SetPropertyAlias(_propertyAlias string) error {
    r._propertyAlias = _propertyAlias
    r.Set("property_alias", _propertyAlias)
    return nil
}

// PropertyAlias Getter
func (r TaobaoFenxiaoProductAddRequest) GetPropertyAlias() string {
    return r._propertyAlias
}
// InputProperties Setter
// 自定义属性。格式为pid:value;pid:value
func (r *TaobaoFenxiaoProductAddRequest) SetInputProperties(_inputProperties string) error {
    r._inputProperties = _inputProperties
    r.Set("input_properties", _inputProperties)
    return nil
}

// InputProperties Getter
func (r TaobaoFenxiaoProductAddRequest) GetInputProperties() string {
    return r._inputProperties
}
// SkuStandardPrices Setter
// sku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddRequest) SetSkuStandardPrices(_skuStandardPrices string) error {
    r._skuStandardPrices = _skuStandardPrices
    r.Set("sku_standard_prices", _skuStandardPrices)
    return nil
}

// SkuStandardPrices Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuStandardPrices() string {
    return r._skuStandardPrices
}
// SkuCostPrices Setter
// sku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddRequest) SetSkuCostPrices(_skuCostPrices string) error {
    r._skuCostPrices = _skuCostPrices
    r.Set("sku_cost_prices", _skuCostPrices)
    return nil
}

// SkuCostPrices Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuCostPrices() string {
    return r._skuCostPrices
}
// SkuOuterIds Setter
// sku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddRequest) SetSkuOuterIds(_skuOuterIds string) error {
    r._skuOuterIds = _skuOuterIds
    r.Set("sku_outer_ids", _skuOuterIds)
    return nil
}

// SkuOuterIds Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuOuterIds() string {
    return r._skuOuterIds
}
// SkuQuantitys Setter
// sku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddRequest) SetSkuQuantitys(_skuQuantitys string) error {
    r._skuQuantitys = _skuQuantitys
    r.Set("sku_quantitys", _skuQuantitys)
    return nil
}

// SkuQuantitys Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuQuantitys() string {
    return r._skuQuantitys
}
// SkuProperties Setter
// sku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddRequest) SetSkuProperties(_skuProperties string) error {
    r._skuProperties = _skuProperties
    r.Set("sku_properties", _skuProperties)
    return nil
}

// SkuProperties Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuProperties() string {
    return r._skuProperties
}
// DealerCostPrice Setter
// 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetDealerCostPrice(_dealerCostPrice string) error {
    r._dealerCostPrice = _dealerCostPrice
    r.Set("dealer_cost_price", _dealerCostPrice)
    return nil
}

// DealerCostPrice Getter
func (r TaobaoFenxiaoProductAddRequest) GetDealerCostPrice() string {
    return r._dealerCostPrice
}
// SkuDealerCostPrices Setter
// sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetSkuDealerCostPrices(_skuDealerCostPrices string) error {
    r._skuDealerCostPrices = _skuDealerCostPrices
    r.Set("sku_dealer_cost_prices", _skuDealerCostPrices)
    return nil
}

// SkuDealerCostPrices Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuDealerCostPrices() string {
    return r._skuDealerCostPrices
}
// ItemId Setter
// 导入的商品ID
func (r *TaobaoFenxiaoProductAddRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoFenxiaoProductAddRequest) GetItemId() int64 {
    return r._itemId
}
// StandardRetailPrice Setter
// 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetStandardRetailPrice(_standardRetailPrice string) error {
    r._standardRetailPrice = _standardRetailPrice
    r.Set("standard_retail_price", _standardRetailPrice)
    return nil
}

// StandardRetailPrice Getter
func (r TaobaoFenxiaoProductAddRequest) GetStandardRetailPrice() string {
    return r._standardRetailPrice
}
// SpuId Setter
// 产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传
func (r *TaobaoFenxiaoProductAddRequest) SetSpuId(_spuId int64) error {
    r._spuId = _spuId
    r.Set("spu_id", _spuId)
    return nil
}

// SpuId Getter
func (r TaobaoFenxiaoProductAddRequest) GetSpuId() int64 {
    return r._spuId
}
