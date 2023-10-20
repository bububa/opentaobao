package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbksctpwdconvertAPIRequest 淘宝客-服务商-淘口令解析&转链 API请求
// taobao.tbk.sc.tpwd.convert
//
// 支持通过淘口令解析商品id，并提供对应的淘客转链接
type TaobaotbksctpwdconvertAPIRequest struct {
	model.Params
	// 需要解析的淘口令
	_passwordContent string
	// 1表示商品转通用计划链接，其他值或不传表示优先转营销计划链接
	_dx string
	// 渠道id
	_relationId string
	// 广告位ID，mm_xx_xx_xx pid三段式中的第三段
	_adzoneId int64
	// 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
	_siteId int64
	// 会员人群ID，用于统计人群推广效果
	_ucrowdId int64
}

// NewTaobaotbksctpwdconvertRequest 初始化TaobaotbksctpwdconvertAPIRequest对象
func NewTaobaotbksctpwdconvertRequest() *TaobaotbksctpwdconvertAPIRequest {
	return &TaobaotbksctpwdconvertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbksctpwdconvertAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.tpwd.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbksctpwdconvertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbksctpwdconvertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPasswordContent is PasswordContent Setter
// 需要解析的淘口令
func (r *TaobaotbksctpwdconvertAPIRequest) SetPasswordContent(_passwordContent string) error {
	r._passwordContent = _passwordContent
	r.Set("password_content", _passwordContent)
	return nil
}

// GetPasswordContent PasswordContent Getter
func (r TaobaotbksctpwdconvertAPIRequest) GetPasswordContent() string {
	return r._passwordContent
}

// SetDx is Dx Setter
// 1表示商品转通用计划链接，其他值或不传表示优先转营销计划链接
func (r *TaobaotbksctpwdconvertAPIRequest) SetDx(_dx string) error {
	r._dx = _dx
	r.Set("dx", _dx)
	return nil
}

// GetDx Dx Getter
func (r TaobaotbksctpwdconvertAPIRequest) GetDx() string {
	return r._dx
}

// SetRelationId is RelationId Setter
// 渠道id
func (r *TaobaotbksctpwdconvertAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaotbksctpwdconvertAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetAdzoneId is AdzoneId Setter
// 广告位ID，mm_xx_xx_xx pid三段式中的第三段
func (r *TaobaotbksctpwdconvertAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaotbksctpwdconvertAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetSiteId is SiteId Setter
// 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
func (r *TaobaotbksctpwdconvertAPIRequest) SetSiteId(_siteId int64) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaotbksctpwdconvertAPIRequest) GetSiteId() int64 {
	return r._siteId
}

// SetUcrowdId is UcrowdId Setter
// 会员人群ID，用于统计人群推广效果
func (r *TaobaotbksctpwdconvertAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaotbksctpwdconvertAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}
