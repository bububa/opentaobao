package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountTokenApplyAPIRequest 申请免登Open Account Token API请求
// taobao.open.account.token.apply
//
// 申请免登Open Account Token
type TaobaoOpenAccountTokenApplyAPIRequest struct {
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

// NewTaobaoOpenAccountTokenApplyRequest 初始化TaobaoOpenAccountTokenApplyAPIRequest对象
func NewTaobaoOpenAccountTokenApplyRequest() *TaobaoOpenAccountTokenApplyAPIRequest {
	return &TaobaoOpenAccountTokenApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountTokenApplyAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.token.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountTokenApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIsvAccountId is IsvAccountId Setter
// isv自己账号的唯一id
func (r *TaobaoOpenAccountTokenApplyAPIRequest) SetIsvAccountId(_isvAccountId string) error {
	r._isvAccountId = _isvAccountId
	r.Set("isv_account_id", _isvAccountId)
	return nil
}

// GetIsvAccountId IsvAccountId Getter
func (r TaobaoOpenAccountTokenApplyAPIRequest) GetIsvAccountId() string {
	return r._isvAccountId
}

// SetUuid is Uuid Setter
// 用于防重放的唯一id
func (r *TaobaoOpenAccountTokenApplyAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaoOpenAccountTokenApplyAPIRequest) GetUuid() string {
	return r._uuid
}

// SetExt is Ext Setter
// 用于透传一些业务附加参数
func (r *TaobaoOpenAccountTokenApplyAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoOpenAccountTokenApplyAPIRequest) GetExt() string {
	return r._ext
}

// SetLoginStateExpireIn is LoginStateExpireIn Setter
// ISV APP的登录态时长单位是秒
func (r *TaobaoOpenAccountTokenApplyAPIRequest) SetLoginStateExpireIn(_loginStateExpireIn int64) error {
	r._loginStateExpireIn = _loginStateExpireIn
	r.Set("login_state_expire_in", _loginStateExpireIn)
	return nil
}

// GetLoginStateExpireIn LoginStateExpireIn Getter
func (r TaobaoOpenAccountTokenApplyAPIRequest) GetLoginStateExpireIn() int64 {
	return r._loginStateExpireIn
}

// SetOpenAccountId is OpenAccountId Setter
// open account id
func (r *TaobaoOpenAccountTokenApplyAPIRequest) SetOpenAccountId(_openAccountId int64) error {
	r._openAccountId = _openAccountId
	r.Set("open_account_id", _openAccountId)
	return nil
}

// GetOpenAccountId OpenAccountId Getter
func (r TaobaoOpenAccountTokenApplyAPIRequest) GetOpenAccountId() int64 {
	return r._openAccountId
}

// SetTokenTimestamp is TokenTimestamp Setter
// 时间戳单位是毫秒
func (r *TaobaoOpenAccountTokenApplyAPIRequest) SetTokenTimestamp(_tokenTimestamp int64) error {
	r._tokenTimestamp = _tokenTimestamp
	r.Set("token_timestamp", _tokenTimestamp)
	return nil
}

// GetTokenTimestamp TokenTimestamp Getter
func (r TaobaoOpenAccountTokenApplyAPIRequest) GetTokenTimestamp() int64 {
	return r._tokenTimestamp
}
