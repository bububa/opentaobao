package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest 阿里体育运动数据同步统一接口 API请求
// alibaba.alisports.datacenter.datasync.sportsdatas
//
// 给单方提供同步运动数据到阿里体育的接口
type AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest struct {
	model.Params
	// 用户阿里体育id
	_userId string
	// 运动一级分类
	_sportsCat1Id int64
	// 运动二级分类
	_sportsCat2Id int64
	// 运动三级分类
	_sportsCat3Id string
	// 运动开始时间，单位：毫秒
	_sportsStartTime int64
	// 运动结束时间，单位：毫秒
	_sportsEndTime int64
	// 时区
	_timezone int64
	// 最小心率
	_minHeartrate int64
	// 最大心率
	_maxHeartrate int64
	// 平均心率
	_avgHeartrate int64
	// 速度，单位：千米/小时
	_speed string
	// 动作计数，如：步数、滑水次数
	_actionCount string
	// 路径
	_path string
	// 数据原始来源
	_subChannel string
	// 里程，单位：米
	_mileage int64
	// 爬高，单位：米
	_climb int64
	// 运动持续时间，单位：毫秒
	_durationTime int64
	// 开始位置，格式：经度,维度
	_startPoint string
	// 预留字段
	_resultOther string
	// 最大速度，单位：千米/小时
	_maxSpeed string
	// 结束位置,格式[经度,纬度]
	_endPoint string
	// 过程数据Json
	_stage string
	// 频率
	_powerFrequency int64
	// 消耗卡路里，单位：千卡
	_calorie string
	// 路径节点数据下载地址
	_pathDataUrl string
	// 过程数据下载地址
	_stageDataUrl string
	// 数据类型：0.普通数据 1.赛事数据
	_dataType int64
	// 设备类型
	_deviceType int64
	// 设备型号(厂商)
	_deviceModel string
	// 设备名称
	_deviceName string
	// 三方数据唯一码
	_messageId string
	// 版本号
	_version string
}

// NewAlibabaAlisportsDatacenterDatasyncSportsdatasRequest 初始化AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest对象
func NewAlibabaAlisportsDatacenterDatasyncSportsdatasRequest() *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest {
	return &AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.datacenter.datasync.sportsdatas"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserId is UserId Setter
// 用户阿里体育id
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetUserId() string {
	return r._userId
}

// SetSportsCat1Id is SportsCat1Id Setter
// 运动一级分类
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetSportsCat1Id(_sportsCat1Id int64) error {
	r._sportsCat1Id = _sportsCat1Id
	r.Set("sports_cat1_id", _sportsCat1Id)
	return nil
}

// GetSportsCat1Id SportsCat1Id Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetSportsCat1Id() int64 {
	return r._sportsCat1Id
}

// SetSportsCat2Id is SportsCat2Id Setter
// 运动二级分类
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetSportsCat2Id(_sportsCat2Id int64) error {
	r._sportsCat2Id = _sportsCat2Id
	r.Set("sports_cat2_id", _sportsCat2Id)
	return nil
}

// GetSportsCat2Id SportsCat2Id Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetSportsCat2Id() int64 {
	return r._sportsCat2Id
}

// SetSportsCat3Id is SportsCat3Id Setter
// 运动三级分类
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetSportsCat3Id(_sportsCat3Id string) error {
	r._sportsCat3Id = _sportsCat3Id
	r.Set("sports_cat3_id", _sportsCat3Id)
	return nil
}

// GetSportsCat3Id SportsCat3Id Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetSportsCat3Id() string {
	return r._sportsCat3Id
}

// SetSportsStartTime is SportsStartTime Setter
// 运动开始时间，单位：毫秒
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetSportsStartTime(_sportsStartTime int64) error {
	r._sportsStartTime = _sportsStartTime
	r.Set("sports_start_time", _sportsStartTime)
	return nil
}

// GetSportsStartTime SportsStartTime Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetSportsStartTime() int64 {
	return r._sportsStartTime
}

