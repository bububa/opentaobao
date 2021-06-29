package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步用户消耗能量 API请求
alibaba.fmhealth.butler.energy.sync

同步用户消耗能量，用户消耗s点或卡路里后，同步给健康平台
*/
type AlibabaFmhealthButlerEnergySyncRequest struct {
    model.Params
    // 阿里用户id
    userId   int64
    // 每日已消耗能量
    value   *BigDecimal
    // “S”- s点 “CAL”- 卡路里
    energyType   string
    // 每日可消耗能量
    target   *BigDecimal
    // 每日运动消耗能量值
    sport   *BigDecimal
}

// 初始化AlibabaFmhealthButlerEnergySyncRequest对象
func NewAlibabaFmhealthButlerEnergySyncRequest() *AlibabaFmhealthButlerEnergySyncRequest{
    return &AlibabaFmhealthButlerEnergySyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFmhealthButlerEnergySyncRequest) GetApiMethodName() string {
    return "alibaba.fmhealth.butler.energy.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFmhealthButlerEnergySyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 阿里用户id
func (r *AlibabaFmhealthButlerEnergySyncRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaFmhealthButlerEnergySyncRequest) GetUserId() int64 {
    return r.userId
}
// Value Setter
// 每日已消耗能量
func (r *AlibabaFmhealthButlerEnergySyncRequest) SetValue(value *BigDecimal) error {
    r.value = value
    r.Set("value", value)
    return nil
}

// Value Getter
func (r AlibabaFmhealthButlerEnergySyncRequest) GetValue() *BigDecimal {
    return r.value
}
// EnergyType Setter
// “S”- s点 “CAL”- 卡路里
func (r *AlibabaFmhealthButlerEnergySyncRequest) SetEnergyType(energyType string) error {
    r.energyType = energyType
    r.Set("energy_type", energyType)
    return nil
}

// EnergyType Getter
func (r AlibabaFmhealthButlerEnergySyncRequest) GetEnergyType() string {
    return r.energyType
}
// Target Setter
// 每日可消耗能量
func (r *AlibabaFmhealthButlerEnergySyncRequest) SetTarget(target *BigDecimal) error {
    r.target = target
    r.Set("target", target)
    return nil
}

// Target Getter
func (r AlibabaFmhealthButlerEnergySyncRequest) GetTarget() *BigDecimal {
    return r.target
}
// Sport Setter
// 每日运动消耗能量值
func (r *AlibabaFmhealthButlerEnergySyncRequest) SetSport(sport *BigDecimal) error {
    r.sport = sport
    r.Set("sport", sport)
    return nil
}

// Sport Getter
func (r AlibabaFmhealthButlerEnergySyncRequest) GetSport() *BigDecimal {
    return r.sport
}
