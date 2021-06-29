package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育同步跑步机设备数据 API请求
alibaba.alisports.datacenter.datasync.treadmill

合作方向阿里体育同步跑步机设备的数据
*/
type AlibabaAlisportsDatacenterDatasyncTreadmillRequest struct {
    model.Params
    // 阿里体育用户id
    userId   string
    // 运动开始时间,秒级别时间戳
    startTime   int64
    // 运动结束时间，秒级别时间戳
    endTime   int64
    // 时区编码，不传默认东八区
    timezone   string
    // 运动位置经纬度
    location   string
    // 国家编码，https://zh.wikipedia.org/wiki/%E5%9C%8B%E5%AE%B6%E5%9C%B0%E5%8D%80%E4%BB%A3%E7%A2%BC
    countryCode   int64
    // 省编码， https://www.ip33.com/area_code.html
    provinceCode   int64
    // 城市编码
    cityCode   int64
    // 平均速度，单位km/h
    speed   *BigDecimal
    // 最大速度，单位km/h
    maxSpeed   *BigDecimal
    // 平均步频
    powerFrequency   int64
    // 累计里程，单位：m
    mileage   int64
    // 累计爬升，单位m
    climb   int64
    // 运动持续时常,单位：秒
    durationTime   int64
    // 总步数
    steps   int64
    // 消耗总热量，单位：卡路里
    calorie   *BigDecimal
    // 最大心率
    maxHeartrate   int64
    // 最小心率
    minHeartrate   int64
    // 平均心率
    avgHeartrate   int64
    // 过程数据收集间隔时间
    collectTimeInterval   int64
    // 过程数据收集间隔时间单位，1.毫秒 2.秒 3.分 4.小时
    collectTimeUnit   int64
    // 速度过程数据,单位km/h
    speedDatas   []string
    // 步频/踏频/桨频过程数据
    motionFrequencyDatas   []string
    // 步幅/踏幅/桨幅过程数据
    hrzMotionRangeDatas   []string
    // 配速过程数据
    tempoDatas   []string
    // 心率过程数据
    heartrateDatas   []int64
    // 三方数据唯一id
    messageId   string
    // 设备类型：4.跑步机 5.单车 6.划船机
    deviceType   int64
    // 设备名称
    deviceName   string
    // 设备型号
    deviceModel   string
}

// 初始化AlibabaAlisportsDatacenterDatasyncTreadmillRequest对象
func NewAlibabaAlisportsDatacenterDatasyncTreadmillRequest() *AlibabaAlisportsDatacenterDatasyncTreadmillRequest{
    return &AlibabaAlisportsDatacenterDatasyncTreadmillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetApiMethodName() string {
    return "alibaba.alisports.datacenter.datasync.treadmill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetUserId() string {
    return r.userId
}
// StartTime Setter
// 运动开始时间,秒级别时间戳
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetStartTime(startTime int64) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetStartTime() int64 {
    return r.startTime
}
// EndTime Setter
// 运动结束时间，秒级别时间戳
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetEndTime(endTime int64) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetEndTime() int64 {
    return r.endTime
}
// Timezone Setter
// 时区编码，不传默认东八区
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetTimezone(timezone string) error {
    r.timezone = timezone
    r.Set("timezone", timezone)
    return nil
}

// Timezone Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetTimezone() string {
    return r.timezone
}
// Location Setter
// 运动位置经纬度
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetLocation(location string) error {
    r.location = location
    r.Set("location", location)
    return nil
}

// Location Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetLocation() string {
    return r.location
}
// CountryCode Setter
// 国家编码，https://zh.wikipedia.org/wiki/%E5%9C%8B%E5%AE%B6%E5%9C%B0%E5%8D%80%E4%BB%A3%E7%A2%BC
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetCountryCode(countryCode int64) error {
    r.countryCode = countryCode
    r.Set("country_code", countryCode)
    return nil
}

// CountryCode Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetCountryCode() int64 {
    return r.countryCode
}
// ProvinceCode Setter
// 省编码， https://www.ip33.com/area_code.html
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetProvinceCode(provinceCode int64) error {
    r.provinceCode = provinceCode
    r.Set("province_code", provinceCode)
    return nil
}

// ProvinceCode Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetProvinceCode() int64 {
    return r.provinceCode
}
// CityCode Setter
// 城市编码
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetCityCode(cityCode int64) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetCityCode() int64 {
    return r.cityCode
}
// Speed Setter
// 平均速度，单位km/h
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetSpeed(speed *BigDecimal) error {
    r.speed = speed
    r.Set("speed", speed)
    return nil
}

// Speed Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetSpeed() *BigDecimal {
    return r.speed
}
// MaxSpeed Setter
// 最大速度，单位km/h
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMaxSpeed(maxSpeed *BigDecimal) error {
    r.maxSpeed = maxSpeed
    r.Set("max_speed", maxSpeed)
    return nil
}

// MaxSpeed Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMaxSpeed() *BigDecimal {
    return r.maxSpeed
}
// PowerFrequency Setter
// 平均步频
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetPowerFrequency(powerFrequency int64) error {
    r.powerFrequency = powerFrequency
    r.Set("power_frequency", powerFrequency)
    return nil
}

// PowerFrequency Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetPowerFrequency() int64 {
    return r.powerFrequency
}
// Mileage Setter
// 累计里程，单位：m
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMileage(mileage int64) error {
    r.mileage = mileage
    r.Set("mileage", mileage)
    return nil
}

