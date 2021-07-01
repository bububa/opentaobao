package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest
阿里体育接入体脂秤数据 API请求
alibaba.alisports.datacenter.datasync.fatscaledata

阿里体育数据中心接入体脂秤数据 */
type AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest struct {
	model.Params
	// 阿里体育用户id
	_aliuid string
	// 测量时间，秒级别时间戳
	_time int64
	// 年龄
	_age int64
	// 身高，单位cm
	_height int64
	// 体重，单位kg
	_weight *BigDecimal
	// bmi
	_bmi *BigDecimal
	// 基础代谢率,单位卡
	_basalMetabolicRate *BigDecimal
	// 去脂体重，单位kg
	_fatFreeMass *BigDecimal
	// 体脂率，百分比12.4%，传12.4
	_bodyFatRate *BigDecimal
	// 脂肪重量，单位kg
	_fatMass *BigDecimal
	// 皮下脂肪率，百分比
	_subcutaneousFatRate *BigDecimal
	// 内脏脂肪率，百分比
	_visceralFatIndex *BigDecimal
	// 肌肉率，百分比
	_muscleRate *BigDecimal
	// 肌肉重量，单位kg
	_muscleMass *BigDecimal
	// 骨骼肌率，百分比
	_skeletalMuscleRate *BigDecimal
	// 水分率，百分比
	_moistureRate *BigDecimal
	// 蛋白质率，百分比
	_proteinRate *BigDecimal
	// 骨量，单位kg
	_boneMass *BigDecimal
	// 体重指数
	_weightIndex int64
	// 身体年龄
	_bodyAge int64
	// 设备类型：1.体脂秤，2智能手表，3智能手环
	_deviceType int64
	// 设备名称
	_deviceName string
	// 设备编号
	_deviceModel string
	// 三方唯一id
	_messageId string
}

// NewAlibabaAlisportsDatacenterDatasyncFatscaledataRequest 初始化AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest对象
func NewAlibabaAlisportsDatacenterDatasyncFatscaledataRequest() *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest {
	return &AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.datacenter.datasync.fatscaledata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// Get Aliuid Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetAliuid() string {
	return r._aliuid
}

// Set is Time Setter
// 测量时间，秒级别时间戳
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetTime(_time int64) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// Get Time Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetTime() int64 {
	return r._time
}

// Set is Age Setter
// 年龄
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetAge(_age int64) error {
	r._age = _age
	r.Set("age", _age)
	return nil
}

// Get Age Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetAge() int64 {
	return r._age
}

// Set is Height Setter
// 身高，单位cm
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// Get Height Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetHeight() int64 {
	return r._height
}

// Set is Weight Setter
// 体重，单位kg
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetWeight(_weight *BigDecimal) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// Get Weight Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetWeight() *BigDecimal {
	return r._weight
}

// Set is Bmi Setter
// bmi
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetBmi(_bmi *BigDecimal) error {
	r._bmi = _bmi
	r.Set("bmi", _bmi)
	return nil
}

// Get Bmi Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetBmi() *BigDecimal {
	return r._bmi
}

// Set is BasalMetabolicRate Setter
// 基础代谢率,单位卡
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetBasalMetabolicRate(_basalMetabolicRate *BigDecimal) error {
	r._basalMetabolicRate = _basalMetabolicRate
	r.Set("basal_metabolic_rate", _basalMetabolicRate)
	return nil
}

// Get BasalMetabolicRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetBasalMetabolicRate() *BigDecimal {
	return r._basalMetabolicRate
}

// Set is FatFreeMass Setter
// 去脂体重，单位kg
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetFatFreeMass(_fatFreeMass *BigDecimal) error {
	r._fatFreeMass = _fatFreeMass
	r.Set("fat_free_mass", _fatFreeMass)
	return nil
}

// Get FatFreeMass Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetFatFreeMass() *BigDecimal {
	return r._fatFreeMass
}

// Set is BodyFatRate Setter
// 体脂率，百分比12.4%，传12.4
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetBodyFatRate(_bodyFatRate *BigDecimal) error {
	r._bodyFatRate = _bodyFatRate
	r.Set("body_fat_rate", _bodyFatRate)
	return nil
}

// Get BodyFatRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetBodyFatRate() *BigDecimal {
	return r._bodyFatRate
}

// Set is FatMass Setter
// 脂肪重量，单位kg
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetFatMass(_fatMass *BigDecimal) error {
	r._fatMass = _fatMass
	r.Set("fat_mass", _fatMass)
	return nil
}

