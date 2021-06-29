package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
减重计划--同步减重计划 API请求
alibaba.fmhealth.weight.lossplan.synclossplan

减重计划--三方同步用户初始化减重计划给我们
*/
type AlibabaFmhealthWeightLossplanSynclossplanRequest struct {
    model.Params
    // 用户id
    tpUserId   int64
    // 性别0女 1男
    gender   int64
    // 生日
    birthday   string
    // 身高170 即一米七
    height   int64
    // 当前体重（今天的体重），单位kg
    weight   string
    // 0创建减肥计划调用；1修改调用；
    type   int64
    // 体年龄
    bodyAge   int64
    // 完成时间
    finishDate   string
    // 每周减重
    lossPerWeek   string
    // 目标体重
    weightGoal   string
    // 减重类型0保持 1减肥
    lossLevel   int64
    // 每日可以摄入的标准总量
    totalCalorie   int64
    // 减重计划开始时间
    beginDate   string
}

// 初始化AlibabaFmhealthWeightLossplanSynclossplanRequest对象
func NewAlibabaFmhealthWeightLossplanSynclossplanRequest() *AlibabaFmhealthWeightLossplanSynclossplanRequest{
    return &AlibabaFmhealthWeightLossplanSynclossplanRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetApiMethodName() string {
    return "alibaba.fmhealth.weight.lossplan.synclossplan"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TpUserId Setter
// 用户id
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetTpUserId(tpUserId int64) error {
    r.tpUserId = tpUserId
    r.Set("tp_user_id", tpUserId)
    return nil
}

// TpUserId Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetTpUserId() int64 {
    return r.tpUserId
}
// Gender Setter
// 性别0女 1男
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetGender(gender int64) error {
    r.gender = gender
    r.Set("gender", gender)
    return nil
}

// Gender Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetGender() int64 {
    return r.gender
}
// Birthday Setter
// 生日
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetBirthday(birthday string) error {
    r.birthday = birthday
    r.Set("birthday", birthday)
    return nil
}

// Birthday Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetBirthday() string {
    return r.birthday
}
// Height Setter
// 身高170 即一米七
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetHeight(height int64) error {
    r.height = height
    r.Set("height", height)
    return nil
}

// Height Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetHeight() int64 {
    return r.height
}
// Weight Setter
// 当前体重（今天的体重），单位kg
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetWeight(weight string) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

// Weight Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetWeight() string {
    return r.weight
}
// Type Setter
// 0创建减肥计划调用；1修改调用；
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetType() int64 {
    return r.type
}
// BodyAge Setter
// 体年龄
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetBodyAge(bodyAge int64) error {
    r.bodyAge = bodyAge
    r.Set("body_age", bodyAge)
    return nil
}

// BodyAge Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetBodyAge() int64 {
    return r.bodyAge
}
// FinishDate Setter
// 完成时间
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetFinishDate(finishDate string) error {
    r.finishDate = finishDate
    r.Set("finish_date", finishDate)
    return nil
}

// FinishDate Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetFinishDate() string {
    return r.finishDate
}
// LossPerWeek Setter
// 每周减重
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetLossPerWeek(lossPerWeek string) error {
    r.lossPerWeek = lossPerWeek
    r.Set("loss_per_week", lossPerWeek)
    return nil
}

// LossPerWeek Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetLossPerWeek() string {
    return r.lossPerWeek
}
// WeightGoal Setter
// 目标体重
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetWeightGoal(weightGoal string) error {
    r.weightGoal = weightGoal
    r.Set("weight_goal", weightGoal)
    return nil
}

// WeightGoal Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetWeightGoal() string {
    return r.weightGoal
}
// LossLevel Setter
// 减重类型0保持 1减肥
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetLossLevel(lossLevel int64) error {
    r.lossLevel = lossLevel
    r.Set("loss_level", lossLevel)
    return nil
}

// LossLevel Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetLossLevel() int64 {
    return r.lossLevel
}
// TotalCalorie Setter
// 每日可以摄入的标准总量
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetTotalCalorie(totalCalorie int64) error {
    r.totalCalorie = totalCalorie
    r.Set("total_calorie", totalCalorie)
    return nil
}

// TotalCalorie Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetTotalCalorie() int64 {
    return r.totalCalorie
}
// BeginDate Setter
// 减重计划开始时间
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetBeginDate() string {
    return r.beginDate
}
