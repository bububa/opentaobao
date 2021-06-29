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
    _alispAppKey   string
    // 最高速度 单位：米/每分
    _maxSpeed   string
    // 平均心率 单位：次/每分
    _averHeartrate   int64
    // 最高心率 单位：次/每分
    _maxHeartrate   int64
    // 最低心率 单位：次/每分
    _minHeartrate   int64
    // 运动轨迹，有序的经纬度集合，json格式 例：[[1,2],[3,4]] [1,2]是一个坐标点，1是经度 2是纬度（有就传，阿里体育较依赖此字段）
    _trail   string
    // 运动开始时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
    _stime   string
    // 运动结束时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
    _etime   string
    // 设备类型 :1.手环;2.手表;3.跑步机;4.智能运动鞋;5.耳机;6.智能运动鞋;7.智能运动Bra8.智能单车;9.智能跳绳10.智能背心11.脚环
    _deviceType   int64
    // 设备名（展示会用到）
    _deviceName   string
    // 设备型号
    _deviceModel   string
    // 平均速度 单位：米/每分
    _averSpeed   string
    // 二级运动量单位 定义：1.爬楼层数(跑步、健走、健身、登山);2.划水次数(游泳)
    _subUnit   int64
    // 二级运动量
    _subNum   string
    // 时间戳精确到秒
    _alispTime   string
    // 接口签名
    _alispSign   string
    // 阿里体育用户id
    _aliuid   string
    // 业务来源二级分类（中英文）
    _source   string
    // 三方运动数据主键id（数据唯一标记，去重使用）
    _dataId   string
    // 运动类型一级分类 定义：1-跑步;2-健走;3-骑行;4-游泳;5-健身;6-篮球;7-足球;8-羽毛球;9-排球;10-乒乓球;11-瑜伽;12-电竞;13-登山;16-椭圆机;19-健身操;20-太极;
    _sportsClass   int64
    // 运动类型二级分类 定义： 1001-室内跑步;1002-室外跑步;2001-室内健走;2002-室外健走;3001-室内骑行;3002-室外骑行;4001-室内游泳;4002-户外游泳
    _sportsType   int64
    // 运动量
    _num   string
    // 运动量单位 1.步数(跑步、健走、椭圆机、登山);2.趟数(游泳);3.平均踏频(骑行);
    _unit   int64
    // 运动消耗卡路里 单位：卡
    _calorie   string
    // 所属赛事（有就传，阿里体育较依赖此字段）
    _match   string
    // 运动距离 单位:米（有就传，阿里体育较依赖此字段）
    _distance   string
    // 运动时长 单位:秒（有就传，阿里体育较依赖此字段）
    _time   string
    // 国家(中英文/运动发生地点，如有就传)
    _country   string
    // 省份(中英文/运动发生地点，如有就传)
    _province   string
    // 城市(中英文/运动发生地点，如有就传)
    _city   string
    // 开始运动地点经纬度，格式：1,2 (1为经度，2为纬度)
    _startPoint   string
    // 结束运动地点经纬度，格式：1,2 (1为经度，2为纬度)
    _endPoint   string
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
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// MaxSpeed Setter
// 最高速度 单位：米/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetMaxSpeed(_maxSpeed string) error {
    r._maxSpeed = _maxSpeed
    r.Set("max_speed", _maxSpeed)
    return nil
}

// MaxSpeed Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetMaxSpeed() string {
    return r._maxSpeed
}
// AverHeartrate Setter
// 平均心率 单位：次/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAverHeartrate(_averHeartrate int64) error {
    r._averHeartrate = _averHeartrate
    r.Set("aver_heartrate", _averHeartrate)
    return nil
}

// AverHeartrate Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAverHeartrate() int64 {
    return r._averHeartrate
}
// MaxHeartrate Setter
// 最高心率 单位：次/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetMaxHeartrate(_maxHeartrate int64) error {
    r._maxHeartrate = _maxHeartrate
    r.Set("max_heartrate", _maxHeartrate)
    return nil
}

// MaxHeartrate Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetMaxHeartrate() int64 {
    return r._maxHeartrate
}
// MinHeartrate Setter
// 最低心率 单位：次/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetMinHeartrate(_minHeartrate int64) error {
    r._minHeartrate = _minHeartrate
    r.Set("min_heartrate", _minHeartrate)
    return nil
}

// MinHeartrate Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetMinHeartrate() int64 {
    return r._minHeartrate
}
// Trail Setter
// 运动轨迹，有序的经纬度集合，json格式 例：[[1,2],[3,4]] [1,2]是一个坐标点，1是经度 2是纬度（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetTrail(_trail string) error {
    r._trail = _trail
    r.Set("trail", _trail)
    return nil
}

// Trail Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetTrail() string {
    return r._trail
}
// Stime Setter
// 运动开始时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetStime(_stime string) error {
    r._stime = _stime
    r.Set("stime", _stime)
    return nil
}

// Stime Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetStime() string {
    return r._stime
}
// Etime Setter
// 运动结束时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetEtime(_etime string) error {
    r._etime = _etime
    r.Set("etime", _etime)
    return nil
}

// Etime Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetEtime() string {
    return r._etime
}
// DeviceType Setter
// 设备类型 :1.手环;2.手表;3.跑步机;4.智能运动鞋;5.耳机;6.智能运动鞋;7.智能运动Bra8.智能单车;9.智能跳绳10.智能背心11.脚环
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetDeviceType(_deviceType int64) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetDeviceType() int64 {
    return r._deviceType
}
// DeviceName Setter
// 设备名（展示会用到）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetDeviceName(_deviceName string) error {
    r._deviceName = _deviceName
    r.Set("device_name", _deviceName)
    return nil
}

