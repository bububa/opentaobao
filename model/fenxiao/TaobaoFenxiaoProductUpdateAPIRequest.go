package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductUpdateAPIRequest 更新产品 API请求
// taobao.fenxiao.product.update
//
// 更新分销平台产品数据，不传更新数据返回失败<br><br/>1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br>
type TaobaoFenxiaoProductUpdateAPIRequest struct {
	model.Params
	// 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）。
	_postageType string
	// 是否有发票，可选值：false（否）、true（是），默认false。
	_haveInvoice string
	// 是否有保修，可选值：false（否）、true（是），默认false。
	_haveQuarantee string
	// 发布状态，可选值：up（上架）、down（下架）、delete（删除），输入非法字符则忽略。
	_status string
	// 产品ID
	_pid int64
	// 产品名称，长度不超过60个字节。
	_name string
	// 采购基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_standardPrice string
	// 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_costPrice string
	// 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_retailPriceLow string
	// 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
	_retailPriceHigh string
	// 商家编码，长度不能超过60个字节。
	_outerId string
	// 产品库存必须是1到999999。
	_quantity int64
	// 产品描述，长度为5到25000字符。
	_desc string
	// 所在地：省，例：“浙江”
	_prov string
	// 所在地：市，例：“杭州”
	_city string
	// 运费模板ID，参考taobao.postages.get。更新时必须指定运费类型为 buyer，否则不更新。
	_postageId int64
	// 平邮费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
	_postageOrdinary string
	// 快递费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
	_postageFast string
	// ems费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
	_postageEms string
	// sku id列表，例：1001,1002,1003。如果传入sku_properties将忽略此参数。
	_skuIds string
	// sku采购价格，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。
	_skuCostPrices string
	// sku库存，单位元，例："10,20,30"，字段必须和sku_ids或sku_properties保持一致。
	_skuQuantitys string
	// sku商家编码 ，单位元，例："S1000,S1002,S1003"，字段必须和上面的id或sku_properties保持一致，如果没有可以写成",,"
	_skuOuterIds string
	// 折扣ID
	_discountId int64
	// sku采购基准价，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。
	_skuStandardPrices string
	// sku属性。格式:pid:vid;pid:vid,表示一组属性如:1627207:3232483;1630696:3284570,表示一组:机身颜色:军绿色;手机套餐:一电一充。多组之间用逗号“,”区分。(属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid)<br/>通过此字段可新增和更新sku。若传入此值将忽略sku_ids字段。sku其他字段与此值保持一致。
	_skuProperties string
	// 根据sku属性删除sku信息。需要按组删除属性。
	_skuPropertiesDel string
	// 产品是否需要授权isAuthz:yes|no <br/>yes:需要授权 <br/>no:不需要授权
	_isAuthz string
	// 产品主图图片空间相对路径或绝对路径
	_picPath string
	// 主图图片，如果pic_path参数不空，则优先使用pic_path，忽略该参数
	_image *model.File
	// 产品属性
	_properties string
	// 属性别名
	_propertyAlias string
	// 自定义属性。格式为pid:value;pid:value
	_inputProperties string
	// 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_dealerCostPrice string
	// sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
	_skuDealerCostPrices string
	// 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
	_categoryId int64
	// 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
	_standardRetailPrice string
}

// NewTaobaoFenxiaoProductUpdateRequest 初始化TaobaoFenxiaoProductUpdateAPIRequest对象
func NewTaobaoFenxiaoProductUpdateRequest() *TaobaoFenxiaoProductUpdateAPIRequest {
	return &TaobaoFenxiaoProductUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPostageType is PostageType Setter
// 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetPostageType(_postageType string) error {
	r._postageType = _postageType
	r.Set("postage_type", _postageType)
	return nil
}

// GetPostageType PostageType Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetPostageType() string {
	return r._postageType
}

// SetHaveInvoice is HaveInvoice Setter
// 是否有发票，可选值：false（否）、true（是），默认false。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetHaveInvoice(_haveInvoice string) error {
	r._haveInvoice = _haveInvoice
	r.Set("have_invoice", _haveInvoice)
	return nil
}

// GetHaveInvoice HaveInvoice Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetHaveInvoice() string {
	return r._haveInvoice
}

// SetHaveQuarantee is HaveQuarantee Setter
// 是否有保修，可选值：false（否）、true（是），默认false。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetHaveQuarantee(_haveQuarantee string) error {
	r._haveQuarantee = _haveQuarantee
	r.Set("have_quarantee", _haveQuarantee)
	return nil
}

// GetHaveQuarantee HaveQuarantee Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetHaveQuarantee() string {
	return r._haveQuarantee
}

// SetStatus is Status Setter
// 发布状态，可选值：up（上架）、down（下架）、delete（删除），输入非法字符则忽略。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetStatus() string {
	return r._status
}

// SetPid is Pid Setter
// 产品ID
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetPid(_pid int64) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetPid() int64 {
	return r._pid
}

// SetName is Name Setter
// 产品名称，长度不超过60个字节。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetName() string {
	return r._name
}

