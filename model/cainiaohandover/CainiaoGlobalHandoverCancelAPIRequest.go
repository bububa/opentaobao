package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalhandovercancelAPIRequest 取消交接单 API请求
// cainiao.global.handover.cancel
//
// 提供给ISV通过该接口取消交接单
type CainiaoglobalhandovercancelAPIRequest struct {
	model.Params
	// 要取消的交接物运单号，即大包运单号
	_trackingNumber string
	// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 多语言
	_locale string
	// 系统自动生成
	_userInfo *UserInfoDto
	// 要取消的交接单id
	_handoverOrderId int64
	// 要取消的交接物id，即大包id
	_handoverContentId int64
}

// NewCainiaoglobalhandovercancelRequest 初始化CainiaoglobalhandovercancelAPIRequest对象
func NewCainiaoglobalhandovercancelRequest() *CainiaoglobalhandovercancelAPIRequest {
	return &CainiaoglobalhandovercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalhandovercancelAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalhandovercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalhandovercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrackingNumber is TrackingNumber Setter
// 要取消的交接物运单号，即大包运单号
func (r *CainiaoglobalhandovercancelAPIRequest) SetTrackingNumber(_trackingNumber string) error {
	r._trackingNumber = _trackingNumber
	r.Set("tracking_number", _trackingNumber)
	return nil
}

// GetTrackingNumber TrackingNumber Getter
func (r CainiaoglobalhandovercancelAPIRequest) GetTrackingNumber() string {
	return r._trackingNumber
}

// SetClient is Client Setter
// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoglobalhandovercancelAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoglobalhandovercancelAPIRequest) GetClient() string {
	return r._client
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoglobalhandovercancelAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoglobalhandovercancelAPIRequest) GetLocale() string {
	return r._locale
}

// SetUserInfo is UserInfo Setter
// 系统自动生成
func (r *CainiaoglobalhandovercancelAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r CainiaoglobalhandovercancelAPIRequest) GetUserInfo() *UserInfoDto {
	return r._userInfo
}

// SetHandoverOrderId is HandoverOrderId Setter
// 要取消的交接单id
func (r *CainiaoglobalhandovercancelAPIRequest) SetHandoverOrderId(_handoverOrderId int64) error {
	r._handoverOrderId = _handoverOrderId
	r.Set("handover_order_id", _handoverOrderId)
	return nil
}

// GetHandoverOrderId HandoverOrderId Getter
func (r CainiaoglobalhandovercancelAPIRequest) GetHandoverOrderId() int64 {
	return r._handoverOrderId
}

// SetHandoverContentId is HandoverContentId Setter
// 要取消的交接物id，即大包id
func (r *CainiaoglobalhandovercancelAPIRequest) SetHandoverContentId(_handoverContentId int64) error {
	r._handoverContentId = _handoverContentId
	r.Set("handover_content_id", _handoverContentId)
	return nil
}

// GetHandoverContentId HandoverContentId Getter
func (r CainiaoglobalhandovercancelAPIRequest) GetHandoverContentId() int64 {
	return r._handoverContentId
}
