package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户个人信息同步接口 API请求
alibaba.alisports.data.sports.syncuserdata

阿里体育数据中心用户个人信息同步接口
*/
type AlibabaAlisportsDataSportsSyncuserdataRequest struct {
    model.Params
    // 应用appkey
    alispAppKey   string
    // 基础代谢率
    metabolize   string
    // 蛋白质含量
    protein   string
    // 骨量
    bone   string
    // 水分率
    water   string
    // 肌肉率
    muscle   string
    // 体脂率
    fat   string
    // 静息心率，单位：次/每分
    heartRate   int64
    // 体重，单位kg
    weight   string
    // 身高，单位m
    height   string
    // 年龄
    age   int64
    // 三方主键id，唯一标识数据
    dataId   string
    // 阿里体育用户id
    aliuid   string
    // 接口签名
    alispSign   string
    // 时间戳精确到秒
    alispTime   string
    // 日期 格式：y-m-d h:i:s
    time   string
}

// 初始化AlibabaAlisportsDataSportsSyncuserdataRequest对象
func NewAlibabaAlisportsDataSportsSyncuserdataRequest() *AlibabaAlisportsDataSportsSyncuserdataRequest{
    return &AlibabaAlisportsDataSportsSyncuserdataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetApiMethodName() string {
    return "alibaba.alisports.data.sports.syncuserdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 应用appkey
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetAlispAppKey() string {
    return r.alispAppKey
}
// Metabolize Setter
// 基础代谢率
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetMetabolize(metabolize string) error {
    r.metabolize = metabolize
    r.Set("metabolize", metabolize)
    return nil
}

// Metabolize Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetMetabolize() string {
    return r.metabolize
}
// Protein Setter
// 蛋白质含量
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetProtein(protein string) error {
    r.protein = protein
    r.Set("protein", protein)
    return nil
}

// Protein Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetProtein() string {
    return r.protein
}
// Bone Setter
// 骨量
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetBone(bone string) error {
    r.bone = bone
    r.Set("bone", bone)
    return nil
}

// Bone Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetBone() string {
    return r.bone
}
// Water Setter
// 水分率
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetWater(water string) error {
    r.water = water
    r.Set("water", water)
    return nil
}

// Water Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetWater() string {
    return r.water
}
// Muscle Setter
// 肌肉率
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetMuscle(muscle string) error {
    r.muscle = muscle
    r.Set("muscle", muscle)
    return nil
}

// Muscle Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetMuscle() string {
    return r.muscle
}
// Fat Setter
// 体脂率
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetFat(fat string) error {
    r.fat = fat
    r.Set("fat", fat)
    return nil
}

// Fat Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetFat() string {
    return r.fat
}
// HeartRate Setter
// 静息心率，单位：次/每分
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetHeartRate(heartRate int64) error {
    r.heartRate = heartRate
    r.Set("heart_rate", heartRate)
    return nil
}

// HeartRate Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetHeartRate() int64 {
    return r.heartRate
}
// Weight Setter
// 体重，单位kg
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetWeight(weight string) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

// Weight Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetWeight() string {
    return r.weight
}
// Height Setter
// 身高，单位m
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetHeight(height string) error {
    r.height = height
    r.Set("height", height)
    return nil
}

// Height Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetHeight() string {
    return r.height
}
// Age Setter
// 年龄
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetAge(age int64) error {
    r.age = age
    r.Set("age", age)
    return nil
}

// Age Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetAge() int64 {
    return r.age
}
// DataId Setter
// 三方主键id，唯一标识数据
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetDataId(dataId string) error {
    r.dataId = dataId
    r.Set("data_id", dataId)
    return nil
}

// DataId Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetDataId() string {
    return r.dataId
}
// Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetAliuid() string {
    return r.aliuid
}
// AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetAlispSign() string {
    return r.alispSign
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetAlispTime() string {
    return r.alispTime
}
// Time Setter
// 日期 格式：y-m-d h:i:s
func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetTime(time string) error {
    r.time = time
    r.Set("time", time)
    return nil
}

// Time Getter
func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetTime() string {
    return r.time
}
