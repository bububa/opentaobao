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
	_passwordcontent string
	// 1表示商品转通用计划链接，其他值或不传表示优先转营销计划链接
	_dx string
	// 渠道id
	_relationid string
	// 广告位ID，mm_xx_xx_xx pid三段式中的第三段
	_adzoneid int64
	// 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
	_siteid int64
	// 会员人群ID，用于统计人群推广效果
	_ucrowdid int64
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

// SetPasswordcontent is Passwordcontent Setter
// 需要解析的淘口令
func (r *TaobaotbksctpwdconvertAPIRequest) SetPasswordcontent(_passwordcontent string) error {
	r._passwordcontent = _passwordcontent
	r.Set("password_content", _passwordcontent)
	return nil
}

// GetPasswordcontent Passwordcontent Getter
func (r TaobaotbksctpwdconvertAPIRequest) GetPasswordcontent() string {
	return r._passwordcontent
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

// SetRelationid is Relationid Setter
// 渠道id
func (r *TaobaotbksctpwdconvertAPIRequest) SetRelationid(_relationid string) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbksctpwdconvertAPIRequest) GetRelationid() string {
	return r._relationid
}

// SetAdzoneid is Adzoneid Setter
// 广告位ID，mm_xx_xx_xx pid三段式中的第三段
func (r *TaobaotbksctpwdconvertAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbksctpwdconvertAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}

// SetSiteid is Siteid Setter
// 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
func (r *TaobaotbksctpwdconvertAPIRequest) SetSiteid(_siteid int64) error {
	r._siteid = _siteid
	r.Set("site_id", _siteid)
	return nil
}

// GetSiteid Siteid Getter
func (r TaobaotbksctpwdconvertAPIRequest) GetSiteid() int64 {
	return r._siteid
}

// SetUcrowdid is Ucrowdid Setter
// 会员人群ID，用于统计人群推广效果
func (r *TaobaotbksctpwdconvertAPIRequest) SetUcrowdid(_ucrowdid int64) error {
	r._ucrowdid = _ucrowdid
	r.Set("ucrowd_id", _ucrowdid)
	return nil
}

// GetUcrowdid Ucrowdid Getter
func (r TaobaotbksctpwdconvertAPIRequest) GetUcrowdid() int64 {
	return r._ucrowdid
}
