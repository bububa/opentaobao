package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportsdatasportssyncstatdataAPIRequest 阿里体育数据中心用户当天累积数据同步接口 API请求
// alibaba.alisports.data.sports.syncstatdata
//
// 阿里体育数据中心用户当天累积数据同步接口
type AlibabaalisportsdatasportssyncstatdataAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 时间戳精确到秒
	_alispTime string
	// 签名
	_alispSign string
	// 阿里体育用户id
	_aliuid string
	// 运动步数
	_steps string
	// 消耗卡路里 单位：卡
	_calorie string
	// 运动距离 单位：米
	_distance string
	// 日期 格式：y-m-d h:i:s
	_time string
	// 设备类型：1手环
	_deviceType string
	// 设备名
	_deviceName string
	// 设备型号
	_deviceModel string
}

// NewAlibabaalisportsdatasportssyncstatdataRequest 初始化AlibabaalisportsdatasportssyncstatdataAPIRequest对象
func NewAlibabaalisportsdatasportssyncstatdataRequest() *AlibabaalisportsdatasportssyncstatdataAPIRequest {
	return &AlibabaalisportsdatasportssyncstatdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.data.sports.syncstatdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 应用appkey
func (r *AlibabaalisportsdatasportssyncstatdataAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispTime is AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaalisportsdatasportssyncstatdataAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispSign is AlispSign Setter
// 签名
func (r *AlibabaalisportsdatasportssyncstatdataAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetAliuid is Aliuid Setter
// 阿里体育用户id
func (r *AlibabaalisportsdatasportssyncstatdataAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// GetAliuid Aliuid Getter
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetAliuid() string {
	return r._aliuid
}

// SetSteps is Steps Setter
// 运动步数
func (r *AlibabaalisportsdatasportssyncstatdataAPIRequest) SetSteps(_steps string) error {
	r._steps = _steps
	r.Set("steps", _steps)
	return nil
}

// GetSteps Steps Getter
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetSteps() string {
	return r._steps
}

// SetCalorie is Calorie Setter
// 消耗卡路里 单位：卡
func (r *AlibabaalisportsdatasportssyncstatdataAPIRequest) SetCalorie(_calorie string) error {
	r._calorie = _calorie
	r.Set("calorie", _calorie)
	return nil
}

// GetCalorie Calorie Getter
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetCalorie() string {
	return r._calorie
}

// SetDistance is Distance Setter
// 运动距离 单位：米
func (r *AlibabaalisportsdatasportssyncstatdataAPIRequest) SetDistance(_distance string) error {
	r._distance = _distance
	r.Set("distance", _distance)
	return nil
}

// GetDistance Distance Getter
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetDistance() string {
	return r._distance
}

// SetTime is Time Setter
// 日期 格式：y-m-d h:i:s
func (r *AlibabaalisportsdatasportssyncstatdataAPIRequest) SetTime(_time string) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// GetTime Time Getter
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetTime() string {
	return r._time
}

// SetDeviceType is DeviceType Setter
// 设备类型：1手环
func (r *AlibabaalisportsdatasportssyncstatdataAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetDeviceName is DeviceName Setter
// 设备名
func (r *AlibabaalisportsdatasportssyncstatdataAPIRequest) SetDeviceName(_deviceName string) error {
	r._deviceName = _deviceName
	r.Set("device_name", _deviceName)
	return nil
}

// GetDeviceName DeviceName Getter
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetDeviceName() string {
	return r._deviceName
}

// SetDeviceModel is DeviceModel Setter
// 设备型号
func (r *AlibabaalisportsdatasportssyncstatdataAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// GetDeviceModel DeviceModel Getter
func (r AlibabaalisportsdatasportssyncstatdataAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}
