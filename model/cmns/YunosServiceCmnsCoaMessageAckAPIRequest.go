package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaMessageAckAPIRequest 消息回执查询 API请求
// yunos.service.cmns.coa.message.ack
//
// 第三方应用开发者调用此接口查询设备是否收到消息，只能查询此appKey床发的消息
type YunosServiceCmnsCoaMessageAckAPIRequest struct {
	model.Params
	// 设备唯一值deviceToken
	_deviceToken string
	// 设备imei
	_imei string
	// 设备uuid
	_uuid string
	// 消息ID
	_mid int64
}

// NewYunosServiceCmnsCoaMessageAckRequest 初始化YunosServiceCmnsCoaMessageAckAPIRequest对象
func NewYunosServiceCmnsCoaMessageAckRequest() *YunosServiceCmnsCoaMessageAckAPIRequest {
	return &YunosServiceCmnsCoaMessageAckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.message.ack"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeviceToken Setter
// 设备唯一值deviceToken
func (r *YunosServiceCmnsCoaMessageAckAPIRequest) SetDeviceToken(_deviceToken string) error {
	r._deviceToken = _deviceToken
	r.Set("device_token", _deviceToken)
	return nil
}

// Get DeviceToken Getter
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetDeviceToken() string {
	return r._deviceToken
}

// Set is Imei Setter
// 设备imei
func (r *YunosServiceCmnsCoaMessageAckAPIRequest) SetImei(_imei string) error {
	r._imei = _imei
	r.Set("imei", _imei)
	return nil
}

// Get Imei Getter
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetImei() string {
	return r._imei
}

// Set is Uuid Setter
// 设备uuid
func (r *YunosServiceCmnsCoaMessageAckAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetUuid() string {
	return r._uuid
}

// Set is Mid Setter
// 消息ID
func (r *YunosServiceCmnsCoaMessageAckAPIRequest) SetMid(_mid int64) error {
	r._mid = _mid
	r.Set("mid", _mid)
	return nil
}

// Get Mid Getter
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetMid() int64 {
	return r._mid
}
