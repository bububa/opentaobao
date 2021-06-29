package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息回执查询 API请求
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

// 初始化YunosServiceCmnsCoaMessageAckRequest对象
func NewYunosServiceCmnsCoaMessageAckRequest() *YunosServiceCmnsCoaMessageAckRequest{
    return &YunosServiceCmnsCoaMessageAckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageAckRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.message.ack"
}

// IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageAckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceToken Setter
// 设备唯一值deviceToken
func (r *YunosServiceCmnsCoaMessageAckRequest) SetDeviceToken(deviceToken string) error {
    r.deviceToken = deviceToken
    r.Set("device_token", deviceToken)
    return nil
}

// DeviceToken Getter
func (r YunosServiceCmnsCoaMessageAckRequest) GetDeviceToken() string {
    return r.deviceToken
}
// Imei Setter
// 设备imei
func (r *YunosServiceCmnsCoaMessageAckRequest) SetImei(imei string) error {
    r.imei = imei
    r.Set("imei", imei)
    return nil
}

// Imei Getter
func (r YunosServiceCmnsCoaMessageAckRequest) GetImei() string {
    return r.imei
}
// Uuid Setter
// 设备uuid
func (r *YunosServiceCmnsCoaMessageAckRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r YunosServiceCmnsCoaMessageAckRequest) GetUuid() string {
    return r.uuid
}
// Mid Setter
// 消息ID
func (r *YunosServiceCmnsCoaMessageAckRequest) SetMid(mid int64) error {
    r.mid = mid
    r.Set("mid", mid)
    return nil
}

// Mid Getter
func (r YunosServiceCmnsCoaMessageAckRequest) GetMid() int64 {
    return r.mid
}
