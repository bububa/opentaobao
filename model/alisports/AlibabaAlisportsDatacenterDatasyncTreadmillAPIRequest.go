package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportsdatacenterdatasynctreadmillAPIRequest 阿里体育同步跑步机设备数据 API请求
// alibaba.alisports.datacenter.datasync.treadmill
//
// 合作方向阿里体育同步跑步机设备的数据
type AlibabaalisportsdatacenterdatasynctreadmillAPIRequest struct {
	model.Params
	// 配速过程数据
	_tempoDatas []int64
	// 心率过程数据
	_heartrateDatas []int64
	// 步幅/踏幅/桨幅过程数据
	_hrzMotionRangeDatas []int64
	// 步频/踏频/桨频过程数据
	_motionFrequencyDatas []int64
	// 速度过程数据,单位km/h
	_speedDatas []int64
	// 时区编码，不传默认东八区
	_timezone string
	// 设备名称
	_deviceName string
	// 三方数据唯一id
	_messageId string
	// 阿里体育用户id
	_userId string
	// 运动位置经纬度
	_location string
	// 设备型号
	_deviceModel string
	// 城市编码
	_cityCode int64
	// 最小心率
	_minHeartrate int64
	// 最大心率
	_maxHeartrate int64
	// 平均心率
	_avgHeartrate int64
	// 平均速度，单位km/h
	_speed float64
	// 过程数据收集间隔时间
	_collectTimeInterval int64
	// 国家编码，https://zh.wikipedia.org/wiki/%E5%9C%8B%E5%AE%B6%E5%9C%B0%E5%8D%80%E4%BB%A3%E7%A2%BC
	_countryCode int64
	// 运动结束时间，秒级别时间戳
	_endTime int64
	// 运动开始时间,秒级别时间戳
	_startTime int64
	// 累计里程，单位：m
	_mileage int64
	// 累计爬升，单位m
	_climb int64
	// 设备类型：4.跑步机 5.单车 6.划船机
	_deviceType int64
	// 运动持续时常,单位：秒
	_durationTime int64
	// 省编码， https://www.ip33.com/area_code.html
	_provinceCode int64
	// 过程数据收集间隔时间单位，1.毫秒 2.秒 3.分 4.小时
	_collectTimeUnit int64
	// 最大速度，单位km/h
	_maxSpeed float64
	// 平均步频
	_powerFrequency int64
	// 消耗总热量，单位：卡路里
	_calorie float64
	// 总步数
	_steps int64
}

// NewAlibabaalisportsdatacenterdatasynctreadmillRequest 初始化AlibabaalisportsdatacenterdatasynctreadmillAPIRequest对象
func NewAlibabaalisportsdatacenterdatasynctreadmillRequest() *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest {
	return &AlibabaalisportsdatacenterdatasynctreadmillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.datacenter.datasync.treadmill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTempoDatas is TempoDatas Setter
// 配速过程数据
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetTempoDatas(_tempoDatas []int64) error {
	r._tempoDatas = _tempoDatas
	r.Set("tempo_datas", _tempoDatas)
	return nil
}

// GetTempoDatas TempoDatas Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetTempoDatas() []int64 {
	return r._tempoDatas
}

// SetHeartrateDatas is HeartrateDatas Setter
// 心率过程数据
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetHeartrateDatas(_heartrateDatas []int64) error {
	r._heartrateDatas = _heartrateDatas
	r.Set("heartrate_datas", _heartrateDatas)
	return nil
}

// GetHeartrateDatas HeartrateDatas Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetHeartrateDatas() []int64 {
	return r._heartrateDatas
}

// SetHrzMotionRangeDatas is HrzMotionRangeDatas Setter
// 步幅/踏幅/桨幅过程数据
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetHrzMotionRangeDatas(_hrzMotionRangeDatas []int64) error {
	r._hrzMotionRangeDatas = _hrzMotionRangeDatas
	r.Set("hrz_motion_range_datas", _hrzMotionRangeDatas)
	return nil
}

// GetHrzMotionRangeDatas HrzMotionRangeDatas Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetHrzMotionRangeDatas() []int64 {
	return r._hrzMotionRangeDatas
}

