package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateProductSmartmatchAPIRequest 联盟物料智能推荐api API请求
// aliexpress.affiliate.product.smartmatch
//
// 联盟物料算法智能推荐
type AliexpressAffiliateProductSmartmatchAPIRequest struct {
	model.Params
	// 接入APP信息
	_app string
	// 请求签名
	_appSignature string
	// 设备信息
	_device string
	// adid或者idfa
	_deviceId string
	// 返回字段列表
	_fields string
	// 关键词
	_keywords string
	// 商品ID
	_productId string
	// 站点信息
	_site string
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
	_targetLanguage string
	// trackingId
	_trackingId string
	// 用户信息
	_user string
	// 请求页数
	_pageNo int64
	// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
	_country string
}

// NewAliexpressAffiliateProductSmartmatchRequest 初始化AliexpressAffiliateProductSmartmatchAPIRequest对象
func NewAliexpressAffiliateProductSmartmatchRequest() *AliexpressAffiliateProductSmartmatchAPIRequest {
	return &AliexpressAffiliateProductSmartmatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.product.smartmatch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is App Setter
// 接入APP信息
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetApp(_app string) error {
	r._app = _app
	r.Set("app", _app)
	return nil
}

// Get App Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetApp() string {
	return r._app
}

// Set is AppSignature Setter
// 请求签名
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// Get AppSignature Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// Set is Device Setter
// 设备信息
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetDevice(_device string) error {
	r._device = _device
	r.Set("device", _device)
	return nil
}

// Get Device Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetDevice() string {
	return r._device
}

// Set is DeviceId Setter
// adid或者idfa
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// Get DeviceId Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// Set is Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// Get Fields Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetFields() string {
	return r._fields
}

// Set is Keywords Setter
// 关键词
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetKeywords(_keywords string) error {
	r._keywords = _keywords
	r.Set("keywords", _keywords)
	return nil
}

// Get Keywords Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetKeywords() string {
	return r._keywords
}

// Set is ProductId Setter
// 商品ID
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetProductId() string {
	return r._productId
}

// Set is Site Setter
// 站点信息
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetSite(_site string) error {
	r._site = _site
	r.Set("site", _site)
	return nil
}

// Get Site Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetSite() string {
	return r._site
}

// Set is TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// Get TargetCurrency Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// Set is TargetLanguage Setter
// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// Get TargetLanguage Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// Set is TrackingId Setter
// trackingId
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// Get TrackingId Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// Set is User Setter
// 用户信息
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetUser(_user string) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// Get User Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetUser() string {
	return r._user
}

// Set is PageNo Setter
// 请求页数
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is Country Setter
// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateProductSmartmatchAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// Get Country Getter
func (r AliexpressAffiliateProductSmartmatchAPIRequest) GetCountry() string {
	return r._country
}
