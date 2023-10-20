package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalhandoverparcelqueryAPIRequest 获取交接单小包信息 API请求
// cainiao.global.handover.parcel.query
//
// 提供给ISV通过该接口查询小包信息
type CainiaoglobalhandoverparcelqueryAPIRequest struct {
	model.Params
	// 多语言
	_locale string
	// 小包的物流订单号,和小包的国际运单号参数任选其一即可
	_orderCode string
	// 小包的国际运单号，和小包的物流订单号参数任选其一即可
	_trackingNumber string
	// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 用户信息
	_userInfo *UserInfoDto
}

// NewCainiaoglobalhandoverparcelqueryRequest 初始化CainiaoglobalhandoverparcelqueryAPIRequest对象
func NewCainiaoglobalhandoverparcelqueryRequest() *CainiaoglobalhandoverparcelqueryAPIRequest {
	return &CainiaoglobalhandoverparcelqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalhandoverparcelqueryAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.parcel.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalhandoverparcelqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalhandoverparcelqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoglobalhandoverparcelqueryAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoglobalhandoverparcelqueryAPIRequest) GetLocale() string {
	return r._locale
}

// SetOrderCode is OrderCode Setter
// 小包的物流订单号,和小包的国际运单号参数任选其一即可
func (r *CainiaoglobalhandoverparcelqueryAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r CainiaoglobalhandoverparcelqueryAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetTrackingNumber is TrackingNumber Setter
// 小包的国际运单号，和小包的物流订单号参数任选其一即可
func (r *CainiaoglobalhandoverparcelqueryAPIRequest) SetTrackingNumber(_trackingNumber string) error {
	r._trackingNumber = _trackingNumber
	r.Set("tracking_number", _trackingNumber)
	return nil
}

// GetTrackingNumber TrackingNumber Getter
func (r CainiaoglobalhandoverparcelqueryAPIRequest) GetTrackingNumber() string {
	return r._trackingNumber
}

// SetClient is Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoglobalhandoverparcelqueryAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoglobalhandoverparcelqueryAPIRequest) GetClient() string {
	return r._client
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *CainiaoglobalhandoverparcelqueryAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r CainiaoglobalhandoverparcelqueryAPIRequest) GetUserInfo() *UserInfoDto {
	return r._userInfo
}
