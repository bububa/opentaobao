package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFmhealthButlerEnergySyncAPIRequest
同步用户消耗能量 API请求
alibaba.fmhealth.butler.energy.sync

同步用户消耗能量，用户消耗s点或卡路里后，同步给健康平台 */
type AlibabaFmhealthButlerEnergySyncAPIRequest struct {
	model.Params
	// 阿里用户id
	_userId int64
	// 每日已消耗能量
	_value *BigDecimal
	// “S”- s点 “CAL”- 卡路里
	_energyType string
	// 每日可消耗能量
	_target *BigDecimal
	// 每日运动消耗能量值
	_sport *BigDecimal
}

// NewAlibabaFmhealthButlerEnergySyncRequest 初始化AlibabaFmhealthButlerEnergySyncAPIRequest对象
func NewAlibabaFmhealthButlerEnergySyncRequest() *AlibabaFmhealthButlerEnergySyncAPIRequest {
	return &AlibabaFmhealthButlerEnergySyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetApiMethodName() string {
	return "alibaba.fmhealth.butler.energy.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserId Setter
// 阿里用户id
func (r *AlibabaFmhealthButlerEnergySyncAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetUserId() int64 {
	return r._userId
}

// Set is Value Setter
// 每日已消耗能量
func (r *AlibabaFmhealthButlerEnergySyncAPIRequest) SetValue(_value *BigDecimal) error {
	r._value = _value
	r.Set("value", _value)
	return nil
}

// Get Value Getter
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetValue() *BigDecimal {
	return r._value
}

// Set is EnergyType Setter
// “S”- s点 “CAL”- 卡路里
func (r *AlibabaFmhealthButlerEnergySyncAPIRequest) SetEnergyType(_energyType string) error {
	r._energyType = _energyType
	r.Set("energy_type", _energyType)
	return nil
}

// Get EnergyType Getter
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetEnergyType() string {
	return r._energyType
}

// Set is Target Setter
// 每日可消耗能量
func (r *AlibabaFmhealthButlerEnergySyncAPIRequest) SetTarget(_target *BigDecimal) error {
	r._target = _target
	r.Set("target", _target)
	return nil
}

// Get Target Getter
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetTarget() *BigDecimal {
	return r._target
}

// Set is Sport Setter
// 每日运动消耗能量值
func (r *AlibabaFmhealthButlerEnergySyncAPIRequest) SetSport(_sport *BigDecimal) error {
	r._sport = _sport
	r.Set("sport", _sport)
	return nil
}

// Get Sport Getter
func (r AlibabaFmhealthButlerEnergySyncAPIRequest) GetSport() *BigDecimal {
	return r._sport
}