// Get FatMass Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetFatMass() *BigDecimal {
	return r._fatMass
}

// Set is SubcutaneousFatRate Setter
// 皮下脂肪率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetSubcutaneousFatRate(_subcutaneousFatRate *BigDecimal) error {
	r._subcutaneousFatRate = _subcutaneousFatRate
	r.Set("subcutaneous_fat_rate", _subcutaneousFatRate)
	return nil
}

// Get SubcutaneousFatRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetSubcutaneousFatRate() *BigDecimal {
	return r._subcutaneousFatRate
}

// Set is VisceralFatIndex Setter
// 内脏脂肪率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetVisceralFatIndex(_visceralFatIndex *BigDecimal) error {
	r._visceralFatIndex = _visceralFatIndex
	r.Set("visceral_fat_index", _visceralFatIndex)
	return nil
}

// Get VisceralFatIndex Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetVisceralFatIndex() *BigDecimal {
	return r._visceralFatIndex
}

// Set is MuscleRate Setter
// 肌肉率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetMuscleRate(_muscleRate *BigDecimal) error {
	r._muscleRate = _muscleRate
	r.Set("muscle_rate", _muscleRate)
	return nil
}

// Get MuscleRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetMuscleRate() *BigDecimal {
	return r._muscleRate
}

// Set is MuscleMass Setter
// 肌肉重量，单位kg
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetMuscleMass(_muscleMass *BigDecimal) error {
	r._muscleMass = _muscleMass
	r.Set("muscle_mass", _muscleMass)
	return nil
}

// Get MuscleMass Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetMuscleMass() *BigDecimal {
	return r._muscleMass
}

// Set is SkeletalMuscleRate Setter
// 骨骼肌率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetSkeletalMuscleRate(_skeletalMuscleRate *BigDecimal) error {
	r._skeletalMuscleRate = _skeletalMuscleRate
	r.Set("skeletal_muscle_rate", _skeletalMuscleRate)
	return nil
}

// Get SkeletalMuscleRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetSkeletalMuscleRate() *BigDecimal {
	return r._skeletalMuscleRate
}

// Set is MoistureRate Setter
// 水分率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetMoistureRate(_moistureRate *BigDecimal) error {
	r._moistureRate = _moistureRate
	r.Set("moisture_rate", _moistureRate)
	return nil
}

// Get MoistureRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetMoistureRate() *BigDecimal {
	return r._moistureRate
}

// Set is ProteinRate Setter
// 蛋白质率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetProteinRate(_proteinRate *BigDecimal) error {
	r._proteinRate = _proteinRate
	r.Set("protein_rate", _proteinRate)
	return nil
}

// Get ProteinRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetProteinRate() *BigDecimal {
	return r._proteinRate
}

// Set is BoneMass Setter
// 骨量，单位kg
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetBoneMass(_boneMass *BigDecimal) error {
	r._boneMass = _boneMass
	r.Set("bone_mass", _boneMass)
	return nil
}

// Get BoneMass Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetBoneMass() *BigDecimal {
	return r._boneMass
}

// Set is WeightIndex Setter
// 体重指数
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetWeightIndex(_weightIndex int64) error {
	r._weightIndex = _weightIndex
	r.Set("weight_index", _weightIndex)
	return nil
}

// Get WeightIndex Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetWeightIndex() int64 {
	return r._weightIndex
}

// Set is BodyAge Setter
// 身体年龄
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetBodyAge(_bodyAge int64) error {
	r._bodyAge = _bodyAge
	r.Set("body_age", _bodyAge)
	return nil
}

// Get BodyAge Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetBodyAge() int64 {
	return r._bodyAge
}

// Set is DeviceType Setter
// 设备类型：1.体脂秤，2智能手表，3智能手环
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// Get DeviceType Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetDeviceType() int64 {
	return r._deviceType
}

// Set is DeviceName Setter
// 设备名称
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetDeviceName(_deviceName string) error {
	r._deviceName = _deviceName
	r.Set("device_name", _deviceName)
	return nil
}

// Get DeviceName Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetDeviceName() string {
	return r._deviceName
}

// Set is DeviceModel Setter
// 设备编号
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// Get DeviceModel Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

// Set is MessageId Setter
// 三方唯一id
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) SetMessageId(_messageId string) error {
	r._messageId = _messageId
	r.Set("message_id", _messageId)
	return nil
}

// Get MessageId Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest) GetMessageId() string {
	return r._messageId
}
