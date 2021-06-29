package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户运动数据同步接口 API请求
alibaba.alisports.data.sports.syncsportsdata

阿里体育数据中心用户运动数据同步接口
*/
type AlibabaAlisportsDataSportsSyncsportsdataRequest struct {
    model.Params
    // 应用appkey
    alispAppKey   string
    // 最高速度 单位：米/每分
    maxSpeed   string
    // 平均心率 单位：次/每分
    averHeartrate   int64
    // 最高心率 单位：次/每分
    maxHeartrate   int64
    // 最低心率 单位：次/每分
    minHeartrate   int64
    // 运动轨迹，有序的经纬度集合，json格式 例：[[1,2],[3,4]] [1,2]是一个坐标点，1是经度 2是纬度（有就传，阿里体育较依赖此字段）
    trail   string
    // 运动开始时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
    stime   string
    // 运动结束时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
    etime   string
    // 设备类型 :1.手环;2.手表;3.跑步机;4.智能运动鞋;5.耳机;6.智能运动鞋;7.智能运动Bra8.智能单车;9.智能跳绳10.智能背心11.脚环
    deviceType   int64
    // 设备名（展示会用到）
    deviceName   string
    // 设备型号
    deviceModel   string
    // 平均速度 单位：米/每分
    averSpeed   string
    // 二级运动量单位 定义：1.爬楼层数(跑步、健走、健身、登山);2.划水次数(游泳)
    subUnit   int64
    // 二级运动量
    subNum   string
    // 时间戳精确到秒
    alispTime   string
    // 接口签名
    alispSign   string
    // 阿里体育用户id
    aliuid   string
    // 业务来源二级分类（中英文）
    source   string
    // 三方运动数据主键id（数据唯一标记，去重使用）
    dataId   string
    // 运动类型一级分类 定义：1-跑步;2-健走;3-骑行;4-游泳;5-健身;6-篮球;7-足球;8-羽毛球;9-排球;10-乒乓球;11-瑜伽;12-电竞;13-登山;16-椭圆机;19-健身操;20-太极;
    sportsClass   int64
    // 运动类型二级分类 定义： 1001-室内跑步;1002-室外跑步;2001-室内健走;2002-室外健走;3001-室内骑行;3002-室外骑行;4001-室内游泳;4002-户外游泳
    sportsType   int64
    // 运动量
    num   string
    // 运动量单位 1.步数(跑步、健走、椭圆机、登山);2.趟数(游泳);3.平均踏频(骑行);
    unit   int64
    // 运动消耗卡路里 单位：卡
    calorie   string
    // 所属赛事（有就传，阿里体育较依赖此字段）
    match   string
    // 运动距离 单位:米（有就传，阿里体育较依赖此字段）
    distance   string
    // 运动时长 单位:秒（有就传，阿里体育较依赖此字段）
    time   string
    // 国家(中英文/运动发生地点，如有就传)
    country   string
    // 省份(中英文/运动发生地点，如有就传)
    province   string
    // 城市(中英文/运动发生地点，如有就传)
    city   string
    // 开始运动地点经纬度，格式：1,2 (1为经度，2为纬度)
    startPoint   string
    // 结束运动地点经纬度，格式：1,2 (1为经度，2为纬度)
    endPoint   string
}

// 初始化AlibabaAlisportsDataSportsSyncsportsdataRequest对象
func NewAlibabaAlisportsDataSportsSyncsportsdataRequest() *AlibabaAlisportsDataSportsSyncsportsdataRequest{
    return &AlibabaAlisportsDataSportsSyncsportsdataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetApiMethodName() string {
    return "alibaba.alisports.data.sports.syncsportsdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 应用appkey
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAlispAppKey() string {
    return r.alispAppKey
}
// MaxSpeed Setter
// 最高速度 单位：米/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetMaxSpeed(maxSpeed string) error {
    r.maxSpeed = maxSpeed
    r.Set("max_speed", maxSpeed)
    return nil
}

// MaxSpeed Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetMaxSpeed() string {
    return r.maxSpeed
}
// AverHeartrate Setter
// 平均心率 单位：次/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAverHeartrate(averHeartrate int64) error {
    r.averHeartrate = averHeartrate
    r.Set("aver_heartrate", averHeartrate)
    return nil
}

// AverHeartrate Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAverHeartrate() int64 {
    return r.averHeartrate
}
// MaxHeartrate Setter
// 最高心率 单位：次/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetMaxHeartrate(maxHeartrate int64) error {
    r.maxHeartrate = maxHeartrate
    r.Set("max_heartrate", maxHeartrate)
    return nil
}

// MaxHeartrate Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetMaxHeartrate() int64 {
    return r.maxHeartrate
}
// MinHeartrate Setter
// 最低心率 单位：次/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetMinHeartrate(minHeartrate int64) error {
    r.minHeartrate = minHeartrate
    r.Set("min_heartrate", minHeartrate)
    return nil
}

// MinHeartrate Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetMinHeartrate() int64 {
    return r.minHeartrate
}
// Trail Setter
// 运动轨迹，有序的经纬度集合，json格式 例：[[1,2],[3,4]] [1,2]是一个坐标点，1是经度 2是纬度（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetTrail(trail string) error {
    r.trail = trail
    r.Set("trail", trail)
    return nil
}

// Trail Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetTrail() string {
    return r.trail
}
// Stime Setter
// 运动开始时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetStime(stime string) error {
    r.stime = stime
    r.Set("stime", stime)
    return nil
}

