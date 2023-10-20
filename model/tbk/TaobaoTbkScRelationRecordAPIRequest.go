package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscrelationrecordAPIRequest 淘宝客-服务商-私域会员标记 API请求
// taobao.tbk.sc.relation.record
//
// 服务商使用。支持淘宝客通过入参私域外部ID，获得待私域会员可标记的链接，会员打开该链接后，可帮助媒体自动生成会员运营id进行标记，同时自动跳转到推广落地页。
type TaobaotbkscrelationrecordAPIRequest struct {
	model.Params
	// 淘宝客自有私域用户自定义标记，如自有私域系统账号标记等
	_externalid string
	// 会员备案授权后，跳转的目标页地址
	_redirecturl string
	// 淘宝客其他平台私域用户自定义标记，1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他
	_externaltype int64
	// 授权类型，选2时可用于更新外部用户标记，默认1：1-新绑，2-更新
	_optype int64
	// 人群标签ID，用户备案授权后会自动添加到该人群
	_ucrowdid int64
}

// NewTaobaotbkscrelationrecordRequest 初始化TaobaotbkscrelationrecordAPIRequest对象
func NewTaobaotbkscrelationrecordRequest() *TaobaotbkscrelationrecordAPIRequest {
	return &TaobaotbkscrelationrecordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscrelationrecordAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.relation.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscrelationrecordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscrelationrecordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExternalid is Externalid Setter
// 淘宝客自有私域用户自定义标记，如自有私域系统账号标记等
func (r *TaobaotbkscrelationrecordAPIRequest) SetExternalid(_externalid string) error {
	r._externalid = _externalid
	r.Set("external_id", _externalid)
	return nil
}

// GetExternalid Externalid Getter
func (r TaobaotbkscrelationrecordAPIRequest) GetExternalid() string {
	return r._externalid
}

// SetRedirecturl is Redirecturl Setter
// 会员备案授权后，跳转的目标页地址
func (r *TaobaotbkscrelationrecordAPIRequest) SetRedirecturl(_redirecturl string) error {
	r._redirecturl = _redirecturl
	r.Set("redirect_url", _redirecturl)
	return nil
}

// GetRedirecturl Redirecturl Getter
func (r TaobaotbkscrelationrecordAPIRequest) GetRedirecturl() string {
	return r._redirecturl
}

// SetExternaltype is Externaltype Setter
// 淘宝客其他平台私域用户自定义标记，1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他
func (r *TaobaotbkscrelationrecordAPIRequest) SetExternaltype(_externaltype int64) error {
	r._externaltype = _externaltype
	r.Set("external_type", _externaltype)
	return nil
}

// GetExternaltype Externaltype Getter
func (r TaobaotbkscrelationrecordAPIRequest) GetExternaltype() int64 {
	return r._externaltype
}

// SetOptype is Optype Setter
// 授权类型，选2时可用于更新外部用户标记，默认1：1-新绑，2-更新
func (r *TaobaotbkscrelationrecordAPIRequest) SetOptype(_optype int64) error {
	r._optype = _optype
	r.Set("op_type", _optype)
	return nil
}

// GetOptype Optype Getter
func (r TaobaotbkscrelationrecordAPIRequest) GetOptype() int64 {
	return r._optype
}

// SetUcrowdid is Ucrowdid Setter
// 人群标签ID，用户备案授权后会自动添加到该人群
func (r *TaobaotbkscrelationrecordAPIRequest) SetUcrowdid(_ucrowdid int64) error {
	r._ucrowdid = _ucrowdid
	r.Set("ucrowd_id", _ucrowdid)
	return nil
}

// GetUcrowdid Ucrowdid Getter
func (r TaobaotbkscrelationrecordAPIRequest) GetUcrowdid() int64 {
	return r._ucrowdid
}