// DeviceName Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetDeviceName() string {
    return r._deviceName
}
// DeviceModel Setter
// 设备型号
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetDeviceModel(_deviceModel string) error {
    r._deviceModel = _deviceModel
    r.Set("device_model", _deviceModel)
    return nil
}

// DeviceModel Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetDeviceModel() string {
    return r._deviceModel
}
// AverSpeed Setter
// 平均速度 单位：米/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAverSpeed(_averSpeed string) error {
    r._averSpeed = _averSpeed
    r.Set("aver_speed", _averSpeed)
    return nil
}

// AverSpeed Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAverSpeed() string {
    return r._averSpeed
}
// SubUnit Setter
// 二级运动量单位 定义：1.爬楼层数(跑步、健走、健身、登山);2.划水次数(游泳)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetSubUnit(_subUnit int64) error {
    r._subUnit = _subUnit
    r.Set("sub_unit", _subUnit)
    return nil
}

// SubUnit Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetSubUnit() int64 {
    return r._subUnit
}
// SubNum Setter
// 二级运动量
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetSubNum(_subNum string) error {
    r._subNum = _subNum
    r.Set("sub_num", _subNum)
    return nil
}

// SubNum Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetSubNum() string {
    return r._subNum
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAlispTime() string {
    return r._alispTime
}
// AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAlispSign() string {
    return r._alispSign
}
// Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetAliuid(_aliuid string) error {
    r._aliuid = _aliuid
    r.Set("aliuid", _aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetAliuid() string {
    return r._aliuid
}
// Source Setter
// 业务来源二级分类（中英文）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetSource() string {
    return r._source
}
// DataId Setter
// 三方运动数据主键id（数据唯一标记，去重使用）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetDataId(_dataId string) error {
    r._dataId = _dataId
    r.Set("data_id", _dataId)
    return nil
}

// DataId Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetDataId() string {
    return r._dataId
}
// SportsClass Setter
// 运动类型一级分类 定义：1-跑步;2-健走;3-骑行;4-游泳;5-健身;6-篮球;7-足球;8-羽毛球;9-排球;10-乒乓球;11-瑜伽;12-电竞;13-登山;16-椭圆机;19-健身操;20-太极;
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetSportsClass(_sportsClass int64) error {
    r._sportsClass = _sportsClass
    r.Set("sports_class", _sportsClass)
    return nil
}

// SportsClass Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetSportsClass() int64 {
    return r._sportsClass
}
// SportsType Setter
// 运动类型二级分类 定义： 1001-室内跑步;1002-室外跑步;2001-室内健走;2002-室外健走;3001-室内骑行;3002-室外骑行;4001-室内游泳;4002-户外游泳
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetSportsType(_sportsType int64) error {
    r._sportsType = _sportsType
    r.Set("sports_type", _sportsType)
    return nil
}

// SportsType Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetSportsType() int64 {
    return r._sportsType
}
// Num Setter
// 运动量
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetNum(_num string) error {
    r._num = _num
    r.Set("num", _num)
    return nil
}

// Num Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetNum() string {
    return r._num
}
// Unit Setter
// 运动量单位 1.步数(跑步、健走、椭圆机、登山);2.趟数(游泳);3.平均踏频(骑行);
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetUnit(_unit int64) error {
    r._unit = _unit
    r.Set("unit", _unit)
    return nil
}

// Unit Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetUnit() int64 {
    return r._unit
}
// Calorie Setter
// 运动消耗卡路里 单位：卡
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetCalorie(_calorie string) error {
    r._calorie = _calorie
    r.Set("calorie", _calorie)
    return nil
}

// Calorie Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetCalorie() string {
    return r._calorie
}
// Match Setter
// 所属赛事（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetMatch(_match string) error {
    r._match = _match
    r.Set("match", _match)
    return nil
}

// Match Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetMatch() string {
    return r._match
}
// Distance Setter
// 运动距离 单位:米（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetDistance(_distance string) error {
    r._distance = _distance
    r.Set("distance", _distance)
    return nil
}

// Distance Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetDistance() string {
    return r._distance
}
// Time Setter
// 运动时长 单位:秒（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetTime(_time string) error {
    r._time = _time
    r.Set("time", _time)
    return nil
}

// Time Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetTime() string {
    return r._time
}
// Country Setter
// 国家(中英文/运动发生地点，如有就传)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetCountry(_country string) error {
    r._country = _country
    r.Set("country", _country)
    return nil
}

// Country Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetCountry() string {
    return r._country
}
// Province Setter
// 省份(中英文/运动发生地点，如有就传)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetProvince(_province string) error {
    r._province = _province
    r.Set("province", _province)
    return nil
}

// Province Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetProvince() string {
    return r._province
}
// City Setter
// 城市(中英文/运动发生地点，如有就传)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetCity() string {
    return r._city
}
// StartPoint Setter
// 开始运动地点经纬度，格式：1,2 (1为经度，2为纬度)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetStartPoint(_startPoint string) error {
    r._startPoint = _startPoint
    r.Set("start_point", _startPoint)
    return nil
}

// StartPoint Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetStartPoint() string {
    return r._startPoint
}
// EndPoint Setter
// 结束运动地点经纬度，格式：1,2 (1为经度，2为纬度)
func (r *AlibabaAlisportsDataSportsSyncsportsdataRequest) SetEndPoint(_endPoint string) error {
    r._endPoint = _endPoint
    r.Set("end_point", _endPoint)
    return nil
}

// EndPoint Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataRequest) GetEndPoint() string {
    return r._endPoint
}
