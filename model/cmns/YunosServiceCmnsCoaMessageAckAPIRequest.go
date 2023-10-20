package cmns

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosServiceCmnsCoaMessageAckAPIRequest) Reset() {
	r._deviceToken = ""
	r._imei = ""
	r._uuid = ""
	r._mid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.message.ack"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceToken is DeviceToken Setter
// 设备唯一值deviceToken
func (r *YunosServiceCmnsCoaMessageAckAPIRequest) SetDeviceToken(_deviceToken string) error {
	r._deviceToken = _deviceToken
	r.Set("device_token", _deviceToken)
	return nil
}

// GetDeviceToken DeviceToken Getter
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetDeviceToken() string {
	return r._deviceToken
}

// SetImei is Imei Setter
// 设备imei
func (r *YunosServiceCmnsCoaMessageAckAPIRequest) SetImei(_imei string) error {
	r._imei = _imei
	r.Set("imei", _imei)
	return nil
}

// GetImei Imei Getter
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetImei() string {
	return r._imei
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *YunosServiceCmnsCoaMessageAckAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetUuid() string {
	return r._uuid
}

// SetMid is Mid Setter
// 消息ID
func (r *YunosServiceCmnsCoaMessageAckAPIRequest) SetMid(_mid int64) error {
	r._mid = _mid
	r.Set("mid", _mid)
	return nil
}

// GetMid Mid Getter
func (r YunosServiceCmnsCoaMessageAckAPIRequest) GetMid() int64 {
	return r._mid
}

var poolYunosServiceCmnsCoaMessageAckAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosServiceCmnsCoaMessageAckRequest()
	},
}

// GetYunosServiceCmnsCoaMessageAckRequest 从 sync.Pool 获取 YunosServiceCmnsCoaMessageAckAPIRequest
func GetYunosServiceCmnsCoaMessageAckAPIRequest() *YunosServiceCmnsCoaMessageAckAPIRequest {
	return poolYunosServiceCmnsCoaMessageAckAPIRequest.Get().(*YunosServiceCmnsCoaMessageAckAPIRequest)
}

// ReleaseYunosServiceCmnsCoaMessageAckAPIRequest 将 YunosServiceCmnsCoaMessageAckAPIRequest 放入 sync.Pool
func ReleaseYunosServiceCmnsCoaMessageAckAPIRequest(v *YunosServiceCmnsCoaMessageAckAPIRequest) {
	v.Reset()
	poolYunosServiceCmnsCoaMessageAckAPIRequest.Put(v)
}
