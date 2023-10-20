package cainiaohandover

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverCloudprintGetAPIRequest 获取面单云打印数据 API请求
// cainiao.global.handover.cloudprint.get
//
// 提供给ISV通过该接口获取面单云打印数据
type CainiaoGlobalHandoverCloudprintGetAPIRequest struct {
	model.Params
	// 大包运单号
	_trackingNumber string
	// 大包物流单LP号
	_orderCode string
	// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 多语言
	_locale string
	// 用户信息
	_userInfo *UserInfoDto
}

// NewCainiaoGlobalHandoverCloudprintGetRequest 初始化CainiaoGlobalHandoverCloudprintGetAPIRequest对象
func NewCainiaoGlobalHandoverCloudprintGetRequest() *CainiaoGlobalHandoverCloudprintGetAPIRequest {
	return &CainiaoGlobalHandoverCloudprintGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGlobalHandoverCloudprintGetAPIRequest) Reset() {
	r._trackingNumber = ""
	r._orderCode = ""
	r._client = ""
	r._locale = ""
	r._userInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverCloudprintGetAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.cloudprint.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverCloudprintGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalHandoverCloudprintGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrackingNumber is TrackingNumber Setter
// 大包运单号
func (r *CainiaoGlobalHandoverCloudprintGetAPIRequest) SetTrackingNumber(_trackingNumber string) error {
	r._trackingNumber = _trackingNumber
	r.Set("tracking_number", _trackingNumber)
	return nil
}

// GetTrackingNumber TrackingNumber Getter
func (r CainiaoGlobalHandoverCloudprintGetAPIRequest) GetTrackingNumber() string {
	return r._trackingNumber
}

// SetOrderCode is OrderCode Setter
// 大包物流单LP号
func (r *CainiaoGlobalHandoverCloudprintGetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r CainiaoGlobalHandoverCloudprintGetAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetClient is Client Setter
// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverCloudprintGetAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoGlobalHandoverCloudprintGetAPIRequest) GetClient() string {
	return r._client
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverCloudprintGetAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoGlobalHandoverCloudprintGetAPIRequest) GetLocale() string {
	return r._locale
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverCloudprintGetAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r CainiaoGlobalHandoverCloudprintGetAPIRequest) GetUserInfo() *UserInfoDto {
	return r._userInfo
}

var poolCainiaoGlobalHandoverCloudprintGetAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGlobalHandoverCloudprintGetRequest()
	},
}

// GetCainiaoGlobalHandoverCloudprintGetRequest 从 sync.Pool 获取 CainiaoGlobalHandoverCloudprintGetAPIRequest
func GetCainiaoGlobalHandoverCloudprintGetAPIRequest() *CainiaoGlobalHandoverCloudprintGetAPIRequest {
	return poolCainiaoGlobalHandoverCloudprintGetAPIRequest.Get().(*CainiaoGlobalHandoverCloudprintGetAPIRequest)
}

// ReleaseCainiaoGlobalHandoverCloudprintGetAPIRequest 将 CainiaoGlobalHandoverCloudprintGetAPIRequest 放入 sync.Pool
func ReleaseCainiaoGlobalHandoverCloudprintGetAPIRequest(v *CainiaoGlobalHandoverCloudprintGetAPIRequest) {
	v.Reset()
	poolCainiaoGlobalHandoverCloudprintGetAPIRequest.Put(v)
}
