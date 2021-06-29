package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加产品规格 API请求
tmall.product.spec.add

增加产品规格
*/
type TmallProductSpecAddRequest struct {
    model.Params
    // 产品ID
    _productId   int64
    // 存放产品规格认证类型-认证图片url映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为调用tmall.product.spec.pic.upload返回的认证图片url文本
    _certifiedPicStr   string
    // 产品规格吊牌价，以分为单位，无默认值，上限999999999
    _labelPrice   int64
    // 存放产品规格认证类型-认证文本映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为认证文本值
    _certifiedTxtStr   string
    // 产品的规格属性
    _specProps   string
    // 规格属性别名,只允许传颜色别名
    _specPropsAlias   string
    // 用户自定义销售属性，结构：pid1:value1;pid2:value2。在
    _customerSpecProps   string
    // 产品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1
    _changeProp   string
    // 产品图片
    _image   *model.File
    // 产品二维码
    _barcode   string
    // 产品货号
    _productCode   string
    // 产品上市时间
    _marketTime   string
}

// 初始化TmallProductSpecAddRequest对象
func NewTmallProductSpecAddRequest() *TmallProductSpecAddRequest{
    return &TmallProductSpecAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductSpecAddRequest) GetApiMethodName() string {
    return "tmall.product.spec.add"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductSpecAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TmallProductSpecAddRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallProductSpecAddRequest) GetProductId() int64 {
    return r._productId
}
// CertifiedPicStr Setter
// 存放产品规格认证类型-认证图片url映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为调用tmall.product.spec.pic.upload返回的认证图片url文本
func (r *TmallProductSpecAddRequest) SetCertifiedPicStr(_certifiedPicStr string) error {
    r._certifiedPicStr = _certifiedPicStr
    r.Set("certified_pic_str", _certifiedPicStr)
    return nil
}

// CertifiedPicStr Getter
func (r TmallProductSpecAddRequest) GetCertifiedPicStr() string {
    return r._certifiedPicStr
}
// LabelPrice Setter
// 产品规格吊牌价，以分为单位，无默认值，上限999999999
func (r *TmallProductSpecAddRequest) SetLabelPrice(_labelPrice int64) error {
    r._labelPrice = _labelPrice
    r.Set("label_price", _labelPrice)
    return nil
}

// LabelPrice Getter
func (r TmallProductSpecAddRequest) GetLabelPrice() int64 {
    return r._labelPrice
}
// CertifiedTxtStr Setter
// 存放产品规格认证类型-认证文本映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为认证文本值
func (r *TmallProductSpecAddRequest) SetCertifiedTxtStr(_certifiedTxtStr string) error {
    r._certifiedTxtStr = _certifiedTxtStr
    r.Set("certified_txt_str", _certifiedTxtStr)
    return nil
}

// CertifiedTxtStr Getter
func (r TmallProductSpecAddRequest) GetCertifiedTxtStr() string {
    return r._certifiedTxtStr
}
// SpecProps Setter
// 产品的规格属性
func (r *TmallProductSpecAddRequest) SetSpecProps(_specProps string) error {
    r._specProps = _specProps
    r.Set("spec_props", _specProps)
    return nil
}

// SpecProps Getter
func (r TmallProductSpecAddRequest) GetSpecProps() string {
    return r._specProps
}
// SpecPropsAlias Setter
// 规格属性别名,只允许传颜色别名
func (r *TmallProductSpecAddRequest) SetSpecPropsAlias(_specPropsAlias string) error {
    r._specPropsAlias = _specPropsAlias
    r.Set("spec_props_alias", _specPropsAlias)
    return nil
}

// SpecPropsAlias Getter
func (r TmallProductSpecAddRequest) GetSpecPropsAlias() string {
    return r._specPropsAlias
}
// CustomerSpecProps Setter
// 用户自定义销售属性，结构：pid1:value1;pid2:value2。在
func (r *TmallProductSpecAddRequest) SetCustomerSpecProps(_customerSpecProps string) error {
    r._customerSpecProps = _customerSpecProps
    r.Set("customer_spec_props", _customerSpecProps)
    return nil
}

// CustomerSpecProps Getter
func (r TmallProductSpecAddRequest) GetCustomerSpecProps() string {
    return r._customerSpecProps
}
// ChangeProp Setter
// 产品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1
func (r *TmallProductSpecAddRequest) SetChangeProp(_changeProp string) error {
    r._changeProp = _changeProp
    r.Set("change_prop", _changeProp)
    return nil
}

// ChangeProp Getter
func (r TmallProductSpecAddRequest) GetChangeProp() string {
    return r._changeProp
}
// Image Setter
// 产品图片
func (r *TmallProductSpecAddRequest) SetImage(_image *model.File) error {
    r._image = _image
    r.Set("image", _image)
    return nil
}

// Image Getter
func (r TmallProductSpecAddRequest) GetImage() *model.File {
    return r._image
}
// Barcode Setter
// 产品二维码
func (r *TmallProductSpecAddRequest) SetBarcode(_barcode string) error {
    r._barcode = _barcode
    r.Set("barcode", _barcode)
    return nil
}

// Barcode Getter
func (r TmallProductSpecAddRequest) GetBarcode() string {
    return r._barcode
}
// ProductCode Setter
// 产品货号
func (r *TmallProductSpecAddRequest) SetProductCode(_productCode string) error {
    r._productCode = _productCode
    r.Set("product_code", _productCode)
    return nil
}

// ProductCode Getter
func (r TmallProductSpecAddRequest) GetProductCode() string {
    return r._productCode
}
// MarketTime Setter
// 产品上市时间
func (r *TmallProductSpecAddRequest) SetMarketTime(_marketTime string) error {
    r._marketTime = _marketTime
    r.Set("market_time", _marketTime)
    return nil
}

// MarketTime Getter
func (r TmallProductSpecAddRequest) GetMarketTime() string {
    return r._marketTime
}
