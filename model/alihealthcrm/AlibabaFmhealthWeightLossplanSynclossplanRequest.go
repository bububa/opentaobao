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
    _tpUserId   int64
    // 性别0女 1男
    _gender   int64
    // 生日
    _birthday   string
    // 身高170 即一米七
    _height   int64
    // 当前体重（今天的体重），单位kg
    _weight   string
    // 0创建减肥计划调用；1修改调用；
    _type   int64
    // 体年龄
    _bodyAge   int64
    // 完成时间
    _finishDate   string
    // 每周减重
    _lossPerWeek   string
    // 目标体重
    _weightGoal   string
    // 减重类型0保持 1减肥
    _lossLevel   int64
    // 每日可以摄入的标准总量
    _totalCalorie   int64
    // 减重计划开始时间
    _beginDate   string
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
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetTpUserId(_tpUserId int64) error {
    r._tpUserId = _tpUserId
    r.Set("tp_user_id", _tpUserId)
    return nil
}

// TpUserId Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetTpUserId() int64 {
    return r._tpUserId
}
// Gender Setter
// 性别0女 1男
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetGender(_gender int64) error {
    r._gender = _gender
    r.Set("gender", _gender)
    return nil
}

// Gender Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetGender() int64 {
    return r._gender
}
// Birthday Setter
// 生日
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetBirthday(_birthday string) error {
    r._birthday = _birthday
    r.Set("birthday", _birthday)
    return nil
}

// Birthday Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetBirthday() string {
    return r._birthday
}
// Height Setter
// 身高170 即一米七
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetHeight(_height int64) error {
    r._height = _height
    r.Set("height", _height)
    return nil
}

// Height Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetHeight() int64 {
    return r._height
}
// Weight Setter
// 当前体重（今天的体重），单位kg
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetWeight(_weight string) error {
    r._weight = _weight
    r.Set("weight", _weight)
    return nil
}

// Weight Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetWeight() string {
    return r._weight
}
// Type Setter
// 0创建减肥计划调用；1修改调用；
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetType() int64 {
    return r._type
}
// BodyAge Setter
// 体年龄
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetBodyAge(_bodyAge int64) error {
    r._bodyAge = _bodyAge
    r.Set("body_age", _bodyAge)
    return nil
}

// BodyAge Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetBodyAge() int64 {
    return r._bodyAge
}
// FinishDate Setter
// 完成时间
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetFinishDate(_finishDate string) error {
    r._finishDate = _finishDate
    r.Set("finish_date", _finishDate)
    return nil
}

// FinishDate Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetFinishDate() string {
    return r._finishDate
}
// LossPerWeek Setter
// 每周减重
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetLossPerWeek(_lossPerWeek string) error {
    r._lossPerWeek = _lossPerWeek
    r.Set("loss_per_week", _lossPerWeek)
    return nil
}

// LossPerWeek Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetLossPerWeek() string {
    return r._lossPerWeek
}
// WeightGoal Setter
// 目标体重
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetWeightGoal(_weightGoal string) error {
    r._weightGoal = _weightGoal
    r.Set("weight_goal", _weightGoal)
    return nil
}

// WeightGoal Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetWeightGoal() string {
    return r._weightGoal
}
// LossLevel Setter
// 减重类型0保持 1减肥
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetLossLevel(_lossLevel int64) error {
    r._lossLevel = _lossLevel
    r.Set("loss_level", _lossLevel)
    return nil
}

// LossLevel Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetLossLevel() int64 {
    return r._lossLevel
}
// TotalCalorie Setter
// 每日可以摄入的标准总量
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetTotalCalorie(_totalCalorie int64) error {
    r._totalCalorie = _totalCalorie
    r.Set("total_calorie", _totalCalorie)
    return nil
}

// TotalCalorie Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetTotalCalorie() int64 {
    return r._totalCalorie
}
// BeginDate Setter
// 减重计划开始时间
func (r *AlibabaFmhealthWeightLossplanSynclossplanRequest) SetBeginDate(_beginDate string) error {
    r._beginDate = _beginDate
    r.Set("begin_date", _beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanRequest) GetBeginDate() string {
    return r._beginDate
}
