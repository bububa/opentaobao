package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallproductspecaddAPIRequest 添加产品规格 API请求
// tmall.product.spec.add
//
// 增加产品规格
type TmallproductspecaddAPIRequest struct {
	model.Params
	// 存放产品规格认证类型-认证图片url映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为调用tmall.product.spec.pic.upload返回的认证图片url文本
	_certifiedPicStr string
	// 存放产品规格认证类型-认证文本映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为认证文本值
	_certifiedTxtStr string
	// 产品的规格属性
	_specProps string
	// 规格属性别名,只允许传颜色别名
	_specPropsAlias string
	// 用户自定义销售属性，结构：pid1:value1;pid2:value2。在
	_customerSpecProps string
	// 产品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1
	_changeProp string
	// 产品二维码
	_barcode string
	// 产品货号
	_productCode string
	// 产品上市时间
	_marketTime string
	// 产品ID
	_productId int64
	// 产品规格吊牌价，以分为单位，无默认值，上限999999999
	_labelPrice int64
	// 产品图片
	_image *model.File
}

// NewTmallproductspecaddRequest 初始化TmallproductspecaddAPIRequest对象
func NewTmallproductspecaddRequest() *TmallproductspecaddAPIRequest {
	return &TmallproductspecaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallproductspecaddAPIRequest) GetApiMethodName() string {
	return "tmall.product.spec.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallproductspecaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallproductspecaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCertifiedPicStr is CertifiedPicStr Setter
// 存放产品规格认证类型-认证图片url映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为调用tmall.product.spec.pic.upload返回的认证图片url文本
func (r *TmallproductspecaddAPIRequest) SetCertifiedPicStr(_certifiedPicStr string) error {
	r._certifiedPicStr = _certifiedPicStr
	r.Set("certified_pic_str", _certifiedPicStr)
	return nil
}

// GetCertifiedPicStr CertifiedPicStr Getter
func (r TmallproductspecaddAPIRequest) GetCertifiedPicStr() string {
	return r._certifiedPicStr
}

// SetCertifiedTxtStr is CertifiedTxtStr Setter
// 存放产品规格认证类型-认证文本映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为认证文本值
func (r *TmallproductspecaddAPIRequest) SetCertifiedTxtStr(_certifiedTxtStr string) error {
	r._certifiedTxtStr = _certifiedTxtStr
	r.Set("certified_txt_str", _certifiedTxtStr)
	return nil
}

// GetCertifiedTxtStr CertifiedTxtStr Getter
func (r TmallproductspecaddAPIRequest) GetCertifiedTxtStr() string {
	return r._certifiedTxtStr
}

// SetSpecProps is SpecProps Setter
// 产品的规格属性
func (r *TmallproductspecaddAPIRequest) SetSpecProps(_specProps string) error {
	r._specProps = _specProps
	r.Set("spec_props", _specProps)
	return nil
}

// GetSpecProps SpecProps Getter
func (r TmallproductspecaddAPIRequest) GetSpecProps() string {
	return r._specProps
}

// SetSpecPropsAlias is SpecPropsAlias Setter
// 规格属性别名,只允许传颜色别名
func (r *TmallproductspecaddAPIRequest) SetSpecPropsAlias(_specPropsAlias string) error {
	r._specPropsAlias = _specPropsAlias
	r.Set("spec_props_alias", _specPropsAlias)
	return nil
}

// GetSpecPropsAlias SpecPropsAlias Getter
func (r TmallproductspecaddAPIRequest) GetSpecPropsAlias() string {
	return r._specPropsAlias
}

// SetCustomerSpecProps is CustomerSpecProps Setter
// 用户自定义销售属性，结构：pid1:value1;pid2:value2。在
func (r *TmallproductspecaddAPIRequest) SetCustomerSpecProps(_customerSpecProps string) error {
	r._customerSpecProps = _customerSpecProps
	r.Set("customer_spec_props", _customerSpecProps)
	return nil
}

// GetCustomerSpecProps CustomerSpecProps Getter
func (r TmallproductspecaddAPIRequest) GetCustomerSpecProps() string {
	return r._customerSpecProps
}

// SetChangeProp is ChangeProp Setter
// 产品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1
func (r *TmallproductspecaddAPIRequest) SetChangeProp(_changeProp string) error {
	r._changeProp = _changeProp
	r.Set("change_prop", _changeProp)
	return nil
}

// GetChangeProp ChangeProp Getter
func (r TmallproductspecaddAPIRequest) GetChangeProp() string {
	return r._changeProp
}

// SetBarcode is Barcode Setter
// 产品二维码
func (r *TmallproductspecaddAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r TmallproductspecaddAPIRequest) GetBarcode() string {
	return r._barcode
}

// SetProductCode is ProductCode Setter
// 产品货号
func (r *TmallproductspecaddAPIRequest) SetProductCode(_productCode string) error {
	r._productCode = _productCode
	r.Set("product_code", _productCode)
	return nil
}

// GetProductCode ProductCode Getter
func (r TmallproductspecaddAPIRequest) GetProductCode() string {
	return r._productCode
}

// SetMarketTime is MarketTime Setter
// 产品上市时间
func (r *TmallproductspecaddAPIRequest) SetMarketTime(_marketTime string) error {
	r._marketTime = _marketTime
	r.Set("market_time", _marketTime)
	return nil
}

// GetMarketTime MarketTime Getter
func (r TmallproductspecaddAPIRequest) GetMarketTime() string {
	return r._marketTime
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TmallproductspecaddAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallproductspecaddAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetLabelPrice is LabelPrice Setter
// 产品规格吊牌价，以分为单位，无默认值，上限999999999
func (r *TmallproductspecaddAPIRequest) SetLabelPrice(_labelPrice int64) error {
	r._labelPrice = _labelPrice
	r.Set("label_price", _labelPrice)
	return nil
}

// GetLabelPrice LabelPrice Getter
func (r TmallproductspecaddAPIRequest) GetLabelPrice() int64 {
	return r._labelPrice
}

// SetImage is Image Setter
// 产品图片
func (r *TmallproductspecaddAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TmallproductspecaddAPIRequest) GetImage() *model.File {
	return r._image
}