// SetSportsEndTime is SportsEndTime Setter
// 运动结束时间，单位：毫秒
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetSportsEndTime(_sportsEndTime int64) error {
	r._sportsEndTime = _sportsEndTime
	r.Set("sports_end_time", _sportsEndTime)
	return nil
}

// GetSportsEndTime SportsEndTime Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetSportsEndTime() int64 {
	return r._sportsEndTime
}

// SetTimezone is Timezone Setter
// 时区
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetTimezone(_timezone int64) error {
	r._timezone = _timezone
	r.Set("timezone", _timezone)
	return nil
}

// GetTimezone Timezone Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetTimezone() int64 {
	return r._timezone
}

// SetMinHeartrate is MinHeartrate Setter
// 最小心率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetMinHeartrate(_minHeartrate int64) error {
	r._minHeartrate = _minHeartrate
	r.Set("min_heartrate", _minHeartrate)
	return nil
}

// GetMinHeartrate MinHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetMinHeartrate() int64 {
	return r._minHeartrate
}

// SetMaxHeartrate is MaxHeartrate Setter
// 最大心率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetMaxHeartrate(_maxHeartrate int64) error {
	r._maxHeartrate = _maxHeartrate
	r.Set("max_heartrate", _maxHeartrate)
	return nil
}

// GetMaxHeartrate MaxHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetMaxHeartrate() int64 {
	return r._maxHeartrate
}

// SetAvgHeartrate is AvgHeartrate Setter
// 平均心率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetAvgHeartrate(_avgHeartrate int64) error {
	r._avgHeartrate = _avgHeartrate
	r.Set("avg_heartrate", _avgHeartrate)
	return nil
}

// GetAvgHeartrate AvgHeartrate Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetAvgHeartrate() int64 {
	return r._avgHeartrate
}

// SetSpeed is Speed Setter
// 速度，单位：千米/小时
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetSpeed(_speed string) error {
	r._speed = _speed
	r.Set("speed", _speed)
	return nil
}

// GetSpeed Speed Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetSpeed() string {
	return r._speed
}

// SetActionCount is ActionCount Setter
// 动作计数，如：步数、滑水次数
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetActionCount(_actionCount string) error {
	r._actionCount = _actionCount
	r.Set("action_count", _actionCount)
	return nil
}

// GetActionCount ActionCount Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetActionCount() string {
	return r._actionCount
}

// SetPath is Path Setter
// 路径
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetPath(_path string) error {
	r._path = _path
	r.Set("path", _path)
	return nil
}

// GetPath Path Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetPath() string {
	return r._path
}

// SetSubChannel is SubChannel Setter
// 数据原始来源
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetSubChannel(_subChannel string) error {
	r._subChannel = _subChannel
	r.Set("sub_channel", _subChannel)
	return nil
}

// GetSubChannel SubChannel Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetSubChannel() string {
	return r._subChannel
}

// SetMileage is Mileage Setter
// 里程，单位：米
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetMileage(_mileage int64) error {
	r._mileage = _mileage
	r.Set("mileage", _mileage)
	return nil
}

// GetMileage Mileage Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetMileage() int64 {
	return r._mileage
}

// SetClimb is Climb Setter
// 爬高，单位：米
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetClimb(_climb int64) error {
	r._climb = _climb
	r.Set("climb", _climb)
	return nil
}

// GetClimb Climb Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetClimb() int64 {
	return r._climb
}

// SetDurationTime is DurationTime Setter
// 运动持续时间，单位：毫秒
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetDurationTime(_durationTime int64) error {
	r._durationTime = _durationTime
	r.Set("duration_time", _durationTime)
	return nil
}

// GetDurationTime DurationTime Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetDurationTime() int64 {
	return r._durationTime
}

// SetStartPoint is StartPoint Setter
// 开始位置，格式：经度,维度
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetStartPoint(_startPoint string) error {
	r._startPoint = _startPoint
	r.Set("start_point", _startPoint)
	return nil
}

