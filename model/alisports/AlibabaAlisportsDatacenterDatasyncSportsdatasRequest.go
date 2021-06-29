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
    _userId   string
    // 运动一级分类
    _sportsCat1Id   int64
    // 运动二级分类
    _sportsCat2Id   int64
    // 运动三级分类
    _sportsCat3Id   string
    // 运动开始时间，单位：毫秒
    _sportsStartTime   int64
    // 运动结束时间，单位：毫秒
    _sportsEndTime   int64
    // 时区
    _timezone   int64
    // 最小心率
    _minHeartrate   int64
    // 最大心率
    _maxHeartrate   int64
    // 平均心率
    _avgHeartrate   int64
    // 速度，单位：千米/小时
    _speed   string
    // 动作计数，如：步数、滑水次数
    _actionCount   string
    // 路径
    _path   string
    // 数据原始来源
    _subChannel   string
    // 里程，单位：米
    _mileage   int64
    // 爬高，单位：米
    _climb   int64
    // 运动持续时间，单位：毫秒
    _durationTime   int64
    // 开始位置，格式：经度,维度
    _startPoint   string
    // 预留字段
    _resultOther   string
    // 最大速度，单位：千米/小时
    _maxSpeed   string
    // 结束位置,格式[经度,纬度]
    _endPoint   string
    // 过程数据Json
    _stage   string
    // 频率
    _powerFrequency   int64
    // 消耗卡路里，单位：千卡
    _calorie   string
    // 路径节点数据下载地址
    _pathDataUrl   string
    // 过程数据下载地址
    _stageDataUrl   string
    // 数据类型：0.普通数据 1.赛事数据
    _dataType   int64
    // 设备类型
    _deviceType   int64
    // 设备型号(厂商)
    _deviceModel   string
    // 设备名称
    _deviceName   string
    // 三方数据唯一码
    _messageId   string
    // 版本号
    _version   string
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
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetUserId() string {
    return r._userId
}
// SportsCat1Id Setter
// 运动一级分类
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsCat1Id(_sportsCat1Id int64) error {
    r._sportsCat1Id = _sportsCat1Id
    r.Set("sports_cat1_id", _sportsCat1Id)
    return nil
}

// SportsCat1Id Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsCat1Id() int64 {
    return r._sportsCat1Id
}
// SportsCat2Id Setter
// 运动二级分类
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsCat2Id(_sportsCat2Id int64) error {
    r._sportsCat2Id = _sportsCat2Id
    r.Set("sports_cat2_id", _sportsCat2Id)
    return nil
}

// SportsCat2Id Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsCat2Id() int64 {
    return r._sportsCat2Id
}
// SportsCat3Id Setter
// 运动三级分类
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsCat3Id(_sportsCat3Id string) error {
    r._sportsCat3Id = _sportsCat3Id
    r.Set("sports_cat3_id", _sportsCat3Id)
    return nil
}

// SportsCat3Id Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsCat3Id() string {
    return r._sportsCat3Id
}
// SportsStartTime Setter
// 运动开始时间，单位：毫秒
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsStartTime(_sportsStartTime int64) error {
    r._sportsStartTime = _sportsStartTime
    r.Set("sports_start_time", _sportsStartTime)
    return nil
}

// SportsStartTime Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsStartTime() int64 {
    return r._sportsStartTime
}
// SportsEndTime Setter
// 运动结束时间，单位：毫秒
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSportsEndTime(_sportsEndTime int64) error {
    r._sportsEndTime = _sportsEndTime
    r.Set("sports_end_time", _sportsEndTime)
    return nil
}

// SportsEndTime Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSportsEndTime() int64 {
    return r._sportsEndTime
}
// Timezone Setter
// 时区
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetTimezone(_timezone int64) error {
    r._timezone = _timezone
    r.Set("timezone", _timezone)
    return nil
}

// Timezone Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetTimezone() int64 {
    return r._timezone
}
// MinHeartrate Setter
// 最小心率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMinHeartrate(_minHeartrate int64) error {
    r._minHeartrate = _minHeartrate
    r.Set("min_heartrate", _minHeartrate)
    return nil
}

// MinHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMinHeartrate() int64 {
    return r._minHeartrate
}
// MaxHeartrate Setter
// 最大心率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMaxHeartrate(_maxHeartrate int64) error {
    r._maxHeartrate = _maxHeartrate
    r.Set("max_heartrate", _maxHeartrate)
    return nil
}

// MaxHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMaxHeartrate() int64 {
    return r._maxHeartrate
}
// AvgHeartrate Setter
// 平均心率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetAvgHeartrate(_avgHeartrate int64) error {
    r._avgHeartrate = _avgHeartrate
    r.Set("avg_heartrate", _avgHeartrate)
    return nil
}

// AvgHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetAvgHeartrate() int64 {
    return r._avgHeartrate
}
// Speed Setter
// 速度，单位：千米/小时
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSpeed(_speed string) error {
    r._speed = _speed
    r.Set("speed", _speed)
    return nil
}

// Speed Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSpeed() string {
    return r._speed
}
// ActionCount Setter
// 动作计数，如：步数、滑水次数
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetActionCount(_actionCount string) error {
    r._actionCount = _actionCount
    r.Set("action_count", _actionCount)
    return nil
}