// SetStandardPrice is StandardPrice Setter
// 采购基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetStandardPrice(_standardPrice string) error {
	r._standardPrice = _standardPrice
	r.Set("standard_price", _standardPrice)
	return nil
}

// GetStandardPrice StandardPrice Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetStandardPrice() string {
	return r._standardPrice
}

// SetCostPrice is CostPrice Setter
// 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetCostPrice(_costPrice string) error {
	r._costPrice = _costPrice
	r.Set("cost_price", _costPrice)
	return nil
}

// GetCostPrice CostPrice Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetCostPrice() string {
	return r._costPrice
}

// SetRetailPriceLow is RetailPriceLow Setter
// 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetRetailPriceLow(_retailPriceLow string) error {
	r._retailPriceLow = _retailPriceLow
	r.Set("retail_price_low", _retailPriceLow)
	return nil
}

// GetRetailPriceLow RetailPriceLow Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetRetailPriceLow() string {
	return r._retailPriceLow
}

// SetRetailPriceHigh is RetailPriceHigh Setter
// 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetRetailPriceHigh(_retailPriceHigh string) error {
	r._retailPriceHigh = _retailPriceHigh
	r.Set("retail_price_high", _retailPriceHigh)
	return nil
}

// GetRetailPriceHigh RetailPriceHigh Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetRetailPriceHigh() string {
	return r._retailPriceHigh
}

// SetOuterId is OuterId Setter
// 商家编码，长度不能超过60个字节。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetQuantity is Quantity Setter
// 产品库存必须是1到999999。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}

// SetDesc is Desc Setter
// 产品描述，长度为5到25000字符。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetDesc() string {
	return r._desc
}

// SetProv is Prov Setter
// 所在地：省，例：“浙江”
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetProv(_prov string) error {
	r._prov = _prov
	r.Set("prov", _prov)
	return nil
}

// GetProv Prov Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetProv() string {
	return r._prov
}

// SetCity is City Setter
// 所在地：市，例：“杭州”
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetCity() string {
	return r._city
}

// SetPostageId is PostageId Setter
// 运费模板ID，参考taobao.postages.get。更新时必须指定运费类型为 buyer，否则不更新。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetPostageId(_postageId int64) error {
	r._postageId = _postageId
	r.Set("postage_id", _postageId)
	return nil
}

// GetPostageId PostageId Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetPostageId() int64 {
	return r._postageId
}

// SetPostageOrdinary is PostageOrdinary Setter
// 平邮费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetPostageOrdinary(_postageOrdinary string) error {
	r._postageOrdinary = _postageOrdinary
	r.Set("postage_ordinary", _postageOrdinary)
	return nil
}

// GetPostageOrdinary PostageOrdinary Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetPostageOrdinary() string {
	return r._postageOrdinary
}

// SetPostageFast is PostageFast Setter
// 快递费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetPostageFast(_postageFast string) error {
	r._postageFast = _postageFast
	r.Set("postage_fast", _postageFast)
	return nil
}

// GetPostageFast PostageFast Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetPostageFast() string {
	return r._postageFast
}

// SetPostageEms is PostageEms Setter
// ems费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetPostageEms(_postageEms string) error {
	r._postageEms = _postageEms
	r.Set("postage_ems", _postageEms)
	return nil
}

// GetPostageEms PostageEms Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetPostageEms() string {
	return r._postageEms
}

// SetSkuIds is SkuIds Setter
// sku id列表，例：1001,1002,1003。如果传入sku_properties将忽略此参数。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetSkuIds(_skuIds string) error {
	r._skuIds = _skuIds
	r.Set("sku_ids", _skuIds)
	return nil
}

// GetSkuIds SkuIds Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetSkuIds() string {
	return r._skuIds
}

// SetSkuCostPrices is SkuCostPrices Setter
// sku采购价格，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetSkuCostPrices(_skuCostPrices string) error {
	r._skuCostPrices = _skuCostPrices
	r.Set("sku_cost_prices", _skuCostPrices)
	return nil
}

// GetSkuCostPrices SkuCostPrices Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetSkuCostPrices() string {
	return r._skuCostPrices
}

// SetSkuQuantitys is SkuQuantitys Setter
// sku库存，单位元，例："10,20,30"，字段必须和sku_ids或sku_properties保持一致。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetSkuQuantitys(_skuQuantitys string) error {
	r._skuQuantitys = _skuQuantitys
	r.Set("sku_quantitys", _skuQuantitys)
	return nil
}

// GetSkuQuantitys SkuQuantitys Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetSkuQuantitys() string {
	return r._skuQuantitys
}

// SetSkuOuterIds is SkuOuterIds Setter
// sku商家编码 ，单位元，例："S1000,S1002,S1003"，字段必须和上面的id或sku_properties保持一致，如果没有可以写成",,"
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetSkuOuterIds(_skuOuterIds string) error {
	r._skuOuterIds = _skuOuterIds
	r.Set("sku_outer_ids", _skuOuterIds)
	return nil
}

