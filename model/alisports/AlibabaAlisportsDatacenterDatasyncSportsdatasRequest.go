package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育运动数据同步统一接口 APIRequest
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

func NewAlibabaAlisportsDatacenterDatasyncSportsdatasRequest() *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest{
    return &AlibabaAlisportsDatacenterDatasyncSportsdatasRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetApiMethodName() string {
    return "alibaba.alisports.datacenter.datasync.sportsdatas"
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsCat1Id(sportsCat1Id int64) error {
    r.sportsCat1Id = sportsCat1Id
    r.Set("sports_cat1_id", sportsCat1Id)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsCat1Id() int64 {
    return r.sportsCat1Id
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsCat2Id(sportsCat2Id int64) error {
    r.sportsCat2Id = sportsCat2Id
    r.Set("sports_cat2_id", sportsCat2Id)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsCat2Id() int64 {
    return r.sportsCat2Id
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsCat3Id(sportsCat3Id string) error {
    r.sportsCat3Id = sportsCat3Id
    r.Set("sports_cat3_id", sportsCat3Id)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsCat3Id() string {
    return r.sportsCat3Id
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsStartTime(sportsStartTime int64) error {
    r.sportsStartTime = sportsStartTime
    r.Set("sports_start_time", sportsStartTime)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsStartTime() int64 {
    return r.sportsStartTime
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsEndTime(sportsEndTime int64) error {
    r.sportsEndTime = sportsEndTime
    r.Set("sports_end_time", sportsEndTime)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsEndTime() int64 {
    return r.sportsEndTime
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetTimezone(timezone int64) error {
    r.timezone = timezone
    r.Set("timezone", timezone)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetTimezone() int64 {
    return r.timezone
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMinHeartrate(minHeartrate int64) error {
    r.minHeartrate = minHeartrate
    r.Set("min_heartrate", minHeartrate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMinHeartrate() int64 {
    return r.minHeartrate
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMaxHeartrate(maxHeartrate int64) error {
    r.maxHeartrate = maxHeartrate
    r.Set("max_heartrate", maxHeartrate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMaxHeartrate() int64 {
    return r.maxHeartrate
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetAvgHeartrate(avgHeartrate int64) error {
    r.avgHeartrate = avgHeartrate
    r.Set("avg_heartrate", avgHeartrate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetAvgHeartrate() int64 {
    return r.avgHeartrate
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSpeed(speed string) error {
    r.speed = speed
    r.Set("speed", speed)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSpeed() string {
    return r.speed
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetActionCount(actionCount string) error {
    r.actionCount = actionCount
    r.Set("action_count", actionCount)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetActionCount() string {
    return r.actionCount
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetPath(path string) error {
    r.path = path
    r.Set("path", path)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetPath() string {
    return r.path
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSubChannel(subChannel string) error {
    r.subChannel = subChannel
    r.Set("sub_channel", subChannel)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSubChannel() string {
    return r.subChannel
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMileage(mileage int64) error {
    r.mileage = mileage
    r.Set("mileage", mileage)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMileage() int64 {
    return r.mileage
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetClimb(climb int64) error {
    r.climb = climb
    r.Set("climb", climb)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetClimb() int64 {
    return r.climb
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDurationTime(durationTime int64) error {
    r.durationTime = durationTime
    r.Set("duration_time", durationTime)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDurationTime() int64 {
    return r.durationTime
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetStartPoint(startPoint string) error {
    r.startPoint = startPoint
    r.Set("start_point", startPoint)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetStartPoint() string {
    return r.startPoint
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetResultOther(resultOther string) error {
    r.resultOther = resultOther
    r.Set("result_other", resultOther)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetResultOther() string {
    return r.resultOther
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMaxSpeed(maxSpeed string) error {
    r.maxSpeed = maxSpeed
    r.Set("max_speed", maxSpeed)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMaxSpeed() string {
    return r.maxSpeed
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetEndPoint(endPoint string) error {
    r.endPoint = endPoint
    r.Set("end_point", endPoint)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetEndPoint() string {
    return r.endPoint
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetStage(stage string) error {
    r.stage = stage
    r.Set("stage", stage)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetStage() string {
    return r.stage
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetPowerFrequency(powerFrequency int64) error {
    r.powerFrequency = powerFrequency
    r.Set("power_frequency", powerFrequency)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetPowerFrequency() int64 {
    return r.powerFrequency
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetCalorie(calorie string) error {
    r.calorie = calorie
    r.Set("calorie", calorie)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetCalorie() string {
    return r.calorie
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetPathDataUrl(pathDataUrl string) error {
    r.pathDataUrl = pathDataUrl
    r.Set("path_data_url", pathDataUrl)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetPathDataUrl() string {
    return r.pathDataUrl
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetStageDataUrl(stageDataUrl string) error {
    r.stageDataUrl = stageDataUrl
    r.Set("stage_data_url", stageDataUrl)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetStageDataUrl() string {
    return r.stageDataUrl
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDataType(dataType int64) error {
    r.dataType = dataType
    r.Set("data_type", dataType)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDataType() int64 {
    return r.dataType
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDeviceType() int64 {
    return r.deviceType
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDeviceModel() string {
    return r.deviceModel
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDeviceName(deviceName string) error {
    r.deviceName = deviceName
    r.Set("device_name", deviceName)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDeviceName() string {
    return r.deviceName
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMessageId(messageId string) error {
    r.messageId = messageId
    r.Set("message_id", messageId)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMessageId() string {
    return r.messageId
}

func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetVersion(version string) error {
    r.version = version
    r.Set("version", version)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetVersion() string {
    return r.version
}

