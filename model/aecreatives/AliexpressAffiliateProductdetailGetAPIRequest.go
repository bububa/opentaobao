package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressaffiliateproductdetailgetAPIRequest 联盟商品详情获取接口 API请求
// aliexpress.affiliate.productdetail.get
//
// 联盟推广商品搜索接口，用于搜索联盟推广商品数据
type AliexpressaffiliateproductdetailgetAPIRequest struct {
	model.Params
	// 安全签名
	_appSignature string
	// commission_rate,sale_price
	_fields string
	// 商品ID列表
	_productIds string
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
	_targetLanguage string
	// trackingId
	_trackingId string
	// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
	_country string
}

// NewAliexpressaffiliateproductdetailgetRequest 初始化AliexpressaffiliateproductdetailgetAPIRequest对象
func NewAliexpressaffiliateproductdetailgetRequest() *AliexpressaffiliateproductdetailgetAPIRequest {
	return &AliexpressaffiliateproductdetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressaffiliateproductdetailgetAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.productdetail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressaffiliateproductdetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressaffiliateproductdetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppSignature is AppSignature Setter
// 安全签名
func (r *AliexpressaffiliateproductdetailgetAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressaffiliateproductdetailgetAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetFields is Fields Setter
// commission_rate,sale_price
func (r *AliexpressaffiliateproductdetailgetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressaffiliateproductdetailgetAPIRequest) GetFields() string {
	return r._fields
}

// SetProductIds is ProductIds Setter
// 商品ID列表
func (r *AliexpressaffiliateproductdetailgetAPIRequest) SetProductIds(_productIds string) error {
	r._productIds = _productIds
	r.Set("product_ids", _productIds)
	return nil
}

// GetProductIds ProductIds Getter
func (r AliexpressaffiliateproductdetailgetAPIRequest) GetProductIds() string {
	return r._productIds
}

// SetTargetCurrency is TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressaffiliateproductdetailgetAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressaffiliateproductdetailgetAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetTargetLanguage is TargetLanguage Setter
// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressaffiliateproductdetailgetAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressaffiliateproductdetailgetAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetTrackingId is TrackingId Setter
// trackingId
func (r *AliexpressaffiliateproductdetailgetAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// GetTrackingId TrackingId Getter
func (r AliexpressaffiliateproductdetailgetAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// SetCountry is Country Setter
// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
func (r *AliexpressaffiliateproductdetailgetAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// GetCountry Country Getter
func (r AliexpressaffiliateproductdetailgetAPIRequest) GetCountry() string {
	return r._country
}
