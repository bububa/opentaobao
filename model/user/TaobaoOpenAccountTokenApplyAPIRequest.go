package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenaccounttokenapplyAPIRequest 申请免登Open Account Token API请求
// taobao.open.account.token.apply
//
// 申请免登Open Account Token
type TaobaoopenaccounttokenapplyAPIRequest struct {
	model.Params
	// isv自己账号的唯一id
	_isvAccountId string
	// 用于防重放的唯一id
	_uuid string
	// 用于透传一些业务附加参数
	_ext string
	// ISV APP的登录态时长单位是秒
	_loginStateExpireIn int64
	// open account id
	_openAccountId int64
	// 时间戳单位是毫秒
	_tokenTimestamp int64
}

// NewTaobaoopenaccounttokenapplyRequest 初始化TaobaoopenaccounttokenapplyAPIRequest对象
func NewTaobaoopenaccounttokenapplyRequest() *TaobaoopenaccounttokenapplyAPIRequest {
	return &TaobaoopenaccounttokenapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenaccounttokenapplyAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.token.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenaccounttokenapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenaccounttokenapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvAccountId is IsvAccountId Setter
// isv自己账号的唯一id
func (r *TaobaoopenaccounttokenapplyAPIRequest) SetIsvAccountId(_isvAccountId string) error {
	r._isvAccountId = _isvAccountId
	r.Set("isv_account_id", _isvAccountId)
	return nil
}

// GetIsvAccountId IsvAccountId Getter
func (r TaobaoopenaccounttokenapplyAPIRequest) GetIsvAccountId() string {
	return r._isvAccountId
}

// SetUuid is Uuid Setter
// 用于防重放的唯一id
func (r *TaobaoopenaccounttokenapplyAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaoopenaccounttokenapplyAPIRequest) GetUuid() string {
	return r._uuid
}

// SetExt is Ext Setter
// 用于透传一些业务附加参数
func (r *TaobaoopenaccounttokenapplyAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoopenaccounttokenapplyAPIRequest) GetExt() string {
	return r._ext
}

// SetLoginStateExpireIn is LoginStateExpireIn Setter
// ISV APP的登录态时长单位是秒
func (r *TaobaoopenaccounttokenapplyAPIRequest) SetLoginStateExpireIn(_loginStateExpireIn int64) error {
	r._loginStateExpireIn = _loginStateExpireIn
	r.Set("login_state_expire_in", _loginStateExpireIn)
	return nil
}

// GetLoginStateExpireIn LoginStateExpireIn Getter
func (r TaobaoopenaccounttokenapplyAPIRequest) GetLoginStateExpireIn() int64 {
	return r._loginStateExpireIn
}

// SetOpenAccountId is OpenAccountId Setter
// open account id
func (r *TaobaoopenaccounttokenapplyAPIRequest) SetOpenAccountId(_openAccountId int64) error {
	r._openAccountId = _openAccountId
	r.Set("open_account_id", _openAccountId)
	return nil
}

// GetOpenAccountId OpenAccountId Getter
func (r TaobaoopenaccounttokenapplyAPIRequest) GetOpenAccountId() int64 {
	return r._openAccountId
}

// SetTokenTimestamp is TokenTimestamp Setter
// 时间戳单位是毫秒
func (r *TaobaoopenaccounttokenapplyAPIRequest) SetTokenTimestamp(_tokenTimestamp int64) error {
	r._tokenTimestamp = _tokenTimestamp
	r.Set("token_timestamp", _tokenTimestamp)
	return nil
}

// GetTokenTimestamp TokenTimestamp Getter
func (r TaobaoopenaccounttokenapplyAPIRequest) GetTokenTimestamp() int64 {
	return r._tokenTimestamp
}
