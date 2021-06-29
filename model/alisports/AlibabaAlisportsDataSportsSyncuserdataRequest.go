package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户个人信息同步接口 APIRequest
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

func NewAlibabaAlisportsDataSportsSyncuserdataRequest() *AlibabaAlisportsDataSportsSyncuserdataRequest{
    return &AlibabaAlisportsDataSportsSyncuserdataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetApiMethodName() string {
    return "alibaba.alisports.data.sports.syncuserdata"
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetAlispAppKey() string {
    return r.alispAppKey
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetMetabolize(metabolize string) error {
    r.metabolize = metabolize
    r.Set("metabolize", metabolize)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetMetabolize() string {
    return r.metabolize
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetProtein(protein string) error {
    r.protein = protein
    r.Set("protein", protein)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetProtein() string {
    return r.protein
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetBone(bone string) error {
    r.bone = bone
    r.Set("bone", bone)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetBone() string {
    return r.bone
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetWater(water string) error {
    r.water = water
    r.Set("water", water)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetWater() string {
    return r.water
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetMuscle(muscle string) error {
    r.muscle = muscle
    r.Set("muscle", muscle)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetMuscle() string {
    return r.muscle
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetFat(fat string) error {
    r.fat = fat
    r.Set("fat", fat)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetFat() string {
    return r.fat
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetHeartRate(heartRate int64) error {
    r.heartRate = heartRate
    r.Set("heart_rate", heartRate)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetHeartRate() int64 {
    return r.heartRate
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetWeight(weight string) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetWeight() string {
    return r.weight
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetHeight(height string) error {
    r.height = height
    r.Set("height", height)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetHeight() string {
    return r.height
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetAge(age int64) error {
    r.age = age
    r.Set("age", age)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetAge() int64 {
    return r.age
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetDataId(dataId string) error {
    r.dataId = dataId
    r.Set("data_id", dataId)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetDataId() string {
    return r.dataId
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetAliuid() string {
    return r.aliuid
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetAlispSign() string {
    return r.alispSign
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetAlispTime() string {
    return r.alispTime
}

func (r *AlibabaAlisportsDataSportsSyncuserdataRequest) SetTime(time string) error {
    r.time = time
    r.Set("time", time)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncuserdataRequest) GetTime() string {
    return r.time
}

