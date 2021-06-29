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
    _deviceToken   string
    // 设备imei
    _imei   string
    // 设备uuid
    _uuid   string
    // 消息ID
    _mid   int64
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
func (r *YunosServiceCmnsCoaMessageAckRequest) SetDeviceToken(_deviceToken string) error {
    r._deviceToken = _deviceToken
    r.Set("device_token", _deviceToken)
    return nil
}

// DeviceToken Getter
func (r YunosServiceCmnsCoaMessageAckRequest) GetDeviceToken() string {
    return r._deviceToken
}
// Imei Setter
// 设备imei
func (r *YunosServiceCmnsCoaMessageAckRequest) SetImei(_imei string) error {
    r._imei = _imei
    r.Set("imei", _imei)
    return nil
}

// Imei Getter
func (r YunosServiceCmnsCoaMessageAckRequest) GetImei() string {
    return r._imei
}
// Uuid Setter
// 设备uuid
func (r *YunosServiceCmnsCoaMessageAckRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r YunosServiceCmnsCoaMessageAckRequest) GetUuid() string {
    return r._uuid
}
// Mid Setter
// 消息ID
func (r *YunosServiceCmnsCoaMessageAckRequest) SetMid(_mid int64) error {
    r._mid = _mid
    r.Set("mid", _mid)
    return nil
}

// Mid Getter
func (r YunosServiceCmnsCoaMessageAckRequest) GetMid() int64 {
    return r._mid
}
