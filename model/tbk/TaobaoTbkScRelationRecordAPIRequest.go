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
	_externalId string
	// 会员备案授权后，跳转的目标页地址
	_redirectUrl string
	// 淘宝客其他平台私域用户自定义标记，1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他
	_externalType int64
	// 授权类型，选2时可用于更新外部用户标记，默认1：1-新绑，2-更新
	_opType int64
	// 人群标签ID，用户备案授权后会自动添加到该人群
	_ucrowdId int64
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

// SetExternalId is ExternalId Setter
// 淘宝客自有私域用户自定义标记，如自有私域系统账号标记等
func (r *TaobaotbkscrelationrecordAPIRequest) SetExternalId(_externalId string) error {
	r._externalId = _externalId
	r.Set("external_id", _externalId)
	return nil
}

// GetExternalId ExternalId Getter
func (r TaobaotbkscrelationrecordAPIRequest) GetExternalId() string {
	return r._externalId
}

// SetRedirectUrl is RedirectUrl Setter
// 会员备案授权后，跳转的目标页地址
func (r *TaobaotbkscrelationrecordAPIRequest) SetRedirectUrl(_redirectUrl string) error {
	r._redirectUrl = _redirectUrl
	r.Set("redirect_url", _redirectUrl)
	return nil
}

// GetRedirectUrl RedirectUrl Getter
func (r TaobaotbkscrelationrecordAPIRequest) GetRedirectUrl() string {
	return r._redirectUrl
}

// SetExternalType is ExternalType Setter
// 淘宝客其他平台私域用户自定义标记，1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他
func (r *TaobaotbkscrelationrecordAPIRequest) SetExternalType(_externalType int64) error {
	r._externalType = _externalType
	r.Set("external_type", _externalType)
	return nil
}

// GetExternalType ExternalType Getter
func (r TaobaotbkscrelationrecordAPIRequest) GetExternalType() int64 {
	return r._externalType
}

// SetOpType is OpType Setter
// 授权类型，选2时可用于更新外部用户标记，默认1：1-新绑，2-更新
func (r *TaobaotbkscrelationrecordAPIRequest) SetOpType(_opType int64) error {
	r._opType = _opType
	r.Set("op_type", _opType)
	return nil
}

// GetOpType OpType Getter
func (r TaobaotbkscrelationrecordAPIRequest) GetOpType() int64 {
	return r._opType
}

// SetUcrowdId is UcrowdId Setter
// 人群标签ID，用户备案授权后会自动添加到该人群
func (r *TaobaotbkscrelationrecordAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaotbkscrelationrecordAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}