// SetMotionFrequencyDatas is MotionFrequencyDatas Setter
// 步频/踏频/桨频过程数据
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetMotionFrequencyDatas(_motionFrequencyDatas []int64) error {
	r._motionFrequencyDatas = _motionFrequencyDatas
	r.Set("motion_frequency_datas", _motionFrequencyDatas)
	return nil
}

// GetMotionFrequencyDatas MotionFrequencyDatas Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetMotionFrequencyDatas() []int64 {
	return r._motionFrequencyDatas
}

// SetSpeedDatas is SpeedDatas Setter
// 速度过程数据,单位km/h
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetSpeedDatas(_speedDatas []int64) error {
	r._speedDatas = _speedDatas
	r.Set("speed_datas", _speedDatas)
	return nil
}

// GetSpeedDatas SpeedDatas Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetSpeedDatas() []int64 {
	return r._speedDatas
}

// SetTimezone is Timezone Setter
// 时区编码，不传默认东八区
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetTimezone(_timezone string) error {
	r._timezone = _timezone
	r.Set("timezone", _timezone)
	return nil
}

// GetTimezone Timezone Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetTimezone() string {
	return r._timezone
}

// SetDeviceName is DeviceName Setter
// 设备名称
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetDeviceName(_deviceName string) error {
	r._deviceName = _deviceName
	r.Set("device_name", _deviceName)
	return nil
}

// GetDeviceName DeviceName Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetDeviceName() string {
	return r._deviceName
}

// SetMessageId is MessageId Setter
// 三方数据唯一id
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetMessageId(_messageId string) error {
	r._messageId = _messageId
	r.Set("message_id", _messageId)
	return nil
}

// GetMessageId MessageId Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetMessageId() string {
	return r._messageId
}

// SetUserId is UserId Setter
// 阿里体育用户id
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetUserId() string {
	return r._userId
}

// SetLocation is Location Setter
// 运动位置经纬度
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetLocation(_location string) error {
	r._location = _location
	r.Set("location", _location)
	return nil
}

// GetLocation Location Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetLocation() string {
	return r._location
}

// SetDeviceModel is DeviceModel Setter
// 设备型号
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// GetDeviceModel DeviceModel Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

// SetCityCode is CityCode Setter
// 城市编码
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetCityCode() int64 {
	return r._cityCode
}

// SetMinHeartrate is MinHeartrate Setter
// 最小心率
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetMinHeartrate(_minHeartrate int64) error {
	r._minHeartrate = _minHeartrate
	r.Set("min_heartrate", _minHeartrate)
	return nil
}

// GetMinHeartrate MinHeartrate Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetMinHeartrate() int64 {
	return r._minHeartrate
}

// SetMaxHeartrate is MaxHeartrate Setter
// 最大心率
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetMaxHeartrate(_maxHeartrate int64) error {
	r._maxHeartrate = _maxHeartrate
	r.Set("max_heartrate", _maxHeartrate)
	return nil
}

// GetMaxHeartrate MaxHeartrate Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetMaxHeartrate() int64 {
	return r._maxHeartrate
}

// SetAvgHeartrate is AvgHeartrate Setter
// 平均心率
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetAvgHeartrate(_avgHeartrate int64) error {
	r._avgHeartrate = _avgHeartrate
	r.Set("avg_heartrate", _avgHeartrate)
	return nil
}

// GetAvgHeartrate AvgHeartrate Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetAvgHeartrate() int64 {
	return r._avgHeartrate
}

// SetSpeed is Speed Setter
// 平均速度，单位km/h
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetSpeed(_speed float64) error {
	r._speed = _speed
	r.Set("speed", _speed)
	return nil
}

// GetSpeed Speed Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetSpeed() float64 {
	return r._speed
}

// SetCollectTimeInterval is CollectTimeInterval Setter
// 过程数据收集间隔时间
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetCollectTimeInterval(_collectTimeInterval int64) error {
	r._collectTimeInterval = _collectTimeInterval
	r.Set("collect_time_interval", _collectTimeInterval)
	return nil
}

