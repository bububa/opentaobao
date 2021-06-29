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
    degAppKey   string
    // 后台活动配置eventkey
    degEventKey   string
    // 前端标识
    source   string
    // 设备uuid
    uuid   string
    // 参数校验
    paramSign   string
    // 传参信息
    degAccessToken   string
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
func (r *TaobaoDegoperationDoLuckydrawRequest) SetDegAppKey(degAppKey string) error {
    r.degAppKey = degAppKey
    r.Set("deg_app_key", degAppKey)
    return nil
}

// DegAppKey Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetDegAppKey() string {
    return r.degAppKey
}
// DegEventKey Setter
// 后台活动配置eventkey
func (r *TaobaoDegoperationDoLuckydrawRequest) SetDegEventKey(degEventKey string) error {
    r.degEventKey = degEventKey
    r.Set("deg_event_key", degEventKey)
    return nil
}

// DegEventKey Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetDegEventKey() string {
    return r.degEventKey
}
// Source Setter
// 前端标识
func (r *TaobaoDegoperationDoLuckydrawRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetSource() string {
    return r.source
}
// Uuid Setter
// 设备uuid
func (r *TaobaoDegoperationDoLuckydrawRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetUuid() string {
    return r.uuid
}
// ParamSign Setter
// 参数校验
func (r *TaobaoDegoperationDoLuckydrawRequest) SetParamSign(paramSign string) error {
    r.paramSign = paramSign
    r.Set("param_sign", paramSign)
    return nil
}

// ParamSign Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetParamSign() string {
    return r.paramSign
}
// DegAccessToken Setter
// 传参信息
func (r *TaobaoDegoperationDoLuckydrawRequest) SetDegAccessToken(degAccessToken string) error {
    r.degAccessToken = degAccessToken
    r.Set("deg_access_token", degAccessToken)
    return nil
}

// DegAccessToken Getter
func (r TaobaoDegoperationDoLuckydrawRequest) GetDegAccessToken() string {
    return r.degAccessToken
}