// Stime Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetStime() string {
    return r.stime
}
// Etime Setter
// 运动结束时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetEtime(etime string) error {
    r.etime = etime
    r.Set("etime", etime)
    return nil
}

// Etime Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetEtime() string {
    return r.etime
}
// DeviceType Setter
// 设备类型 :1.手环;2.手表;3.跑步机;4.智能运动鞋;5.耳机;6.智能运动鞋;7.智能运动Bra8.智能单车;9.智能跳绳10.智能背心11.脚环
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetDeviceType() int64 {
    return r.deviceType
}
// DeviceName Setter
// 设备名（展示会用到）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetDeviceName(deviceName string) error {
    r.deviceName = deviceName
    r.Set("device_name", deviceName)
    return nil
}

// DeviceName Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetDeviceName() string {
    return r.deviceName
}
// DeviceModel Setter
// 设备型号
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

// DeviceModel Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetDeviceModel() string {
    return r.deviceModel
}
// AverSpeed Setter
// 平均速度 单位：米/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAverSpeed(averSpeed string) error {
    r.averSpeed = averSpeed
    r.Set("aver_speed", averSpeed)
    return nil
}

// AverSpeed Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAverSpeed() string {
    return r.averSpeed
}
// SubUnit Setter
// 二级运动量单位 定义：1.爬楼层数(跑步、健走、健身、登山);2.划水次数(游泳)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetSubUnit(subUnit int64) error {
    r.subUnit = subUnit
    r.Set("sub_unit", subUnit)
    return nil
}

// SubUnit Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetSubUnit() int64 {
    return r.subUnit
}
// SubNum Setter
// 二级运动量
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetSubNum(subNum string) error {
    r.subNum = subNum
    r.Set("sub_num", subNum)
    return nil
}

// SubNum Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetSubNum() string {
    return r.subNum
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAlispTime() string {
    return r.alispTime
}
// AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAlispSign() string {
    return r.alispSign
}
// Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAliuid() string {
    return r.aliuid
}
// Source Setter
// 业务来源二级分类（中英文）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetSource() string {
    return r.source
}
// DataId Setter
// 三方运动数据主键id（数据唯一标记，去重使用）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetDataId(dataId string) error {
    r.dataId = dataId
    r.Set("data_id", dataId)
    return nil
}

// DataId Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetDataId() string {
    return r.dataId
}
// SportsClass Setter
// 运动类型一级分类 定义：1-跑步;2-健走;3-骑行;4-游泳;5-健身;6-篮球;7-足球;8-羽毛球;9-排球;10-乒乓球;11-瑜伽;12-电竞;13-登山;16-椭圆机;19-健身操;20-太极;
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetSportsClass(sportsClass int64) error {
    r.sportsClass = sportsClass
    r.Set("sports_class", sportsClass)
    return nil
}

// SportsClass Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetSportsClass() int64 {
    return r.sportsClass
}
// SportsType Setter
// 运动类型二级分类 定义： 1001-室内跑步;1002-室外跑步;2001-室内健走;2002-室外健走;3001-室内骑行;3002-室外骑行;4001-室内游泳;4002-户外游泳
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetSportsType(sportsType int64) error {
    r.sportsType = sportsType
    r.Set("sports_type", sportsType)
    return nil
}

// SportsType Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetSportsType() int64 {
    return r.sportsType
}
// Num Setter
// 运动量
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetNum(num string) error {
    r.num = num
    r.Set("num", num)
    return nil
}

// Num Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetNum() string {
    return r.num
}
// Unit Setter
// 运动量单位 1.步数(跑步、健走、椭圆机、登山);2.趟数(游泳);3.平均踏频(骑行);
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetUnit(unit int64) error {
    r.unit = unit
    r.Set("unit", unit)
    return nil
}

// Unit Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetUnit() int64 {
    return r.unit
}
// Calorie Setter
// 运动消耗卡路里 单位：卡
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetCalorie(calorie string) error {
    r.calorie = calorie
    r.Set("calorie", calorie)
    return nil
}

// Calorie Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetCalorie() string {
    return r.calorie
}
// Match Setter
// 所属赛事（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetMatch(match string) error {
    r.match = match
    r.Set("match", match)
    return nil
}

// Match Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetMatch() string {
    return r.match
}
// Distance Setter
// 运动距离 单位:米（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetDistance(distance string) error {
    r.distance = distance
    r.Set("distance", distance)
    return nil
}

// Distance Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetDistance() string {
    return r.distance
}
// Time Setter
// 运动时长 单位:秒（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetTime(time string) error {
    r.time = time
    r.Set("time", time)
    return nil
}

// Time Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetTime() string {
    return r.time
}
// Country Setter
// 国家(中英文/运动发生地点，如有就传)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

// Country Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetCountry() string {
    return r.country
}
// Province Setter
// 省份(中英文/运动发生地点，如有就传)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

// Province Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetProvince() string {
    return r.province
}
// City Setter
// 城市(中英文/运动发生地点，如有就传)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetCity() string {
    return r.city
}
// StartPoint Setter
// 开始运动地点经纬度，格式：1,2 (1为经度，2为纬度)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetStartPoint(startPoint string) error {
    r.startPoint = startPoint
    r.Set("start_point", startPoint)
    return nil
}

// StartPoint Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetStartPoint() string {
    return r.startPoint
}
// EndPoint Setter
// 结束运动地点经纬度，格式：1,2 (1为经度，2为纬度)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetEndPoint(endPoint string) error {
    r.endPoint = endPoint
    r.Set("end_point", endPoint)
    return nil
}

// EndPoint Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetEndPoint() string {
    return r.endPoint
}
