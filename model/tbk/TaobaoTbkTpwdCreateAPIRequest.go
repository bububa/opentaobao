package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbktpwdcreateAPIRequest 淘宝客-公用-淘口令生成 API请求
// taobao.tbk.tpwd.create
//
// 提供淘口令生成接口。提交需要生成淘口令的内容、logo、url等参数，生成淘口令，其中关键信息为￥SADadW￥，后续可基于淘口令进行文案包装组装用于传播。
type TaobaotbktpwdcreateAPIRequest struct {
	model.Params
	// 联盟官方渠道获取的淘客推广链接，请注意，不要随意篡改官方生成的链接，否则可能无法生成淘口令
	_url string
	// 兼容旧版本api参数，无实际作用
	_text string
	// 兼容旧版本api参数，无实际作用
	_logo string
	// 兼容旧版本api参数，无实际作用
	_ext string
	// 兼容旧版本api参数，无实际作用
	_userId string
}

// NewTaobaotbktpwdcreateRequest 初始化TaobaotbktpwdcreateAPIRequest对象
func NewTaobaotbktpwdcreateRequest() *TaobaotbktpwdcreateAPIRequest {
	return &TaobaotbktpwdcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbktpwdcreateAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.tpwd.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbktpwdcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbktpwdcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUrl is Url Setter
// 联盟官方渠道获取的淘客推广链接，请注意，不要随意篡改官方生成的链接，否则可能无法生成淘口令
func (r *TaobaotbktpwdcreateAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r TaobaotbktpwdcreateAPIRequest) GetUrl() string {
	return r._url
}

// SetText is Text Setter
// 兼容旧版本api参数，无实际作用
func (r *TaobaotbktpwdcreateAPIRequest) SetText(_text string) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// GetText Text Getter
func (r TaobaotbktpwdcreateAPIRequest) GetText() string {
	return r._text
}

// SetLogo is Logo Setter
// 兼容旧版本api参数，无实际作用
func (r *TaobaotbktpwdcreateAPIRequest) SetLogo(_logo string) error {
	r._logo = _logo
	r.Set("logo", _logo)
	return nil
}

// GetLogo Logo Getter
func (r TaobaotbktpwdcreateAPIRequest) GetLogo() string {
	return r._logo
}

// SetExt is Ext Setter
// 兼容旧版本api参数，无实际作用
func (r *TaobaotbktpwdcreateAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaotbktpwdcreateAPIRequest) GetExt() string {
	return r._ext
}

// SetUserId is UserId Setter
// 兼容旧版本api参数，无实际作用
func (r *TaobaotbktpwdcreateAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaotbktpwdcreateAPIRequest) GetUserId() string {
	return r._userId
}
