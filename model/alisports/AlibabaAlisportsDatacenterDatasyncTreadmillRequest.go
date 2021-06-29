package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育同步跑步机设备数据 APIRequest
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

func NewAlibabaAlisportsDatacenterDatasyncTreadmillRequest() *AlibabaAlisportsDatacenterDatasyncTreadmillRequest{
    return &AlibabaAlisportsDatacenterDatasyncTreadmillRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetApiMethodName() string {
    return "alibaba.alisports.datacenter.datasync.treadmill"
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetStartTime(startTime int64) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetStartTime() int64 {
    return r.startTime
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetEndTime(endTime int64) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetEndTime() int64 {
    return r.endTime
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetTimezone(timezone string) error {
    r.timezone = timezone
    r.Set("timezone", timezone)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetTimezone() string {
    return r.timezone
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetLocation(location string) error {
    r.location = location
    r.Set("location", location)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetLocation() string {
    return r.location
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetCountryCode(countryCode int64) error {
    r.countryCode = countryCode
    r.Set("country_code", countryCode)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetCountryCode() int64 {
    return r.countryCode
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetProvinceCode(provinceCode int64) error {
    r.provinceCode = provinceCode
    r.Set("province_code", provinceCode)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetProvinceCode() int64 {
    return r.provinceCode
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetCityCode(cityCode int64) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetCityCode() int64 {
    return r.cityCode
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetSpeed(speed *BigDecimal) error {
    r.speed = speed
    r.Set("speed", speed)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetSpeed() *BigDecimal {
    return r.speed
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMaxSpeed(maxSpeed *BigDecimal) error {
    r.maxSpeed = maxSpeed
    r.Set("max_speed", maxSpeed)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMaxSpeed() *BigDecimal {
    return r.maxSpeed
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetPowerFrequency(powerFrequency int64) error {
    r.powerFrequency = powerFrequency
    r.Set("power_frequency", powerFrequency)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetPowerFrequency() int64 {
    return r.powerFrequency
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMileage(mileage int64) error {
    r.mileage = mileage
    r.Set("mileage", mileage)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMileage() int64 {
    return r.mileage
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetClimb(climb int64) error {
    r.climb = climb
    r.Set("climb", climb)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetClimb() int64 {
    return r.climb
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetDurationTime(durationTime int64) error {
    r.durationTime = durationTime
    r.Set("duration_time", durationTime)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetDurationTime() int64 {
    return r.durationTime
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetSteps(steps int64) error {
    r.steps = steps
    r.Set("steps", steps)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetSteps() int64 {
    return r.steps
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetCalorie(calorie *BigDecimal) error {
    r.calorie = calorie
    r.Set("calorie", calorie)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetCalorie() *BigDecimal {
    return r.calorie
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMaxHeartrate(maxHeartrate int64) error {
    r.maxHeartrate = maxHeartrate
    r.Set("max_heartrate", maxHeartrate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMaxHeartrate() int64 {
    return r.maxHeartrate
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMinHeartrate(minHeartrate int64) error {
    r.minHeartrate = minHeartrate
    r.Set("min_heartrate", minHeartrate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMinHeartrate() int64 {
    return r.minHeartrate
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetAvgHeartrate(avgHeartrate int64) error {
    r.avgHeartrate = avgHeartrate
    r.Set("avg_heartrate", avgHeartrate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetAvgHeartrate() int64 {
    return r.avgHeartrate
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetCollectTimeInterval(collectTimeInterval int64) error {
    r.collectTimeInterval = collectTimeInterval
    r.Set("collect_time_interval", collectTimeInterval)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetCollectTimeInterval() int64 {
    return r.collectTimeInterval
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetCollectTimeUnit(collectTimeUnit int64) error {
    r.collectTimeUnit = collectTimeUnit
    r.Set("collect_time_unit", collectTimeUnit)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetCollectTimeUnit() int64 {
    return r.collectTimeUnit
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetSpeedDatas(speedDatas []string) error {
    r.speedDatas = speedDatas
    r.Set("speed_datas", speedDatas)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetSpeedDatas() []string {
    return r.speedDatas
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMotionFrequencyDatas(motionFrequencyDatas []string) error {
    r.motionFrequencyDatas = motionFrequencyDatas
    r.Set("motion_frequency_datas", motionFrequencyDatas)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMotionFrequencyDatas() []string {
    return r.motionFrequencyDatas
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetHrzMotionRangeDatas(hrzMotionRangeDatas []string) error {
    r.hrzMotionRangeDatas = hrzMotionRangeDatas
    r.Set("hrz_motion_range_datas", hrzMotionRangeDatas)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetHrzMotionRangeDatas() []string {
    return r.hrzMotionRangeDatas
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetTempoDatas(tempoDatas []string) error {
    r.tempoDatas = tempoDatas
    r.Set("tempo_datas", tempoDatas)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetTempoDatas() []string {
    return r.tempoDatas
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetHeartrateDatas(heartrateDatas []int64) error {
    r.heartrateDatas = heartrateDatas
    r.Set("heartrate_datas", heartrateDatas)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetHeartrateDatas() []int64 {
    return r.heartrateDatas
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetMessageId(messageId string) error {
    r.messageId = messageId
    r.Set("message_id", messageId)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetMessageId() string {
    return r.messageId
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetDeviceType() int64 {
    return r.deviceType
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetDeviceName(deviceName string) error {
    r.deviceName = deviceName
    r.Set("device_name", deviceName)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetDeviceName() string {
    return r.deviceName
}

func (r *AlibabaAlisportsDatacenterDatasyncTreadmillRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncTreadmillRequest) GetDeviceModel() string {
    return r.deviceModel
}