// GetSkuOuterIds SkuOuterIds Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetSkuOuterIds() string {
	return r._skuOuterIds
}

// SetDiscountId is DiscountId Setter
// 折扣ID
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetDiscountId(_discountId int64) error {
	r._discountId = _discountId
	r.Set("discount_id", _discountId)
	return nil
}

// GetDiscountId DiscountId Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetDiscountId() int64 {
	return r._discountId
}

// SetSkuStandardPrices is SkuStandardPrices Setter
// sku采购基准价，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetSkuStandardPrices(_skuStandardPrices string) error {
	r._skuStandardPrices = _skuStandardPrices
	r.Set("sku_standard_prices", _skuStandardPrices)
	return nil
}

// GetSkuStandardPrices SkuStandardPrices Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetSkuStandardPrices() string {
	return r._skuStandardPrices
}

// SetSkuProperties is SkuProperties Setter
// sku属性。格式:pid:vid;pid:vid,表示一组属性如:1627207:3232483;1630696:3284570,表示一组:机身颜色:军绿色;手机套餐:一电一充。多组之间用逗号“,”区分。(属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid)<br/>通过此字段可新增和更新sku。若传入此值将忽略sku_ids字段。sku其他字段与此值保持一致。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetSkuProperties(_skuProperties string) error {
	r._skuProperties = _skuProperties
	r.Set("sku_properties", _skuProperties)
	return nil
}

// GetSkuProperties SkuProperties Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetSkuProperties() string {
	return r._skuProperties
}

// SetSkuPropertiesDel is SkuPropertiesDel Setter
// 根据sku属性删除sku信息。需要按组删除属性。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetSkuPropertiesDel(_skuPropertiesDel string) error {
	r._skuPropertiesDel = _skuPropertiesDel
	r.Set("sku_properties_del", _skuPropertiesDel)
	return nil
}

// GetSkuPropertiesDel SkuPropertiesDel Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetSkuPropertiesDel() string {
	return r._skuPropertiesDel
}

// SetIsAuthz is IsAuthz Setter
// 产品是否需要授权isAuthz:yes|no <br/>yes:需要授权 <br/>no:不需要授权
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetIsAuthz(_isAuthz string) error {
	r._isAuthz = _isAuthz
	r.Set("is_authz", _isAuthz)
	return nil
}

// GetIsAuthz IsAuthz Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetIsAuthz() string {
	return r._isAuthz
}

// SetPicPath is PicPath Setter
// 产品主图图片空间相对路径或绝对路径
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetPicPath(_picPath string) error {
	r._picPath = _picPath
	r.Set("pic_path", _picPath)
	return nil
}

// GetPicPath PicPath Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetPicPath() string {
	return r._picPath
}

// SetImage is Image Setter
// 主图图片，如果pic_path参数不空，则优先使用pic_path，忽略该参数
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetImage() *model.File {
	return r._image
}

// SetProperties is Properties Setter
// 产品属性
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetProperties() string {
	return r._properties
}

// SetPropertyAlias is PropertyAlias Setter
// 属性别名
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetPropertyAlias(_propertyAlias string) error {
	r._propertyAlias = _propertyAlias
	r.Set("property_alias", _propertyAlias)
	return nil
}

// GetPropertyAlias PropertyAlias Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetPropertyAlias() string {
	return r._propertyAlias
}

// SetInputProperties is InputProperties Setter
// 自定义属性。格式为pid:value;pid:value
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetInputProperties(_inputProperties string) error {
	r._inputProperties = _inputProperties
	r.Set("input_properties", _inputProperties)
	return nil
}

// GetInputProperties InputProperties Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetInputProperties() string {
	return r._inputProperties
}

// SetDealerCostPrice is DealerCostPrice Setter
// 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetDealerCostPrice(_dealerCostPrice string) error {
	r._dealerCostPrice = _dealerCostPrice
	r.Set("dealer_cost_price", _dealerCostPrice)
	return nil
}

// GetDealerCostPrice DealerCostPrice Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetDealerCostPrice() string {
	return r._dealerCostPrice
}

// SetSkuDealerCostPrices is SkuDealerCostPrices Setter
// sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetSkuDealerCostPrices(_skuDealerCostPrices string) error {
	r._skuDealerCostPrices = _skuDealerCostPrices
	r.Set("sku_dealer_cost_prices", _skuDealerCostPrices)
	return nil
}

// GetSkuDealerCostPrices SkuDealerCostPrices Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetSkuDealerCostPrices() string {
	return r._skuDealerCostPrices
}

// SetCategoryId is CategoryId Setter
// 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetStandardRetailPrice is StandardRetailPrice Setter
// 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductUpdateAPIRequest) SetStandardRetailPrice(_standardRetailPrice string) error {
	r._standardRetailPrice = _standardRetailPrice
	r.Set("standard_retail_price", _standardRetailPrice)
	return nil
}

// GetStandardRetailPrice StandardRetailPrice Getter
func (r TaobaoFenxiaoProductUpdateAPIRequest) GetStandardRetailPrice() string {
	return r._standardRetailPrice
}
