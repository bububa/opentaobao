package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodegoperationdoluckydrawAPIRequest 激励抽奖 API请求
// taobao.degoperation.do.luckydraw
//
// 激励平台抽奖接口。用户可以通过接口完成抽奖功能
type TaobaodegoperationdoluckydrawAPIRequest struct {
	model.Params
	// 后台活动配置appkey
	_degAppKey string
	// 后台活动配置eventkey
	_degEventKey string
	// 传参信息
	_degAccessToken string
	// 前端标识
	_source string
	// 设备uuid
	_uuid string
	// 参数校验
	_paramSign string
}

// NewTaobaodegoperationdoluckydrawRequest 初始化TaobaodegoperationdoluckydrawAPIRequest对象
func NewTaobaodegoperationdoluckydrawRequest() *TaobaodegoperationdoluckydrawAPIRequest {
	return &TaobaodegoperationdoluckydrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodegoperationdoluckydrawAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.do.luckydraw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodegoperationdoluckydrawAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodegoperationdoluckydrawAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDegAppKey is DegAppKey Setter
// 后台活动配置appkey
func (r *TaobaodegoperationdoluckydrawAPIRequest) SetDegAppKey(_degAppKey string) error {
	r._degAppKey = _degAppKey
	r.Set("deg_app_key", _degAppKey)
	return nil
}

// GetDegAppKey DegAppKey Getter
func (r TaobaodegoperationdoluckydrawAPIRequest) GetDegAppKey() string {
	return r._degAppKey
}

// SetDegEventKey is DegEventKey Setter
// 后台活动配置eventkey
func (r *TaobaodegoperationdoluckydrawAPIRequest) SetDegEventKey(_degEventKey string) error {
	r._degEventKey = _degEventKey
	r.Set("deg_event_key", _degEventKey)
	return nil
}

// GetDegEventKey DegEventKey Getter
func (r TaobaodegoperationdoluckydrawAPIRequest) GetDegEventKey() string {
	return r._degEventKey
}

// SetDegAccessToken is DegAccessToken Setter
// 传参信息
func (r *TaobaodegoperationdoluckydrawAPIRequest) SetDegAccessToken(_degAccessToken string) error {
	r._degAccessToken = _degAccessToken
	r.Set("deg_access_token", _degAccessToken)
	return nil
}

// GetDegAccessToken DegAccessToken Getter
func (r TaobaodegoperationdoluckydrawAPIRequest) GetDegAccessToken() string {
	return r._degAccessToken
}

// SetSource is Source Setter
// 前端标识
func (r *TaobaodegoperationdoluckydrawAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaodegoperationdoluckydrawAPIRequest) GetSource() string {
	return r._source
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *TaobaodegoperationdoluckydrawAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaodegoperationdoluckydrawAPIRequest) GetUuid() string {
	return r._uuid
}

// SetParamSign is ParamSign Setter
// 参数校验
func (r *TaobaodegoperationdoluckydrawAPIRequest) SetParamSign(_paramSign string) error {
	r._paramSign = _paramSign
	r.Set("param_sign", _paramSign)
	return nil
}

// GetParamSign ParamSign Getter
func (r TaobaodegoperationdoluckydrawAPIRequest) GetParamSign() string {
	return r._paramSign
}
