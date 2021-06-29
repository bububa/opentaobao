package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据uuid用户抽奖次数限制 APIRequest
taobao.degoperation.get.info.uuid

根据uuid用户抽奖次数限制
*/
type TaobaoDegoperationGetInfoUuidRequest struct {
    model.Params

    // 活动后台配置eventkey
    degAppKey   string 

    // 活动后台配置appkey
    degEventKey   string 

    // 设备id
    uuid   string 

}

func NewTaobaoDegoperationGetInfoUuidRequest() *TaobaoDegoperationGetInfoUuidRequest{
    return &TaobaoDegoperationGetInfoUuidRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDegoperationGetInfoUuidRequest) GetApiMethodName() string {
    return "taobao.degoperation.get.info.uuid"
}

func (r TaobaoDegoperationGetInfoUuidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDegoperationGetInfoUuidRequest) SetDegAppKey(degAppKey string) error {
    r.degAppKey = degAppKey
    r.Set("deg_app_key", degAppKey)
    return nil
}

func (r TaobaoDegoperationGetInfoUuidRequest) GetDegAppKey() string {
    return r.degAppKey
}

func (r *TaobaoDegoperationGetInfoUuidRequest) SetDegEventKey(degEventKey string) error {
    r.degEventKey = degEventKey
    r.Set("deg_event_key", degEventKey)
    return nil
}

func (r TaobaoDegoperationGetInfoUuidRequest) GetDegEventKey() string {
    return r.degEventKey
}

func (r *TaobaoDegoperationGetInfoUuidRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r TaobaoDegoperationGetInfoUuidRequest) GetUuid() string {
    return r.uuid
}

