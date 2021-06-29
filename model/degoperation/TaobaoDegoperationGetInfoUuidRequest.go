package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据uuid用户抽奖次数限制 API请求
taobao.degoperation.get.info.uuid

根据uuid用户抽奖次数限制
*/
type TaobaoDegoperationGetInfoUuidRequest struct {
    model.Params
    // 活动后台配置eventkey
    _degAppKey   string
    // 活动后台配置appkey
    _degEventKey   string
    // 设备id
    _uuid   string
}

// 初始化TaobaoDegoperationGetInfoUuidRequest对象
func NewTaobaoDegoperationGetInfoUuidRequest() *TaobaoDegoperationGetInfoUuidRequest{
    return &TaobaoDegoperationGetInfoUuidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationGetInfoUuidRequest) GetApiMethodName() string {
    return "taobao.degoperation.get.info.uuid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationGetInfoUuidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DegAppKey Setter
// 活动后台配置eventkey
func (r *TaobaoDegoperationGetInfoUuidRequest) SetDegAppKey(_degAppKey string) error {
    r._degAppKey = _degAppKey
    r.Set("deg_app_key", _degAppKey)
    return nil
}

// DegAppKey Getter
func (r TaobaoDegoperationGetInfoUuidRequest) GetDegAppKey() string {
    return r._degAppKey
}
// DegEventKey Setter
// 活动后台配置appkey
func (r *TaobaoDegoperationGetInfoUuidRequest) SetDegEventKey(_degEventKey string) error {
    r._degEventKey = _degEventKey
    r.Set("deg_event_key", _degEventKey)
    return nil
}

// DegEventKey Getter
func (r TaobaoDegoperationGetInfoUuidRequest) GetDegEventKey() string {
    return r._degEventKey
}
// Uuid Setter
// 设备id
func (r *TaobaoDegoperationGetInfoUuidRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r TaobaoDegoperationGetInfoUuidRequest) GetUuid() string {
    return r._uuid
}