// GetStartPoint StartPoint Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetStartPoint() string {
	return r._startPoint
}

// SetResultOther is ResultOther Setter
// 预留字段
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetResultOther(_resultOther string) error {
	r._resultOther = _resultOther
	r.Set("result_other", _resultOther)
	return nil
}

// GetResultOther ResultOther Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetResultOther() string {
	return r._resultOther
}

// SetMaxSpeed is MaxSpeed Setter
// 最大速度，单位：千米/小时
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetMaxSpeed(_maxSpeed string) error {
	r._maxSpeed = _maxSpeed
	r.Set("max_speed", _maxSpeed)
	return nil
}

// GetMaxSpeed MaxSpeed Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetMaxSpeed() string {
	return r._maxSpeed
}

// SetEndPoint is EndPoint Setter
// 结束位置,格式[经度,纬度]
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetEndPoint(_endPoint string) error {
	r._endPoint = _endPoint
	r.Set("end_point", _endPoint)
	return nil
}

// GetEndPoint EndPoint Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetEndPoint() string {
	return r._endPoint
}

// SetStage is Stage Setter
// 过程数据Json
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetStage(_stage string) error {
	r._stage = _stage
	r.Set("stage", _stage)
	return nil
}

// GetStage Stage Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetStage() string {
	return r._stage
}

// SetPowerFrequency is PowerFrequency Setter
// 频率
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetPowerFrequency(_powerFrequency int64) error {
	r._powerFrequency = _powerFrequency
	r.Set("power_frequency", _powerFrequency)
	return nil
}

// GetPowerFrequency PowerFrequency Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetPowerFrequency() int64 {
	return r._powerFrequency
}

// SetCalorie is Calorie Setter
// 消耗卡路里，单位：千卡
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetCalorie(_calorie string) error {
	r._calorie = _calorie
	r.Set("calorie", _calorie)
	return nil
}

// GetCalorie Calorie Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetCalorie() string {
	return r._calorie
}

// SetPathDataUrl is PathDataUrl Setter
// 路径节点数据下载地址
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetPathDataUrl(_pathDataUrl string) error {
	r._pathDataUrl = _pathDataUrl
	r.Set("path_data_url", _pathDataUrl)
	return nil
}

// GetPathDataUrl PathDataUrl Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetPathDataUrl() string {
	return r._pathDataUrl
}

// SetStageDataUrl is StageDataUrl Setter
// 过程数据下载地址
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetStageDataUrl(_stageDataUrl string) error {
	r._stageDataUrl = _stageDataUrl
	r.Set("stage_data_url", _stageDataUrl)
	return nil
}

// GetStageDataUrl StageDataUrl Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetStageDataUrl() string {
	return r._stageDataUrl
}

// SetDataType is DataType Setter
// 数据类型：0.普通数据 1.赛事数据
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetDataType(_dataType int64) error {
	r._dataType = _dataType
	r.Set("data_type", _dataType)
	return nil
}

// GetDataType DataType Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetDataType() int64 {
	return r._dataType
}

// SetDeviceType is DeviceType Setter
// 设备类型
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetDeviceType() int64 {
	return r._deviceType
}

// SetDeviceModel is DeviceModel Setter
// 设备型号(厂商)
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// GetDeviceModel DeviceModel Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

// SetDeviceName is DeviceName Setter
// 设备名称
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetDeviceName(_deviceName string) error {
	r._deviceName = _deviceName
	r.Set("device_name", _deviceName)
	return nil
}

// GetDeviceName DeviceName Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetDeviceName() string {
	return r._deviceName
}

// SetMessageId is MessageId Setter
// 三方数据唯一码
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetMessageId(_messageId string) error {
	r._messageId = _messageId
	r.Set("message_id", _messageId)
	return nil
}

// GetMessageId MessageId Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetMessageId() string {
	return r._messageId
}

// SetVersion is Version Setter
// 版本号
func (r *AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) SetVersion(_version string) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest) GetVersion() string {
	return r._version
}
