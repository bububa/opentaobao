package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsDataSportsSyncsportsdataAPIRequest
阿里体育数据中心用户运动数据同步接口 API请求
alibaba.alisports.data.sports.syncsportsdata

阿里体育数据中心用户运动数据同步接口 */
type AlibabaAlisportsDataSportsSyncsportsdataAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 最高速度 单位：米/每分
	_maxSpeed string
	// 平均心率 单位：次/每分
	_averHeartrate int64
	// 最高心率 单位：次/每分
	_maxHeartrate int64
	// 最低心率 单位：次/每分
	_minHeartrate int64
	// 运动轨迹，有序的经纬度集合，json格式 例：[[1,2],[3,4]] [1,2]是一个坐标点，1是经度 2是纬度（有就传，阿里体育较依赖此字段）
	_trail string
	// 运动开始时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
	_stime string
	// 运动结束时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
	_etime string
	// 设备类型 :1.手环;2.手表;3.跑步机;4.智能运动鞋;5.耳机;6.智能运动鞋;7.智能运动Bra8.智能单车;9.智能跳绳10.智能背心11.脚环
	_deviceType int64
	// 设备名（展示会用到）
	_deviceName string
	// 设备型号
	_deviceModel string
	// 平均速度 单位：米/每分
	_averSpeed string
	// 二级运动量单位 定义：1.爬楼层数(跑步、健走、健身、登山);2.划水次数(游泳)
	_subUnit int64
	// 二级运动量
	_subNum string
	// 时间戳精确到秒
	_alispTime string
	// 接口签名
	_alispSign string
	// 阿里体育用户id
	_aliuid string
	// 业务来源二级分类（中英文）
	_source string
	// 三方运动数据主键id（数据唯一标记，去重使用）
	_dataId string
	// 运动类型一级分类 定义：1-跑步;2-健走;3-骑行;4-游泳;5-健身;6-篮球;7-足球;8-羽毛球;9-排球;10-乒乓球;11-瑜伽;12-电竞;13-登山;16-椭圆机;19-健身操;20-太极;
	_sportsClass int64
	// 运动类型二级分类 定义： 1001-室内跑步;1002-室外跑步;2001-室内健走;2002-室外健走;3001-室内骑行;3002-室外骑行;4001-室内游泳;4002-户外游泳
	_sportsType int64
	// 运动量
	_num string
	// 运动量单位 1.步数(跑步、健走、椭圆机、登山);2.趟数(游泳);3.平均踏频(骑行);
	_unit int64
	// 运动消耗卡路里 单位：卡
	_calorie string
	// 所属赛事（有就传，阿里体育较依赖此字段）
	_match string
	// 运动距离 单位:米（有就传，阿里体育较依赖此字段）
	_distance string
	// 运动时长 单位:秒（有就传，阿里体育较依赖此字段）
	_time string
	// 国家(中英文/运动发生地点，如有就传)
	_country string
	// 省份(中英文/运动发生地点，如有就传)
	_province string
	// 城市(中英文/运动发生地点，如有就传)
	_city string
	// 开始运动地点经纬度，格式：1,2 (1为经度，2为纬度)
	_startPoint string
	// 结束运动地点经纬度，格式：1,2 (1为经度，2为纬度)
	_endPoint string
}

// NewAlibabaAlisportsDataSportsSyncsportsdataRequest 初始化AlibabaAlisportsDataSportsSyncsportsdataAPIRequest对象
func NewAlibabaAlisportsDataSportsSyncsportsdataRequest() *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest {
	return &AlibabaAlisportsDataSportsSyncsportsdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.data.sports.syncsportsdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlispAppKey Setter
// 应用appkey
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// Get AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// Set is MaxSpeed Setter
// 最高速度 单位：米/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetMaxSpeed(_maxSpeed string) error {
	r._maxSpeed = _maxSpeed
	r.Set("max_speed", _maxSpeed)
	return nil
}