// Mileage Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMileage() int64 {
    return r.mileage
}
// Climb Setter
// 累计爬升，单位m
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetClimb(climb int64) error {
    r.climb = climb
    r.Set("climb", climb)
    return nil
}

// Climb Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetClimb() int64 {
    return r.climb
}
// DurationTime Setter
// 运动持续时常,单位：秒
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetDurationTime(durationTime int64) error {
    r.durationTime = durationTime
    r.Set("duration_time", durationTime)
    return nil
}

// DurationTime Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetDurationTime() int64 {
    return r.durationTime
}
// Steps Setter
// 总步数
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetSteps(steps int64) error {
    r.steps = steps
    r.Set("steps", steps)
    return nil
}

// Steps Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetSteps() int64 {
    return r.steps
}
// Calorie Setter
// 消耗总热量，单位：卡路里
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetCalorie(calorie *BigDecimal) error {
    r.calorie = calorie
    r.Set("calorie", calorie)
    return nil
}

// Calorie Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetCalorie() *BigDecimal {
    return r.calorie
}
// MaxHeartrate Setter
// 最大心率
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMaxHeartrate(maxHeartrate int64) error {
    r.maxHeartrate = maxHeartrate
    r.Set("max_heartrate", maxHeartrate)
    return nil
}

// MaxHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMaxHeartrate() int64 {
    return r.maxHeartrate
}
// MinHeartrate Setter
// 最小心率
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMinHeartrate(minHeartrate int64) error {
    r.minHeartrate = minHeartrate
    r.Set("min_heartrate", minHeartrate)
    return nil
}

// MinHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMinHeartrate() int64 {
    return r.minHeartrate
}
// AvgHeartrate Setter
// 平均心率
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetAvgHeartrate(avgHeartrate int64) error {
    r.avgHeartrate = avgHeartrate
    r.Set("avg_heartrate", avgHeartrate)
    return nil
}

// AvgHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetAvgHeartrate() int64 {
    return r.avgHeartrate
}
// CollectTimeInterval Setter
// 过程数据收集间隔时间
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetCollectTimeInterval(collectTimeInterval int64) error {
    r.collectTimeInterval = collectTimeInterval
    r.Set("collect_time_interval", collectTimeInterval)
    return nil
}

// CollectTimeInterval Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetCollectTimeInterval() int64 {
    return r.collectTimeInterval
}
// CollectTimeUnit Setter
// 过程数据收集间隔时间单位，1.毫秒 2.秒 3.分 4.小时
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetCollectTimeUnit(collectTimeUnit int64) error {
    r.collectTimeUnit = collectTimeUnit
    r.Set("collect_time_unit", collectTimeUnit)
    return nil
}

// CollectTimeUnit Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetCollectTimeUnit() int64 {
    return r.collectTimeUnit
}
// SpeedDatas Setter
// 速度过程数据,单位km/h
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetSpeedDatas(speedDatas []string) error {
    r.speedDatas = speedDatas
    r.Set("speed_datas", speedDatas)
    return nil
}

// SpeedDatas Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetSpeedDatas() []string {
    return r.speedDatas
}
// MotionFrequencyDatas Setter
// 步频/踏频/桨频过程数据
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMotionFrequencyDatas(motionFrequencyDatas []string) error {
    r.motionFrequencyDatas = motionFrequencyDatas
    r.Set("motion_frequency_datas", motionFrequencyDatas)
    return nil
}

// MotionFrequencyDatas Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMotionFrequencyDatas() []string {
    return r.motionFrequencyDatas
}
// HrzMotionRangeDatas Setter
// 步幅/踏幅/桨幅过程数据
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetHrzMotionRangeDatas(hrzMotionRangeDatas []string) error {
    r.hrzMotionRangeDatas = hrzMotionRangeDatas
    r.Set("hrz_motion_range_datas", hrzMotionRangeDatas)
    return nil
}

// HrzMotionRangeDatas Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetHrzMotionRangeDatas() []string {
    return r.hrzMotionRangeDatas
}
// TempoDatas Setter
// 配速过程数据
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetTempoDatas(tempoDatas []string) error {
    r.tempoDatas = tempoDatas
    r.Set("tempo_datas", tempoDatas)
    return nil
}

// TempoDatas Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetTempoDatas() []string {
    return r.tempoDatas
}
// HeartrateDatas Setter
// 心率过程数据
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetHeartrateDatas(heartrateDatas []int64) error {
    r.heartrateDatas = heartrateDatas
    r.Set("heartrate_datas", heartrateDatas)
    return nil
}

// HeartrateDatas Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetHeartrateDatas() []int64 {
    return r.heartrateDatas
}
// MessageId Setter
// 三方数据唯一id
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMessageId(messageId string) error {
    r.messageId = messageId
    r.Set("message_id", messageId)
    return nil
}

// MessageId Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMessageId() string {
    return r.messageId
}
// DeviceType Setter
// 设备类型：4.跑步机 5.单车 6.划船机
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetDeviceType() int64 {
    return r.deviceType
}
// DeviceName Setter
// 设备名称
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetDeviceName(deviceName string) error {
    r.deviceName = deviceName
    r.Set("device_name", deviceName)
    return nil
}

// DeviceName Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetDeviceName() string {
    return r.deviceName
}
// DeviceModel Setter
// 设备型号
func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

// DeviceModel Getter
func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetDeviceModel() string {
    return r.deviceModel
}