// ActionCount Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetActionCount() string {
    return r._actionCount
}
// Path Setter
// 路径
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetPath(_path string) error {
    r._path = _path
    r.Set("path", _path)
    return nil
}

// Path Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetPath() string {
    return r._path
}
// SubChannel Setter
// 数据原始来源
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetSubChannel(_subChannel string) error {
    r._subChannel = _subChannel
    r.Set("sub_channel", _subChannel)
    return nil
}

// SubChannel Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetSubChannel() string {
    return r._subChannel
}
// Mileage Setter
// 里程，单位：米
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMileage(_mileage int64) error {
    r._mileage = _mileage
    r.Set("mileage", _mileage)
    return nil
}

// Mileage Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMileage() int64 {
    return r._mileage
}
// Climb Setter
// 爬高，单位：米
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetClimb(_climb int64) error {
    r._climb = _climb
    r.Set("climb", _climb)
    return nil
}

// Climb Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetClimb() int64 {
    return r._climb
}
// DurationTime Setter
// 运动持续时间，单位：毫秒
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDurationTime(_durationTime int64) error {
    r._durationTime = _durationTime
    r.Set("duration_time", _durationTime)
    return nil
}

// DurationTime Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDurationTime() int64 {
    return r._durationTime
}
// StartPoint Setter
// 开始位置，格式：经度,维度
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetStartPoint(_startPoint string) error {
    r._startPoint = _startPoint
    r.Set("start_point", _startPoint)
    return nil
}

// StartPoint Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetStartPoint() string {
    return r._startPoint
}
// ResultOther Setter
// 预留字段
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetResultOther(_resultOther string) error {
    r._resultOther = _resultOther
    r.Set("result_other", _resultOther)
    return nil
}

// ResultOther Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetResultOther() string {
    return r._resultOther
}
// MaxSpeed Setter
// 最大速度，单位：千米/小时
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMaxSpeed(_maxSpeed string) error {
    r._maxSpeed = _maxSpeed
    r.Set("max_speed", _maxSpeed)
    return nil
}

// MaxSpeed Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMaxSpeed() string {
    return r._maxSpeed
}
// EndPoint Setter
// 结束位置,格式[经度,纬度]
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetEndPoint(_endPoint string) error {
    r._endPoint = _endPoint
    r.Set("end_point", _endPoint)
    return nil
}

// EndPoint Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetEndPoint() string {
    return r._endPoint
}
// Stage Setter
// 过程数据Json
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetStage(_stage string) error {
    r._stage = _stage
    r.Set("stage", _stage)
    return nil
}

// Stage Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetStage() string {
    return r._stage
}
// PowerFrequency Setter
// 频率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetPowerFrequency(_powerFrequency int64) error {
    r._powerFrequency = _powerFrequency
    r.Set("power_frequency", _powerFrequency)
    return nil
}

// PowerFrequency Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetPowerFrequency() int64 {
    return r._powerFrequency
}
// Calorie Setter
// 消耗卡路里，单位：千卡
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetCalorie(_calorie string) error {
    r._calorie = _calorie
    r.Set("calorie", _calorie)
    return nil
}

// Calorie Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetCalorie() string {
    return r._calorie
}
// PathDataUrl Setter
// 路径节点数据下载地址
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetPathDataUrl(_pathDataUrl string) error {
    r._pathDataUrl = _pathDataUrl
    r.Set("path_data_url", _pathDataUrl)
    return nil
}

// PathDataUrl Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetPathDataUrl() string {
    return r._pathDataUrl
}
// StageDataUrl Setter
// 过程数据下载地址
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetStageDataUrl(_stageDataUrl string) error {
    r._stageDataUrl = _stageDataUrl
    r.Set("stage_data_url", _stageDataUrl)
    return nil
}

// StageDataUrl Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetStageDataUrl() string {
    return r._stageDataUrl
}
// DataType Setter
// 数据类型：0.普通数据 1.赛事数据
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDataType(_dataType int64) error {
    r._dataType = _dataType
    r.Set("data_type", _dataType)
    return nil
}

// DataType Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDataType() int64 {
    return r._dataType
}
// DeviceType Setter
// 设备类型
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDeviceType(_deviceType int64) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDeviceType() int64 {
    return r._deviceType
}
// DeviceModel Setter
// 设备型号(厂商)
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDeviceModel(_deviceModel string) error {
    r._deviceModel = _deviceModel
    r.Set("device_model", _deviceModel)
    return nil
}

// DeviceModel Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDeviceModel() string {
    return r._deviceModel
}
// DeviceName Setter
// 设备名称
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetDeviceName(_deviceName string) error {
    r._deviceName = _deviceName
    r.Set("device_name", _deviceName)
    return nil
}

// DeviceName Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetDeviceName() string {
    return r._deviceName
}
// MessageId Setter
// 三方数据唯一码
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetMessageId(_messageId string) error {
    r._messageId = _messageId
    r.Set("message_id", _messageId)
    return nil
}

// MessageId Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetMessageId() string {
    return r._messageId
}
// Version Setter
// 版本号
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) SetVersion(_version string) error {
    r._version = _version
    r.Set("version", _version)
    return nil
}

// Version Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasRequest) GetVersion() string {
    return r._version
}
