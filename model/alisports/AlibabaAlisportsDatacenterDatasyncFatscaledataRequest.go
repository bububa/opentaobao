package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育接入体脂秤数据 API请求
alibaba.alisports.datacenter.datasync.fatscaledata

阿里体育数据中心接入体脂秤数据
*/
type AlibabaAlisportsDatacenterDatasyncFatscaledataRequest struct {
    model.Params
    // 阿里体育用户id
    _aliuid   string
    // 测量时间，秒级别时间戳
    _time   int64
    // 年龄
    _age   int64
    // 身高，单位cm
    _height   int64
    // 体重，单位kg
    _weight   *BigDecimal
    // bmi
    _bmi   *BigDecimal
    // 基础代谢率,单位卡
    _basalMetabolicRate   *BigDecimal
    // 去脂体重，单位kg
    _fatFreeMass   *BigDecimal
    // 体脂率，百分比12.4%，传12.4
    _bodyFatRate   *BigDecimal
    // 脂肪重量，单位kg
    _fatMass   *BigDecimal
    // 皮下脂肪率，百分比
    _subcutaneousFatRate   *BigDecimal
    // 内脏脂肪率，百分比
    _visceralFatIndex   *BigDecimal
    // 肌肉率，百分比
    _muscleRate   *BigDecimal
    // 肌肉重量，单位kg
    _muscleMass   *BigDecimal
    // 骨骼肌率，百分比
    _skeletalMuscleRate   *BigDecimal
    // 水分率，百分比
    _moistureRate   *BigDecimal
    // 蛋白质率，百分比
    _proteinRate   *BigDecimal
    // 骨量，单位kg
    _boneMass   *BigDecimal
    // 体重指数
    _weightIndex   int64
    // 身体年龄
    _bodyAge   int64
    // 设备类型：1.体脂秤，2智能手表，3智能手环
    _deviceType   int64
    // 设备名称
    _deviceName   string
    // 设备编号
    _deviceModel   string
    // 三方唯一id
    _messageId   string
}

// 初始化AlibabaAlisportsDatacenterDatasyncFatscaledataRequest对象
func NewAlibabaAlisportsDatacenterDatasyncFatscaledataRequest() *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest{
    return &AlibabaAlisportsDatacenterDatasyncFatscaledataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetApiMethodName() string {
    return "alibaba.alisports.datacenter.datasync.fatscaledata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetAliuid(_aliuid string) error {
    r._aliuid = _aliuid
    r.Set("aliuid", _aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetAliuid() string {
    return r._aliuid
}
// Time Setter
// 测量时间，秒级别时间戳
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetTime(_time int64) error {
    r._time = _time
    r.Set("time", _time)
    return nil
}

// Time Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetTime() int64 {
    return r._time
}
// Age Setter
// 年龄
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetAge(_age int64) error {
    r._age = _age
    r.Set("age", _age)
    return nil
}

// Age Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetAge() int64 {
    return r._age
}
// Height Setter
// 身高，单位cm
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetHeight(_height int64) error {
    r._height = _height
    r.Set("height", _height)
    return nil
}

// Height Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetHeight() int64 {
    return r._height
}
// Weight Setter
// 体重，单位kg
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetWeight(_weight *BigDecimal) error {
    r._weight = _weight
    r.Set("weight", _weight)
    return nil
}

// Weight Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetWeight() *BigDecimal {
    return r._weight
}
// Bmi Setter
// bmi
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetBmi(_bmi *BigDecimal) error {
    r._bmi = _bmi
    r.Set("bmi", _bmi)
    return nil
}

// Bmi Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetBmi() *BigDecimal {
    return r._bmi
}
// BasalMetabolicRate Setter
// 基础代谢率,单位卡
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetBasalMetabolicRate(_basalMetabolicRate *BigDecimal) error {
    r._basalMetabolicRate = _basalMetabolicRate
    r.Set("basal_metabolic_rate", _basalMetabolicRate)
    return nil
}

// BasalMetabolicRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetBasalMetabolicRate() *BigDecimal {
    return r._basalMetabolicRate
}
// FatFreeMass Setter
// 去脂体重，单位kg
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetFatFreeMass(_fatFreeMass *BigDecimal) error {
    r._fatFreeMass = _fatFreeMass
    r.Set("fat_free_mass", _fatFreeMass)
    return nil
}

// FatFreeMass Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetFatFreeMass() *BigDecimal {
    return r._fatFreeMass
}
// BodyFatRate Setter
// 体脂率，百分比12.4%，传12.4
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetBodyFatRate(_bodyFatRate *BigDecimal) error {
    r._bodyFatRate = _bodyFatRate
    r.Set("body_fat_rate", _bodyFatRate)
    return nil
}

// BodyFatRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetBodyFatRate() *BigDecimal {
    return r._bodyFatRate
}
// FatMass Setter
// 脂肪重量，单位kg
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetFatMass(_fatMass *BigDecimal) error {
    r._fatMass = _fatMass
    r.Set("fat_mass", _fatMass)
    return nil
}

// FatMass Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetFatMass() *BigDecimal {
    return r._fatMass
}
// SubcutaneousFatRate Setter
// 皮下脂肪率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetSubcutaneousFatRate(_subcutaneousFatRate *BigDecimal) error {
    r._subcutaneousFatRate = _subcutaneousFatRate
    r.Set("subcutaneous_fat_rate", _subcutaneousFatRate)
    return nil
}

// SubcutaneousFatRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetSubcutaneousFatRate() *BigDecimal {
    return r._subcutaneousFatRate
}
// VisceralFatIndex Setter
// 内脏脂肪率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetVisceralFatIndex(_visceralFatIndex *BigDecimal) error {
    r._visceralFatIndex = _visceralFatIndex
    r.Set("visceral_fat_index", _visceralFatIndex)
    return nil
}

// VisceralFatIndex Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetVisceralFatIndex() *BigDecimal {
    return r._visceralFatIndex
}
// MuscleRate Setter
// 肌肉率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetMuscleRate(_muscleRate *BigDecimal) error {
    r._muscleRate = _muscleRate
    r.Set("muscle_rate", _muscleRate)
    return nil
}

// MuscleRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetMuscleRate() *BigDecimal {
    return r._muscleRate
}
// MuscleMass Setter
// 肌肉重量，单位kg
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetMuscleMass(_muscleMass *BigDecimal) error {
    r._muscleMass = _muscleMass
    r.Set("muscle_mass", _muscleMass)
    return nil
}

// MuscleMass Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetMuscleMass() *BigDecimal {
    return r._muscleMass
}
// SkeletalMuscleRate Setter
// 骨骼肌率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetSkeletalMuscleRate(_skeletalMuscleRate *BigDecimal) error {
    r._skeletalMuscleRate = _skeletalMuscleRate
    r.Set("skeletal_muscle_rate", _skeletalMuscleRate)
    return nil
}

// SkeletalMuscleRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetSkeletalMuscleRate() *BigDecimal {
    return r._skeletalMuscleRate
}
// MoistureRate Setter
// 水分率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetMoistureRate(_moistureRate *BigDecimal) error {
    r._moistureRate = _moistureRate
    r.Set("moisture_rate", _moistureRate)
    return nil
}

// MoistureRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetMoistureRate() *BigDecimal {
    return r._moistureRate
}
// ProteinRate Setter
// 蛋白质率，百分比
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetProteinRate(_proteinRate *BigDecimal) error {
    r._proteinRate = _proteinRate
    r.Set("protein_rate", _proteinRate)
    return nil
}

// ProteinRate Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetProteinRate() *BigDecimal {
    return r._proteinRate
}
// BoneMass Setter
// 骨量，单位kg
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetBoneMass(_boneMass *BigDecimal) error {
    r._boneMass = _boneMass
    r.Set("bone_mass", _boneMass)
    return nil
}

// BoneMass Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetBoneMass() *BigDecimal {
    return r._boneMass
}
// WeightIndex Setter
// 体重指数
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetWeightIndex(_weightIndex int64) error {
    r._weightIndex = _weightIndex
    r.Set("weight_index", _weightIndex)
    return nil
}

// WeightIndex Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetWeightIndex() int64 {
    return r._weightIndex
}
// BodyAge Setter
// 身体年龄
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetBodyAge(_bodyAge int64) error {
    r._bodyAge = _bodyAge
    r.Set("body_age", _bodyAge)
    return nil
}

// BodyAge Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetBodyAge() int64 {
    return r._bodyAge
}
// DeviceType Setter
// 设备类型：1.体脂秤，2智能手表，3智能手环
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetDeviceType(_deviceType int64) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetDeviceType() int64 {
    return r._deviceType
}
// DeviceName Setter
// 设备名称
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetDeviceName(_deviceName string) error {
    r._deviceName = _deviceName
    r.Set("device_name", _deviceName)
    return nil
}

// DeviceName Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetDeviceName() string {
    return r._deviceName
}
// DeviceModel Setter
// 设备编号
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetDeviceModel(_deviceModel string) error {
    r._deviceModel = _deviceModel
    r.Set("device_model", _deviceModel)
    return nil
}

// DeviceModel Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetDeviceModel() string {
    return r._deviceModel
}
// MessageId Setter
// 三方唯一id
func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetMessageId(_messageId string) error {
    r._messageId = _messageId
    r.Set("message_id", _messageId)
    return nil
}

// MessageId Getter
func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetMessageId() string {
    return r._messageId
}
