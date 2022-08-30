package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationDoLuckydrawAPIRequest 激励抽奖 API请求
// taobao.degoperation.do.luckydraw
//
// 激励平台抽奖接口。用户可以通过接口完成抽奖功能
type TaobaoDegoperationDoLuckydrawAPIRequest struct {
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

// NewTaobaoDegoperationDoLuckydrawRequest 初始化TaobaoDegoperationDoLuckydrawAPIRequest对象
func NewTaobaoDegoperationDoLuckydrawRequest() *TaobaoDegoperationDoLuckydrawAPIRequest {
	return &TaobaoDegoperationDoLuckydrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationDoLuckydrawAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.do.luckydraw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationDoLuckydrawAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDegAppKey is DegAppKey Setter
// 后台活动配置appkey
func (r *TaobaoDegoperationDoLuckydrawAPIRequest) SetDegAppKey(_degAppKey string) error {
	r._degAppKey = _degAppKey
	r.Set("deg_app_key", _degAppKey)
	return nil
}

// GetDegAppKey DegAppKey Getter
func (r TaobaoDegoperationDoLuckydrawAPIRequest) GetDegAppKey() string {
	return r._degAppKey
}

// SetDegEventKey is DegEventKey Setter
// 后台活动配置eventkey
func (r *TaobaoDegoperationDoLuckydrawAPIRequest) SetDegEventKey(_degEventKey string) error {
	r._degEventKey = _degEventKey
	r.Set("deg_event_key", _degEventKey)
	return nil
}

// GetDegEventKey DegEventKey Getter
func (r TaobaoDegoperationDoLuckydrawAPIRequest) GetDegEventKey() string {
	return r._degEventKey
}

// SetDegAccessToken is DegAccessToken Setter
// 传参信息
func (r *TaobaoDegoperationDoLuckydrawAPIRequest) SetDegAccessToken(_degAccessToken string) error {
	r._degAccessToken = _degAccessToken
	r.Set("deg_access_token", _degAccessToken)
	return nil
}

// GetDegAccessToken DegAccessToken Getter
func (r TaobaoDegoperationDoLuckydrawAPIRequest) GetDegAccessToken() string {
	return r._degAccessToken
}

// SetSource is Source Setter
// 前端标识
func (r *TaobaoDegoperationDoLuckydrawAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoDegoperationDoLuckydrawAPIRequest) GetSource() string {
	return r._source
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *TaobaoDegoperationDoLuckydrawAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaoDegoperationDoLuckydrawAPIRequest) GetUuid() string {
	return r._uuid
}

// SetParamSign is ParamSign Setter
// 参数校验
func (r *TaobaoDegoperationDoLuckydrawAPIRequest) SetParamSign(_paramSign string) error {
	r._paramSign = _paramSign
	r.Set("param_sign", _paramSign)
	return nil
}

// GetParamSign ParamSign Getter
func (r TaobaoDegoperationDoLuckydrawAPIRequest) GetParamSign() string {
	return r._paramSign
}