// Get MaxSpeed Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetMaxSpeed() string {
	return r._maxSpeed
}

// Set is AverHeartrate Setter
// 平均心率 单位：次/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetAverHeartrate(_averHeartrate int64) error {
	r._averHeartrate = _averHeartrate
	r.Set("aver_heartrate", _averHeartrate)
	return nil
}

// Get AverHeartrate Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetAverHeartrate() int64 {
	return r._averHeartrate
}

// Set is MaxHeartrate Setter
// 最高心率 单位：次/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetMaxHeartrate(_maxHeartrate int64) error {
	r._maxHeartrate = _maxHeartrate
	r.Set("max_heartrate", _maxHeartrate)
	return nil
}

// Get MaxHeartrate Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetMaxHeartrate() int64 {
	return r._maxHeartrate
}

// Set is MinHeartrate Setter
// 最低心率 单位：次/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetMinHeartrate(_minHeartrate int64) error {
	r._minHeartrate = _minHeartrate
	r.Set("min_heartrate", _minHeartrate)
	return nil
}

// Get MinHeartrate Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetMinHeartrate() int64 {
	return r._minHeartrate
}

// Set is Trail Setter
// 运动轨迹，有序的经纬度集合，json格式 例：[[1,2],[3,4]] [1,2]是一个坐标点，1是经度 2是纬度（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetTrail(_trail string) error {
	r._trail = _trail
	r.Set("trail", _trail)
	return nil
}

// Get Trail Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetTrail() string {
	return r._trail
}

// Set is Stime Setter
// 运动开始时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetStime(_stime string) error {
	r._stime = _stime
	r.Set("stime", _stime)
	return nil
}

// Get Stime Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetStime() string {
	return r._stime
}

// Set is Etime Setter
// 运动结束时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetEtime(_etime string) error {
	r._etime = _etime
	r.Set("etime", _etime)
	return nil
}

// Get Etime Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetEtime() string {
	return r._etime
}

// Set is DeviceType Setter
// 设备类型 :1.手环;2.手表;3.跑步机;4.智能运动鞋;5.耳机;6.智能运动鞋;7.智能运动Bra8.智能单车;9.智能跳绳10.智能背心11.脚环
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// Get DeviceType Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetDeviceType() int64 {
	return r._deviceType
}

// Set is DeviceName Setter
// 设备名（展示会用到）
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetDeviceName(_deviceName string) error {
	r._deviceName = _deviceName
	r.Set("device_name", _deviceName)
	return nil
}

// Get DeviceName Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetDeviceName() string {
	return r._deviceName
}

// Set is DeviceModel Setter
// 设备型号
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// Get DeviceModel Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

// Set is AverSpeed Setter
// 平均速度 单位：米/每分
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetAverSpeed(_averSpeed string) error {
	r._averSpeed = _averSpeed
	r.Set("aver_speed", _averSpeed)
	return nil
}

// Get AverSpeed Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetAverSpeed() string {
	return r._averSpeed
}

// Set is SubUnit Setter
// 二级运动量单位 定义：1.爬楼层数(跑步、健走、健身、登山);2.划水次数(游泳)
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetSubUnit(_subUnit int64) error {
	r._subUnit = _subUnit
	r.Set("sub_unit", _subUnit)
	return nil
}

// Get SubUnit Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetSubUnit() int64 {
	return r._subUnit
}

// Set is SubNum Setter
// 二级运动量
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetSubNum(_subNum string) error {
	r._subNum = _subNum
	r.Set("sub_num", _subNum)
	return nil
}

// Get SubNum Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetSubNum() string {
	return r._subNum
}

// Set is AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// Get AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// Set is AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// Get AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// Set is Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// Get Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetAliuid() string {
	return r._aliuid
}

// Set is Source Setter
// 业务来源二级分类（中英文）
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// Get Source Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetSource() string {
	return r._source
}

