package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息回执查询 APIRequest
yunos.service.cmns.coa.message.ack

第三方应用开发者调用此接口查询设备是否收到消息，只能查询此appKey床发的消息
*/
type YunosServiceCmnsCoaMessageAckRequest struct {
    model.Params

    // 设备唯一值deviceToken
    deviceToken   string 

    // 设备imei
    imei   string 

    // 设备uuid
    uuid   string 

    // 消息ID
    mid   int64 

}

func NewYunosServiceCmnsCoaMessageAckRequest() *YunosServiceCmnsCoaMessageAckRequest{
    return &YunosServiceCmnsCoaMessageAckRequest{
        Params: model.NewParams(),
    }
}

func (r YunosServiceCmnsCoaMessageAckRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.message.ack"
}

func (r YunosServiceCmnsCoaMessageAckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosServiceCmnsCoaMessageAckRequest) SetDeviceToken(deviceToken string) error {
    r.deviceToken = deviceToken
    r.Set("device_token", deviceToken)
    return nil
}

func (r YunosServiceCmnsCoaMessageAckRequest) GetDeviceToken() string {
    return r.deviceToken
}

func (r *YunosServiceCmnsCoaMessageAckRequest) SetImei(imei string) error {
    r.imei = imei
    r.Set("imei", imei)
    return nil
}

func (r YunosServiceCmnsCoaMessageAckRequest) GetImei() string {
    return r.imei
}

func (r *YunosServiceCmnsCoaMessageAckRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r YunosServiceCmnsCoaMessageAckRequest) GetUuid() string {
    return r.uuid
}

func (r *YunosServiceCmnsCoaMessageAckRequest) SetMid(mid int64) error {
    r.mid = mid
    r.Set("mid", mid)
    return nil
}

func (r YunosServiceCmnsCoaMessageAckRequest) GetMid() int64 {
    return r.mid
}

