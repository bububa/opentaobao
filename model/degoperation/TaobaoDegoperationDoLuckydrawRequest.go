package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
激励抽奖 API请求
taobao.degoperation.do.luckydraw

激励平台抽奖接口。用户可以通过接口完成抽奖功能
*/
type TaobaoDegoperationDoLuckydrawRequest struct {
    model.Params
    // 后台活动配置appkey
    _degAppKey   string
    // 后台活动配置eventkey
    _degEventKey   string
    // 前端标识
    _source   string
    // 设备uuid
    _uuid   string
    // 参数校验
    _paramSign   string
    // 传参信息
    _degAccessToken   string
}

// 初始化TaobaoDegoperationDoLuckydrawRequest对象
func NewTaobaoDegoperationDoLuckydrawRequest() *TaobaoDegoperationDoLuckydrawRequest{
    return &TaobaoDegoperationDoLuckydrawRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationDoLuckydrawRequest) GetApiMethodName() string {
    return "taobao.degoperation.do.luckydraw"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationDoLuckydrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DegAppKey Setter
// 后台活动配置appkey
func (r *TaobaoDegoperationDoLuckydrawRequest) SetDegAppKey(_degAppKey string) error {
    r._degAppKey = _degAppKey
    r.Set("deg_app_key", _degAppKey)
    return nil
}

// DegAppKey Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetDegAppKey() string {
    return r._degAppKey
}
// DegEventKey Setter
// 后台活动配置eventkey
func (r *TaobaoDegoperationDoLuckydrawRequest) SetDegEventKey(_degEventKey string) error {
    r._degEventKey = _degEventKey
    r.Set("deg_event_key", _degEventKey)
    return nil
}

// DegEventKey Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetDegEventKey() string {
    return r._degEventKey
}
// Source Setter
// 前端标识
func (r *TaobaoDegoperationDoLuckydrawRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetSource() string {
    return r._source
}
// Uuid Setter
// 设备uuid
func (r *TaobaoDegoperationDoLuckydrawRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetUuid() string {
    return r._uuid
}
// ParamSign Setter
// 参数校验
func (r *TaobaoDegoperationDoLuckydrawRequest) SetParamSign(_paramSign string) error {
    r._paramSign = _paramSign
    r.Set("param_sign", _paramSign)
    return nil
}

// ParamSign Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetParamSign() string {
    return r._paramSign
}
// DegAccessToken Setter
// 传参信息
func (r *TaobaoDegoperationDoLuckydrawRequest) SetDegAccessToken(_degAccessToken string) error {
    r._degAccessToken = _degAccessToken
    r.Set("deg_access_token", _degAccessToken)
    return nil
}

// DegAccessToken Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetDegAccessToken() string {
    return r._degAccessToken
}
