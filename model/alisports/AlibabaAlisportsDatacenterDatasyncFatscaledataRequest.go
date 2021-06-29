package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育接入体脂秤数据 APIRequest
alibaba.alisports.datacenter.datasync.fatscaledata

阿里体育数据中心接入体脂秤数据
*/
type AlibabaAlisportsDatacenterDatasyncFatscaledataRequest struct {
    model.Params

    // 阿里体育用户id
    aliuid   string 

    // 测量时间，秒级别时间戳
    time   int64 

    // 年龄
    age   int64 

    // 身高，单位cm
    height   int64 

    // 体重，单位kg
    weight   *BigDecimal 

    // bmi
    bmi   *BigDecimal 

    // 基础代谢率,单位卡
    basalMetabolicRate   *BigDecimal 

    // 去脂体重，单位kg
    fatFreeMass   *BigDecimal 

    // 体脂率，百分比12.4%，传12.4
    bodyFatRate   *BigDecimal 

    // 脂肪重量，单位kg
    fatMass   *BigDecimal 

    // 皮下脂肪率，百分比
    subcutaneousFatRate   *BigDecimal 

    // 内脏脂肪率，百分比
    visceralFatIndex   *BigDecimal 

    // 肌肉率，百分比
    muscleRate   *BigDecimal 

    // 肌肉重量，单位kg
    muscleMass   *BigDecimal 

    // 骨骼肌率，百分比
    skeletalMuscleRate   *BigDecimal 

    // 水分率，百分比
    moistureRate   *BigDecimal 

    // 蛋白质率，百分比
    proteinRate   *BigDecimal 

    // 骨量，单位kg
    boneMass   *BigDecimal 

    // 体重指数
    weightIndex   int64 

    // 身体年龄
    bodyAge   int64 

    // 设备类型：1.体脂秤，2智能手表，3智能手环
    deviceType   int64 

    // 设备名称
    deviceName   string 

    // 设备编号
    deviceModel   string 

    // 三方唯一id
    messageId   string 

}

func NewAlibabaAlisportsDatacenterDatasyncFatscaledataRequest() *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest{
    return &AlibabaAlisportsDatacenterDatasyncFatscaledataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetApiMethodName() string {
    return "alibaba.alisports.datacenter.datasync.fatscaledata"
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetAliuid() string {
    return r.aliuid
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetTime(time int64) error {
    r.time = time
    r.Set("time", time)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetTime() int64 {
    return r.time
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetAge(age int64) error {
    r.age = age
    r.Set("age", age)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetAge() int64 {
    return r.age
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetHeight(height int64) error {
    r.height = height
    r.Set("height", height)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetHeight() int64 {
    return r.height
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetWeight(weight *BigDecimal) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetWeight() *BigDecimal {
    return r.weight
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetBmi(bmi *BigDecimal) error {
    r.bmi = bmi
    r.Set("bmi", bmi)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetBmi() *BigDecimal {
    return r.bmi
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetBasalMetabolicRate(basalMetabolicRate *BigDecimal) error {
    r.basalMetabolicRate = basalMetabolicRate
    r.Set("basal_metabolic_rate", basalMetabolicRate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetBasalMetabolicRate() *BigDecimal {
    return r.basalMetabolicRate
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetFatFreeMass(fatFreeMass *BigDecimal) error {
    r.fatFreeMass = fatFreeMass
    r.Set("fat_free_mass", fatFreeMass)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetFatFreeMass() *BigDecimal {
    return r.fatFreeMass
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetBodyFatRate(bodyFatRate *BigDecimal) error {
    r.bodyFatRate = bodyFatRate
    r.Set("body_fat_rate", bodyFatRate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetBodyFatRate() *BigDecimal {
    return r.bodyFatRate
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetFatMass(fatMass *BigDecimal) error {
    r.fatMass = fatMass
    r.Set("fat_mass", fatMass)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetFatMass() *BigDecimal {
    return r.fatMass
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetSubcutaneousFatRate(subcutaneousFatRate *BigDecimal) error {
    r.subcutaneousFatRate = subcutaneousFatRate
    r.Set("subcutaneous_fat_rate", subcutaneousFatRate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetSubcutaneousFatRate() *BigDecimal {
    return r.subcutaneousFatRate
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetVisceralFatIndex(visceralFatIndex *BigDecimal) error {
    r.visceralFatIndex = visceralFatIndex
    r.Set("visceral_fat_index", visceralFatIndex)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetVisceralFatIndex() *BigDecimal {
    return r.visceralFatIndex
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetMuscleRate(muscleRate *BigDecimal) error {
    r.muscleRate = muscleRate
    r.Set("muscle_rate", muscleRate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetMuscleRate() *BigDecimal {
    return r.muscleRate
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetMuscleMass(muscleMass *BigDecimal) error {
    r.muscleMass = muscleMass
    r.Set("muscle_mass", muscleMass)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetMuscleMass() *BigDecimal {
    return r.muscleMass
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetSkeletalMuscleRate(skeletalMuscleRate *BigDecimal) error {
    r.skeletalMuscleRate = skeletalMuscleRate
    r.Set("skeletal_muscle_rate", skeletalMuscleRate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetSkeletalMuscleRate() *BigDecimal {
    return r.skeletalMuscleRate
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetMoistureRate(moistureRate *BigDecimal) error {
    r.moistureRate = moistureRate
    r.Set("moisture_rate", moistureRate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetMoistureRate() *BigDecimal {
    return r.moistureRate
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetProteinRate(proteinRate *BigDecimal) error {
    r.proteinRate = proteinRate
    r.Set("protein_rate", proteinRate)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetProteinRate() *BigDecimal {
    return r.proteinRate
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetBoneMass(boneMass *BigDecimal) error {
    r.boneMass = boneMass
    r.Set("bone_mass", boneMass)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetBoneMass() *BigDecimal {
    return r.boneMass
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetWeightIndex(weightIndex int64) error {
    r.weightIndex = weightIndex
    r.Set("weight_index", weightIndex)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetWeightIndex() int64 {
    return r.weightIndex
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetBodyAge(bodyAge int64) error {
    r.bodyAge = bodyAge
    r.Set("body_age", bodyAge)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetBodyAge() int64 {
    return r.bodyAge
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetDeviceType() int64 {
    return r.deviceType
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetDeviceName(deviceName string) error {
    r.deviceName = deviceName
    r.Set("device_name", deviceName)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetDeviceName() string {
    return r.deviceName
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetDeviceModel() string {
    return r.deviceModel
}

func (r *AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) SetMessageId(messageId string) error {
    r.messageId = messageId
    r.Set("message_id", messageId)
    return nil
}

func (r AlibabaAlisportsDatacenterDatasyncFatscaledataRequest) GetMessageId() string {
    return r.messageId
}

