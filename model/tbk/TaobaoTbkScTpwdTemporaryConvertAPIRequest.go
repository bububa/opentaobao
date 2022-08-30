package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScTpwdTemporaryConvertAPIRequest 淘宝客-服务商-淘口令解析&&转链（临时接口） API请求
// taobao.tbk.sc.tpwd.temporary.convert
//
// 支持通过淘口令解析商品id，并提供对应的淘客转链接
type TaobaoTbkScTpwdTemporaryConvertAPIRequest struct {
	model.Params
	// 需要解析的淘口令
	_passwordContent string
	// 1表示商品转通用计划链接，其他值或不传表示优先转营销计划链接
	_dx string
	// 广告位ID，mm_xx_xx_xx pid三段式中的第三段
	_adzoneId int64
	// 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
	_siteId int64
	// 会员人群ID，用于统计人群推广效果
	_ucrowdId int64
}

// NewTaobaoTbkScTpwdTemporaryConvertRequest 初始化TaobaoTbkScTpwdTemporaryConvertAPIRequest对象
func NewTaobaoTbkScTpwdTemporaryConvertRequest() *TaobaoTbkScTpwdTemporaryConvertAPIRequest {
	return &TaobaoTbkScTpwdTemporaryConvertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScTpwdTemporaryConvertAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.tpwd.temporary.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScTpwdTemporaryConvertAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPasswordContent is PasswordContent Setter
// 需要解析的淘口令
func (r *TaobaoTbkScTpwdTemporaryConvertAPIRequest) SetPasswordContent(_passwordContent string) error {
	r._passwordContent = _passwordContent
	r.Set("password_content", _passwordContent)
	return nil
}

// GetPasswordContent PasswordContent Getter
func (r TaobaoTbkScTpwdTemporaryConvertAPIRequest) GetPasswordContent() string {
	return r._passwordContent
}

// SetDx is Dx Setter
// 1表示商品转通用计划链接，其他值或不传表示优先转营销计划链接
func (r *TaobaoTbkScTpwdTemporaryConvertAPIRequest) SetDx(_dx string) error {
	r._dx = _dx
	r.Set("dx", _dx)
	return nil
}

// GetDx Dx Getter
func (r TaobaoTbkScTpwdTemporaryConvertAPIRequest) GetDx() string {
	return r._dx
}

// SetAdzoneId is AdzoneId Setter
// 广告位ID，mm_xx_xx_xx pid三段式中的第三段
func (r *TaobaoTbkScTpwdTemporaryConvertAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkScTpwdTemporaryConvertAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetSiteId is SiteId Setter
// 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
func (r *TaobaoTbkScTpwdTemporaryConvertAPIRequest) SetSiteId(_siteId int64) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaoTbkScTpwdTemporaryConvertAPIRequest) GetSiteId() int64 {
	return r._siteId
}

// SetUcrowdId is UcrowdId Setter
// 会员人群ID，用于统计人群推广效果
func (r *TaobaoTbkScTpwdTemporaryConvertAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaoTbkScTpwdTemporaryConvertAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}
