package alisports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsDataSportsSyncstatdataAPIRequest 阿里体育数据中心用户当天累积数据同步接口 API请求
// alibaba.alisports.data.sports.syncstatdata
//
// 阿里体育数据中心用户当天累积数据同步接口
type AlibabaAlisportsDataSportsSyncstatdataAPIRequest struct {
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

// NewAlibabaAlisportsDataSportsSyncstatdataRequest 初始化AlibabaAlisportsDataSportsSyncstatdataAPIRequest对象
func NewAlibabaAlisportsDataSportsSyncstatdataRequest() *AlibabaAlisportsDataSportsSyncstatdataAPIRequest {
	return &AlibabaAlisportsDataSportsSyncstatdataAPIRequest{
		Params: model.NewParams(11),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) Reset() {
	r._alispAppKey = ""
	r._alispTime = ""
	r._alispSign = ""
	r._aliuid = ""
	r._steps = ""
	r._calorie = ""
	r._distance = ""
	r._time = ""
	r._deviceType = ""
	r._deviceName = ""
	r._deviceModel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.data.sports.syncstatdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 应用appkey
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispTime is AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispSign is AlispSign Setter
// 签名
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetAliuid is Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// GetAliuid Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetAliuid() string {
	return r._aliuid
}

// SetSteps is Steps Setter
// 运动步数
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) SetSteps(_steps string) error {
	r._steps = _steps
	r.Set("steps", _steps)
	return nil
}

// GetSteps Steps Getter
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetSteps() string {
	return r._steps
}

// SetCalorie is Calorie Setter
// 消耗卡路里 单位：卡
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) SetCalorie(_calorie string) error {
	r._calorie = _calorie
	r.Set("calorie", _calorie)
	return nil
}

// GetCalorie Calorie Getter
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetCalorie() string {
	return r._calorie
}

// SetDistance is Distance Setter
// 运动距离 单位：米
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) SetDistance(_distance string) error {
	r._distance = _distance
	r.Set("distance", _distance)
	return nil
}

// GetDistance Distance Getter
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetDistance() string {
	return r._distance
}

// SetTime is Time Setter
// 日期 格式：y-m-d h:i:s
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) SetTime(_time string) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// GetTime Time Getter
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetTime() string {
	return r._time
}

// SetDeviceType is DeviceType Setter
// 设备类型：1手环
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetDeviceName is DeviceName Setter
// 设备名
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) SetDeviceName(_deviceName string) error {
	r._deviceName = _deviceName
	r.Set("device_name", _deviceName)
	return nil
}

// GetDeviceName DeviceName Getter
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetDeviceName() string {
	return r._deviceName
}

// SetDeviceModel is DeviceModel Setter
// 设备型号
func (r *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// GetDeviceModel DeviceModel Getter
func (r AlibabaAlisportsDataSportsSyncstatdataAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

var poolAlibabaAlisportsDataSportsSyncstatdataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlisportsDataSportsSyncstatdataRequest()
	},
}

// GetAlibabaAlisportsDataSportsSyncstatdataRequest 从 sync.Pool 获取 AlibabaAlisportsDataSportsSyncstatdataAPIRequest
func GetAlibabaAlisportsDataSportsSyncstatdataAPIRequest() *AlibabaAlisportsDataSportsSyncstatdataAPIRequest {
	return poolAlibabaAlisportsDataSportsSyncstatdataAPIRequest.Get().(*AlibabaAlisportsDataSportsSyncstatdataAPIRequest)
}

// ReleaseAlibabaAlisportsDataSportsSyncstatdataAPIRequest 将 AlibabaAlisportsDataSportsSyncstatdataAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlisportsDataSportsSyncstatdataAPIRequest(v *AlibabaAlisportsDataSportsSyncstatdataAPIRequest) {
	v.Reset()
	poolAlibabaAlisportsDataSportsSyncstatdataAPIRequest.Put(v)
}
