package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosservicecmnscoamessageackAPIRequest 消息回执查询 API请求
// yunos.service.cmns.coa.message.ack
//
// 第三方应用开发者调用此接口查询设备是否收到消息，只能查询此appKey床发的消息
type YunosservicecmnscoamessageackAPIRequest struct {
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

// NewYunosservicecmnscoamessageackRequest 初始化YunosservicecmnscoamessageackAPIRequest对象
func NewYunosservicecmnscoamessageackRequest() *YunosservicecmnscoamessageackAPIRequest {
	return &YunosservicecmnscoamessageackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosservicecmnscoamessageackAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.message.ack"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosservicecmnscoamessageackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosservicecmnscoamessageackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceToken is DeviceToken Setter
// 设备唯一值deviceToken
func (r *YunosservicecmnscoamessageackAPIRequest) SetDeviceToken(_deviceToken string) error {
	r._deviceToken = _deviceToken
	r.Set("device_token", _deviceToken)
	return nil
}

// GetDeviceToken DeviceToken Getter
func (r YunosservicecmnscoamessageackAPIRequest) GetDeviceToken() string {
	return r._deviceToken
}

// SetImei is Imei Setter
// 设备imei
func (r *YunosservicecmnscoamessageackAPIRequest) SetImei(_imei string) error {
	r._imei = _imei
	r.Set("imei", _imei)
	return nil
}

// GetImei Imei Getter
func (r YunosservicecmnscoamessageackAPIRequest) GetImei() string {
	return r._imei
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *YunosservicecmnscoamessageackAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r YunosservicecmnscoamessageackAPIRequest) GetUuid() string {
	return r._uuid
}

// SetMid is Mid Setter
// 消息ID
func (r *YunosservicecmnscoamessageackAPIRequest) SetMid(_mid int64) error {
	r._mid = _mid
	r.Set("mid", _mid)
	return nil
}

// GetMid Mid Getter
func (r YunosservicecmnscoamessageackAPIRequest) GetMid() int64 {
	return r._mid
}
