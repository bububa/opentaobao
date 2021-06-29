package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育运动数据同步统一接口 API请求
alibaba.alisports.datacenter.datasync.sportsdatas

给单方提供同步运动数据到阿里体育的接口
*/
type AlibabaAlisportsDatacenterDatasyncSportsdatasRequest struct {
    model.Params
    // 用户阿里体育id
    userId   string
    // 运动一级分类
    sportsCat1Id   int64
    // 运动二级分类
    sportsCat2Id   int64
    // 运动三级分类
    sportsCat3Id   string
    // 运动开始时间，单位：毫秒
    sportsStartTime   int64
    // 运动结束时间，单位：毫秒
    sportsEndTime   int64
    // 时区
    timezone   int64
    // 最小心率
    minHeartrate   int64
    // 最大心率
    maxHeartrate   int64
    // 平均心率
    avgHeartrate   int64
    // 速度，单位：千米/小时
    speed   string
    // 动作计数，如：步数、滑水次数
    actionCount   string
    // 路径
    path   string
    // 数据原始来源
    subChannel   string
    // 里程，单位：米
    mileage   int64
    // 爬高，单位：米
    climb   int64
    // 运动持续时间，单位：毫秒
    durationTime   int64
    // 开始位置，格式：经度,维度
    startPoint   string
    // 预留字段
    resultOther   string
    // 最大速度，单位：千米/小时
    maxSpeed   string
    // 结束位置,格式[经度,纬度]
    endPoint   string
    // 过程数据Json
    stage   string
    // 频率
    powerFrequency   int64
    // 消耗卡路里，单位：千卡
    calorie   string
    // 路径节点数据下载地址
    pathDataUrl   string
    // 过程数据下载地址
    stageDataUrl   string
    // 数据类型：0.普通数据 1.赛事数据
    dataType   int64
    // 设备类型
    deviceType   int64
    // 设备型号(厂商)
    deviceModel   string
    // 设备名称
    deviceName   string
    // 三方数据唯一码
    messageId   string
    // 版本号
    version   string
}

// 初始化AlibabaAlisportsDatacenterDatasyncSportsdatasRequest对象
func NewAlibabaAlisportsDatacenterDatasyncSportsdatasRequest() *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest{
    return &AlibabaAlisportsDatacenterDatasyncSportsdatasRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetApiMethodName() string {
    return "alibaba.alisports.datacenter.datasync.sportsdatas"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户阿里体育id
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetUserId() string {
    return r.userId
}
// SportsCat1Id Setter
// 运动一级分类
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsCat1Id(sportsCat1Id int64) error {
    r.sportsCat1Id = sportsCat1Id
    r.Set("sports_cat1_id", sportsCat1Id)
    return nil
}

// SportsCat1Id Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsCat1Id() int64 {
    return r.sportsCat1Id
}
// SportsCat2Id Setter
// 运动二级分类
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsCat2Id(sportsCat2Id int64) error {
    r.sportsCat2Id = sportsCat2Id
    r.Set("sports_cat2_id", sportsCat2Id)
    return nil
}

// SportsCat2Id Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsCat2Id() int64 {
    return r.sportsCat2Id
}
// SportsCat3Id Setter
// 运动三级分类
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsCat3Id(sportsCat3Id string) error {
    r.sportsCat3Id = sportsCat3Id
    r.Set("sports_cat3_id", sportsCat3Id)
    return nil
}

// SportsCat3Id Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsCat3Id() string {
    return r.sportsCat3Id
}
// SportsStartTime Setter
// 运动开始时间，单位：毫秒
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsStartTime(sportsStartTime int64) error {
    r.sportsStartTime = sportsStartTime
    r.Set("sports_start_time", sportsStartTime)
    return nil
}

// SportsStartTime Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsStartTime() int64 {
    return r.sportsStartTime
}
// SportsEndTime Setter
// 运动结束时间，单位：毫秒
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsEndTime(sportsEndTime int64) error {
    r.sportsEndTime = sportsEndTime
    r.Set("sports_end_time", sportsEndTime)
    return nil
}

// SportsEndTime Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsEndTime() int64 {
    return r.sportsEndTime
}
// Timezone Setter
// 时区
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetTimezone(timezone int64) error {
    r.timezone = timezone
    r.Set("timezone", timezone)
    return nil
}

// Timezone Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetTimezone() int64 {
    return r.timezone
}
// MinHeartrate Setter
// 最小心率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMinHeartrate(minHeartrate int64) error {
    r.minHeartrate = minHeartrate
    r.Set("min_heartrate", minHeartrate)
    return nil
}

// MinHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMinHeartrate() int64 {
    return r.minHeartrate
}
// MaxHeartrate Setter
// 最大心率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMaxHeartrate(maxHeartrate int64) error {
    r.maxHeartrate = maxHeartrate
    r.Set("max_heartrate", maxHeartrate)
    return nil
}

// MaxHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMaxHeartrate() int64 {
    return r.maxHeartrate
}
// AvgHeartrate Setter
// 平均心率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetAvgHeartrate(avgHeartrate int64) error {
    r.avgHeartrate = avgHeartrate
    r.Set("avg_heartrate", avgHeartrate)
    return nil
}

// AvgHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetAvgHeartrate() int64 {
    return r.avgHeartrate
}
// Speed Setter
// 速度，单位：千米/小时
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSpeed(speed string) error {
    r.speed = speed
    r.Set("speed", speed)
    return nil
}

// Speed Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSpeed() string {
    return r.speed
}
// ActionCount Setter
// 动作计数，如：步数、滑水次数
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetActionCount(actionCount string) error {
    r.actionCount = actionCount
    r.Set("action_count", actionCount)
    return nil
}

// ActionCount Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetActionCount() string {
    return r.actionCount
}
// Path Setter
// 路径
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetPath(path string) error {
    r.path = path
    r.Set("path", path)
    return nil
}

// Path Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetPath() string {
    return r.path
}
// SubChannel Setter
// 数据原始来源
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSubChannel(subChannel string) error {
    r.subChannel = subChannel
    r.Set("sub_channel", subChannel)
    return nil
}

// SubChannel Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSubChannel() string {
    return r.subChannel
}
// Mileage Setter
// 里程，单位：米
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMileage(mileage int64) error {
    r.mileage = mileage
    r.Set("mileage", mileage)
    return nil
}

// Mileage Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMileage() int64 {
    return r.mileage
}
// Climb Setter
// 爬高，单位：米
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetClimb(climb int64) error {
    r.climb = climb
    r.Set("climb", climb)
    return nil
}

// Climb Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetClimb() int64 {
    return r.climb
}
// DurationTime Setter
// 运动持续时间，单位：毫秒
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDurationTime(durationTime int64) error {
    r.durationTime = durationTime
    r.Set("duration_time", durationTime)
    return nil
}

// DurationTime Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDurationTime() int64 {
    return r.durationTime
}
// StartPoint Setter
// 开始位置，格式：经度,维度
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetStartPoint(startPoint string) error {
    r.startPoint = startPoint
    r.Set("start_point", startPoint)
    return nil
}

// StartPoint Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetStartPoint() string {
    return r.startPoint
}
// ResultOther Setter
// 预留字段
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetResultOther(resultOther string) error {
    r.resultOther = resultOther
    r.Set("result_other", resultOther)
    return nil
}

// ResultOther Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetResultOther() string {
    return r.resultOther
}
// MaxSpeed Setter
// 最大速度，单位：千米/小时
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMaxSpeed(maxSpeed string) error {
    r.maxSpeed = maxSpeed
    r.Set("max_speed", maxSpeed)
    return nil
}

// MaxSpeed Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMaxSpeed() string {
    return r.maxSpeed
}
// EndPoint Setter
// 结束位置,格式[经度,纬度]
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetEndPoint(endPoint string) error {
    r.endPoint = endPoint
    r.Set("end_point", endPoint)
    return nil
}

// EndPoint Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetEndPoint() string {
    return r.endPoint
}
// Stage Setter
// 过程数据Json
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetStage(stage string) error {
    r.stage = stage
    r.Set("stage", stage)
    return nil
}

// Stage Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetStage() string {
    return r.stage
}
// PowerFrequency Setter
// 频率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetPowerFrequency(powerFrequency int64) error {
    r.powerFrequency = powerFrequency
    r.Set("power_frequency", powerFrequency)
    return nil
}

// PowerFrequency Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetPowerFrequency() int64 {
    return r.powerFrequency
}
// Calorie Setter
// 消耗卡路里，单位：千卡
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetCalorie(calorie string) error {
    r.calorie = calorie
    r.Set("calorie", calorie)
    return nil
}

// Calorie Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetCalorie() string {
    return r.calorie
}
// PathDataUrl Setter
// 路径节点数据下载地址
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetPathDataUrl(pathDataUrl string) error {
    r.pathDataUrl = pathDataUrl
    r.Set("path_data_url", pathDataUrl)
    return nil
}

// PathDataUrl Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetPathDataUrl() string {
    return r.pathDataUrl
}
// StageDataUrl Setter
// 过程数据下载地址
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetStageDataUrl(stageDataUrl string) error {
    r.stageDataUrl = stageDataUrl
    r.Set("stage_data_url", stageDataUrl)
    return nil
}

// StageDataUrl Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetStageDataUrl() string {
    return r.stageDataUrl
}
// DataType Setter
// 数据类型：0.普通数据 1.赛事数据
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDataType(dataType int64) error {
    r.dataType = dataType
    r.Set("data_type", dataType)
    return nil
}

// DataType Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDataType() int64 {
    return r.dataType
}
// DeviceType Setter
// 设备类型
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDeviceType() int64 {
    return r.deviceType
}
// DeviceModel Setter
// 设备型号(厂商)
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

// DeviceModel Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDeviceModel() string {
    return r.deviceModel
}
// DeviceName Setter
// 设备名称
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDeviceName(deviceName string) error {
    r.deviceName = deviceName
    r.Set("device_name", deviceName)
    return nil
}

// DeviceName Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDeviceName() string {
    return r.deviceName
}
// MessageId Setter
// 三方数据唯一码
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMessageId(messageId string) error {
    r.messageId = messageId
    r.Set("message_id", messageId)
    return nil
}

// MessageId Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMessageId() string {
    return r.messageId
}
// Version Setter
// 版本号
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetVersion(version string) error {
    r.version = version
    r.Set("version", version)
    return nil
}

// Version Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetVersion() string {
    return r.version
}