// GetCollectTimeInterval CollectTimeInterval Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetCollectTimeInterval() int64 {
	return r._collectTimeInterval
}

// SetCountryCode is CountryCode Setter
// 国家编码，https://zh.wikipedia.org/wiki/%E5%9C%8B%E5%AE%B6%E5%9C%B0%E5%8D%80%E4%BB%A3%E7%A2%BC
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetCountryCode(_countryCode int64) error {
	r._countryCode = _countryCode
	r.Set("country_code", _countryCode)
	return nil
}

// GetCountryCode CountryCode Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetCountryCode() int64 {
	return r._countryCode
}

// SetEndTime is EndTime Setter
// 运动结束时间，秒级别时间戳
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetEndTime(_endTime int64) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetEndTime() int64 {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 运动开始时间,秒级别时间戳
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetStartTime(_startTime int64) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetStartTime() int64 {
	return r._startTime
}

// SetMileage is Mileage Setter
// 累计里程，单位：m
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetMileage(_mileage int64) error {
	r._mileage = _mileage
	r.Set("mileage", _mileage)
	return nil
}

// GetMileage Mileage Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetMileage() int64 {
	return r._mileage
}

// SetClimb is Climb Setter
// 累计爬升，单位m
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetClimb(_climb int64) error {
	r._climb = _climb
	r.Set("climb", _climb)
	return nil
}

// GetClimb Climb Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetClimb() int64 {
	return r._climb
}

// SetDeviceType is DeviceType Setter
// 设备类型：4.跑步机 5.单车 6.划船机
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetDeviceType() int64 {
	return r._deviceType
}

// SetDurationTime is DurationTime Setter
// 运动持续时常,单位：秒
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetDurationTime(_durationTime int64) error {
	r._durationTime = _durationTime
	r.Set("duration_time", _durationTime)
	return nil
}

// GetDurationTime DurationTime Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetDurationTime() int64 {
	return r._durationTime
}

// SetProvinceCode is ProvinceCode Setter
// 省编码， https://www.ip33.com/area_code.html
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetProvinceCode(_provinceCode int64) error {
	r._provinceCode = _provinceCode
	r.Set("province_code", _provinceCode)
	return nil
}

// GetProvinceCode ProvinceCode Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetProvinceCode() int64 {
	return r._provinceCode
}

// SetCollectTimeUnit is CollectTimeUnit Setter
// 过程数据收集间隔时间单位，1.毫秒 2.秒 3.分 4.小时
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetCollectTimeUnit(_collectTimeUnit int64) error {
	r._collectTimeUnit = _collectTimeUnit
	r.Set("collect_time_unit", _collectTimeUnit)
	return nil
}

// GetCollectTimeUnit CollectTimeUnit Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetCollectTimeUnit() int64 {
	return r._collectTimeUnit
}

// SetMaxSpeed is MaxSpeed Setter
// 最大速度，单位km/h
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetMaxSpeed(_maxSpeed float64) error {
	r._maxSpeed = _maxSpeed
	r.Set("max_speed", _maxSpeed)
	return nil
}

// GetMaxSpeed MaxSpeed Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetMaxSpeed() float64 {
	return r._maxSpeed
}

// SetPowerFrequency is PowerFrequency Setter
// 平均步频
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetPowerFrequency(_powerFrequency int64) error {
	r._powerFrequency = _powerFrequency
	r.Set("power_frequency", _powerFrequency)
	return nil
}

// GetPowerFrequency PowerFrequency Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetPowerFrequency() int64 {
	return r._powerFrequency
}

// SetCalorie is Calorie Setter
// 消耗总热量，单位：卡路里
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetCalorie(_calorie float64) error {
	r._calorie = _calorie
	r.Set("calorie", _calorie)
	return nil
}

// GetCalorie Calorie Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetCalorie() float64 {
	return r._calorie
}

// SetSteps is Steps Setter
// 总步数
func (r *AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) SetSteps(_steps int64) error {
	r._steps = _steps
	r.Set("steps", _steps)
	return nil
}

// GetSteps Steps Getter
func (r AlibabaalisportsdatacenterdatasynctreadmillAPIRequest) GetSteps() int64 {
	return r._steps
}
