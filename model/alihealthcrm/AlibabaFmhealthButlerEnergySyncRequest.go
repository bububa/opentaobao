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
    _userId   int64
    // 每日已消耗能量
    _value   *BigDecimal
    // “S”- s点 “CAL”- 卡路里
    _energyType   string
    // 每日可消耗能量
    _target   *BigDecimal
    // 每日运动消耗能量值
    _sport   *BigDecimal
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
func (r *AlibabaFmhealthButlerEnergySyncRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaFmhealthButlerEnergySyncRequest) GetUserId() int64 {
    return r._userId
}
// Value Setter
// 每日已消耗能量
func (r *AlibabaFmhealthButlerEnergySyncRequest) SetValue(_value *BigDecimal) error {
    r._value = _value
    r.Set("value", _value)
    return nil
}

// Value Getter
func (r AlibabaFmhealthButlerEnergySyncRequest) GetValue() *BigDecimal {
    return r._value
}
// EnergyType Setter
// “S”- s点 “CAL”- 卡路里
func (r *AlibabaFmhealthButlerEnergySyncRequest) SetEnergyType(_energyType string) error {
    r._energyType = _energyType
    r.Set("energy_type", _energyType)
    return nil
}

// EnergyType Getter
func (r AlibabaFmhealthButlerEnergySyncRequest) GetEnergyType() string {
    return r._energyType
}
// Target Setter
// 每日可消耗能量
func (r *AlibabaFmhealthButlerEnergySyncRequest) SetTarget(_target *BigDecimal) error {
    r._target = _target
    r.Set("target", _target)
    return nil
}

// Target Getter
func (r AlibabaFmhealthButlerEnergySyncRequest) GetTarget() *BigDecimal {
    return r._target
}
// Sport Setter
// 每日运动消耗能量值
func (r *AlibabaFmhealthButlerEnergySyncRequest) SetSport(_sport *BigDecimal) error {
    r._sport = _sport
    r.Set("sport", _sport)
    return nil
}

// Sport Getter
func (r AlibabaFmhealthButlerEnergySyncRequest) GetSport() *BigDecimal {
    return r._sport
}