// Set is DataId Setter
// 三方运动数据主键id（数据唯一标记，去重使用）
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetDataId(_dataId string) error {
	r._dataId = _dataId
	r.Set("data_id", _dataId)
	return nil
}

// Get DataId Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetDataId() string {
	return r._dataId
}

// Set is SportsClass Setter
// 运动类型一级分类 定义：1-跑步;2-健走;3-骑行;4-游泳;5-健身;6-篮球;7-足球;8-羽毛球;9-排球;10-乒乓球;11-瑜伽;12-电竞;13-登山;16-椭圆机;19-健身操;20-太极;
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetSportsClass(_sportsClass int64) error {
	r._sportsClass = _sportsClass
	r.Set("sports_class", _sportsClass)
	return nil
}

// Get SportsClass Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetSportsClass() int64 {
	return r._sportsClass
}

// Set is SportsType Setter
// 运动类型二级分类 定义： 1001-室内跑步;1002-室外跑步;2001-室内健走;2002-室外健走;3001-室内骑行;3002-室外骑行;4001-室内游泳;4002-户外游泳
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetSportsType(_sportsType int64) error {
	r._sportsType = _sportsType
	r.Set("sports_type", _sportsType)
	return nil
}

// Get SportsType Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetSportsType() int64 {
	return r._sportsType
}

// Set is Num Setter
// 运动量
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetNum(_num string) error {
	r._num = _num
	r.Set("num", _num)
	return nil
}

// Get Num Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetNum() string {
	return r._num
}

// Set is Unit Setter
// 运动量单位 1.步数(跑步、健走、椭圆机、登山);2.趟数(游泳);3.平均踏频(骑行);
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetUnit(_unit int64) error {
	r._unit = _unit
	r.Set("unit", _unit)
	return nil
}

// Get Unit Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetUnit() int64 {
	return r._unit
}

// Set is Calorie Setter
// 运动消耗卡路里 单位：卡
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetCalorie(_calorie string) error {
	r._calorie = _calorie
	r.Set("calorie", _calorie)
	return nil
}

// Get Calorie Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetCalorie() string {
	return r._calorie
}

// Set is Match Setter
// 所属赛事（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetMatch(_match string) error {
	r._match = _match
	r.Set("match", _match)
	return nil
}

// Get Match Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetMatch() string {
	return r._match
}

// Set is Distance Setter
// 运动距离 单位:米（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetDistance(_distance string) error {
	r._distance = _distance
	r.Set("distance", _distance)
	return nil
}

// Get Distance Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetDistance() string {
	return r._distance
}

// Set is Time Setter
// 运动时长 单位:秒（有就传，阿里体育较依赖此字段）
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetTime(_time string) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// Get Time Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetTime() string {
	return r._time
}

// Set is Country Setter
// 国家(中英文/运动发生地点，如有就传)
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// Get Country Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetCountry() string {
	return r._country
}

// Set is Province Setter
// 省份(中英文/运动发生地点，如有就传)
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// Get Province Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetProvince() string {
	return r._province
}

// Set is City Setter
// 城市(中英文/运动发生地点，如有就传)
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// Get City Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetCity() string {
	return r._city
}

// Set is StartPoint Setter
// 开始运动地点经纬度，格式：1,2 (1为经度，2为纬度)
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetStartPoint(_startPoint string) error {
	r._startPoint = _startPoint
	r.Set("start_point", _startPoint)
	return nil
}

// Get StartPoint Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetStartPoint() string {
	return r._startPoint
}

// Set is EndPoint Setter
// 结束运动地点经纬度，格式：1,2 (1为经度，2为纬度)
func (r *AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) SetEndPoint(_endPoint string) error {
	r._endPoint = _endPoint
	r.Set("end_point", _endPoint)
	return nil
}

// Get EndPoint Getter
func (r AlibabaAlisportsDataSportsSyncsportsdataAPIRequest) GetEndPoint() string {
	return r._endPoint
}
