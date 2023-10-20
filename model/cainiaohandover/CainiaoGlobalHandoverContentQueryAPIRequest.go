package cainiaohandover

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverContentQueryAPIRequest 查询大包详情 API请求
// cainiao.global.handover.content.query
//
// 查询大包详情
type CainiaoGlobalHandoverContentQueryAPIRequest struct {
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

// NewCainiaoGlobalHandoverContentQueryRequest 初始化CainiaoGlobalHandoverContentQueryAPIRequest对象
func NewCainiaoGlobalHandoverContentQueryRequest() *CainiaoGlobalHandoverContentQueryAPIRequest {
	return &CainiaoGlobalHandoverContentQueryAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGlobalHandoverContentQueryAPIRequest) Reset() {
	r._orderCode = ""
	r._trackingNumber = ""
	r._client = ""
	r._locale = ""
	r._userInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.content.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 交接物物流订单编码,和交接物运单号参数可以任选其一即可
func (r *CainiaoGlobalHandoverContentQueryAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetTrackingNumber is TrackingNumber Setter
// 交接物运单号，和交接物物流订单编码参数任选其一即可
func (r *CainiaoGlobalHandoverContentQueryAPIRequest) SetTrackingNumber(_trackingNumber string) error {
	r._trackingNumber = _trackingNumber
	r.Set("tracking_number", _trackingNumber)
	return nil
}

// GetTrackingNumber TrackingNumber Getter
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetTrackingNumber() string {
	return r._trackingNumber
}

// SetClient is Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverContentQueryAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetClient() string {
	return r._client
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverContentQueryAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetLocale() string {
	return r._locale
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverContentQueryAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r CainiaoGlobalHandoverContentQueryAPIRequest) GetUserInfo() *UserInfoDto {
	return r._userInfo
}

var poolCainiaoGlobalHandoverContentQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGlobalHandoverContentQueryRequest()
	},
}

// GetCainiaoGlobalHandoverContentQueryRequest 从 sync.Pool 获取 CainiaoGlobalHandoverContentQueryAPIRequest
func GetCainiaoGlobalHandoverContentQueryAPIRequest() *CainiaoGlobalHandoverContentQueryAPIRequest {
	return poolCainiaoGlobalHandoverContentQueryAPIRequest.Get().(*CainiaoGlobalHandoverContentQueryAPIRequest)
}

// ReleaseCainiaoGlobalHandoverContentQueryAPIRequest 将 CainiaoGlobalHandoverContentQueryAPIRequest 放入 sync.Pool
func ReleaseCainiaoGlobalHandoverContentQueryAPIRequest(v *CainiaoGlobalHandoverContentQueryAPIRequest) {
	v.Reset()
	poolCainiaoGlobalHandoverContentQueryAPIRequest.Put(v)
}
