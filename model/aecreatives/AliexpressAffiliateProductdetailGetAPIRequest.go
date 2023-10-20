package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateProductdetailGetAPIRequest 联盟商品详情获取接口 API请求
// aliexpress.affiliate.productdetail.get
//
// 联盟推广商品搜索接口，用于搜索联盟推广商品数据
type AliexpressAffiliateProductdetailGetAPIRequest struct {
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

// NewAliexpressAffiliateProductdetailGetRequest 初始化AliexpressAffiliateProductdetailGetAPIRequest对象
func NewAliexpressAffiliateProductdetailGetRequest() *AliexpressAffiliateProductdetailGetAPIRequest {
	return &AliexpressAffiliateProductdetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateProductdetailGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.productdetail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateProductdetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAffiliateProductdetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppSignature is AppSignature Setter
// 安全签名
func (r *AliexpressAffiliateProductdetailGetAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressAffiliateProductdetailGetAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetFields is Fields Setter
// commission_rate,sale_price
func (r *AliexpressAffiliateProductdetailGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressAffiliateProductdetailGetAPIRequest) GetFields() string {
	return r._fields
}

// SetProductIds is ProductIds Setter
// 商品ID列表
func (r *AliexpressAffiliateProductdetailGetAPIRequest) SetProductIds(_productIds string) error {
	r._productIds = _productIds
	r.Set("product_ids", _productIds)
	return nil
}

// GetProductIds ProductIds Getter
func (r AliexpressAffiliateProductdetailGetAPIRequest) GetProductIds() string {
	return r._productIds
}

// SetTargetCurrency is TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateProductdetailGetAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressAffiliateProductdetailGetAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetTargetLanguage is TargetLanguage Setter
// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressAffiliateProductdetailGetAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressAffiliateProductdetailGetAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetTrackingId is TrackingId Setter
// trackingId
func (r *AliexpressAffiliateProductdetailGetAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// GetTrackingId TrackingId Getter
func (r AliexpressAffiliateProductdetailGetAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// SetCountry is Country Setter
// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateProductdetailGetAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// GetCountry Country Getter
func (r AliexpressAffiliateProductdetailGetAPIRequest) GetCountry() string {
	return r._country
}
