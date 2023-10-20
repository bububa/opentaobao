package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalhandovercontentqueryAPIRequest 查询大包详情 API请求
// cainiao.global.handover.content.query
//
// 查询大包详情
type CainiaoglobalhandovercontentqueryAPIRequest struct {
	model.Params
	// 交接物物流订单编码,和交接物运单号参数可以任选其一即可
	_orderCode string
	// 交接物运单号，和交接物物流订单编码参数任选其一即可
	_trackingNumber string
	// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 多语言
	_locale string
	// 用户信息
	_userInfo *UserInfoDto
}

// NewCainiaoglobalhandovercontentqueryRequest 初始化CainiaoglobalhandovercontentqueryAPIRequest对象
func NewCainiaoglobalhandovercontentqueryRequest() *CainiaoglobalhandovercontentqueryAPIRequest {
	return &CainiaoglobalhandovercontentqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalhandovercontentqueryAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.content.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalhandovercontentqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalhandovercontentqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 交接物物流订单编码,和交接物运单号参数可以任选其一即可
func (r *CainiaoglobalhandovercontentqueryAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r CainiaoglobalhandovercontentqueryAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetTrackingNumber is TrackingNumber Setter
// 交接物运单号，和交接物物流订单编码参数任选其一即可
func (r *CainiaoglobalhandovercontentqueryAPIRequest) SetTrackingNumber(_trackingNumber string) error {
	r._trackingNumber = _trackingNumber
	r.Set("tracking_number", _trackingNumber)
	return nil
}

// GetTrackingNumber TrackingNumber Getter
func (r CainiaoglobalhandovercontentqueryAPIRequest) GetTrackingNumber() string {
	return r._trackingNumber
}

// SetClient is Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoglobalhandovercontentqueryAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoglobalhandovercontentqueryAPIRequest) GetClient() string {
	return r._client
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoglobalhandovercontentqueryAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoglobalhandovercontentqueryAPIRequest) GetLocale() string {
	return r._locale
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *CainiaoglobalhandovercontentqueryAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r CainiaoglobalhandovercontentqueryAPIRequest) GetUserInfo() *UserInfoDto {
	return r._userInfo
}
