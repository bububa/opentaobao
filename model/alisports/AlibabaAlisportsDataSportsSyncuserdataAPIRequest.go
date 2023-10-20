package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportsdatasportssyncuserdataAPIRequest 阿里体育数据中心用户个人信息同步接口 API请求
// alibaba.alisports.data.sports.syncuserdata
//
// 阿里体育数据中心用户个人信息同步接口
type AlibabaalisportsdatasportssyncuserdataAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 基础代谢率
	_metabolize string
	// 蛋白质含量
	_protein string
	// 骨量
	_bone string
	// 水分率
	_water string
	// 肌肉率
	_muscle string
	// 体脂率
	_fat string
	// 体重，单位kg
	_weight string
	// 身高，单位m
	_height string
	// 三方主键id，唯一标识数据
	_dataId string
	// 阿里体育用户id
	_aliuid string
	// 接口签名
	_alispSign string
	// 时间戳精确到秒
	_alispTime string
	// 日期 格式：y-m-d h:i:s
	_time string
	// 静息心率，单位：次/每分
	_heartRate int64
	// 年龄
	_age int64
}

// NewAlibabaalisportsdatasportssyncuserdataRequest 初始化AlibabaalisportsdatasportssyncuserdataAPIRequest对象
func NewAlibabaalisportsdatasportssyncuserdataRequest() *AlibabaalisportsdatasportssyncuserdataAPIRequest {
	return &AlibabaalisportsdatasportssyncuserdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.data.sports.syncuserdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 应用appkey
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetMetabolize is Metabolize Setter
// 基础代谢率
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetMetabolize(_metabolize string) error {
	r._metabolize = _metabolize
	r.Set("metabolize", _metabolize)
	return nil
}

// GetMetabolize Metabolize Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetMetabolize() string {
	return r._metabolize
}

// SetProtein is Protein Setter
// 蛋白质含量
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetProtein(_protein string) error {
	r._protein = _protein
	r.Set("protein", _protein)
	return nil
}

// GetProtein Protein Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetProtein() string {
	return r._protein
}

// SetBone is Bone Setter
// 骨量
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetBone(_bone string) error {
	r._bone = _bone
	r.Set("bone", _bone)
	return nil
}

// GetBone Bone Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetBone() string {
	return r._bone
}

// SetWater is Water Setter
// 水分率
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetWater(_water string) error {
	r._water = _water
	r.Set("water", _water)
	return nil
}

// GetWater Water Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetWater() string {
	return r._water
}

// SetMuscle is Muscle Setter
// 肌肉率
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetMuscle(_muscle string) error {
	r._muscle = _muscle
	r.Set("muscle", _muscle)
	return nil
}

// GetMuscle Muscle Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetMuscle() string {
	return r._muscle
}

// SetFat is Fat Setter
// 体脂率
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetFat(_fat string) error {
	r._fat = _fat
	r.Set("fat", _fat)
	return nil
}

// GetFat Fat Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetFat() string {
	return r._fat
}

// SetWeight is Weight Setter
// 体重，单位kg
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetWeight(_weight string) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetWeight() string {
	return r._weight
}

// SetHeight is Height Setter
// 身高，单位m
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetHeight(_height string) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// GetHeight Height Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetHeight() string {
	return r._height
}

// SetDataId is DataId Setter
// 三方主键id，唯一标识数据
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetDataId(_dataId string) error {
	r._dataId = _dataId
	r.Set("data_id", _dataId)
	return nil
}

// GetDataId DataId Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetDataId() string {
	return r._dataId
}

// SetAliuid is Aliuid Setter
// 阿里体育用户id
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// GetAliuid Aliuid Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetAliuid() string {
	return r._aliuid
}

// SetAlispSign is AlispSign Setter
// 接口签名
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetAlispTime is AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetTime is Time Setter
// 日期 格式：y-m-d h:i:s
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetTime(_time string) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// GetTime Time Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetTime() string {
	return r._time
}

// SetHeartRate is HeartRate Setter
// 静息心率，单位：次/每分
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetHeartRate(_heartRate int64) error {
	r._heartRate = _heartRate
	r.Set("heart_rate", _heartRate)
	return nil
}

// GetHeartRate HeartRate Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetHeartRate() int64 {
	return r._heartRate
}

// SetAge is Age Setter
// 年龄
func (r *AlibabaalisportsdatasportssyncuserdataAPIRequest) SetAge(_age int64) error {
	r._age = _age
	r.Set("age", _age)
	return nil
}

// GetAge Age Getter
func (r AlibabaalisportsdatasportssyncuserdataAPIRequest) GetAge() int64 {
	return r._age
}
