package gameact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeActivityLuckydrawAPIRequest 抽奖 API请求
// taobao.de.activity.luckydraw
//
// 用于激励平台对外提供抽奖功能，包括但不限于集分宝、红包、宝点、淘金币、淘彩票等
type TaobaoDeActivityLuckydrawAPIRequest struct {
	model.Params
	// 运营和cp约定的事件唯一标示
	_eventKey string
	// 时间戳
	_sequenceId int64
	// 用户的串ID
	_accountId string
	// 机器设备号
	_machineId string
	// 确认签名key
	_confirmKey string
	// 行为Key
	_behaviorKey string
	// 渠道
	_channel string
	// 使用市场
	_market string
	// 盒型号
	_deviceModel string
	// 魔盒分发渠道
	_distribChannel string
	// 魔盒UUID
	_uuid string
}

// NewTaobaoDeActivityLuckydrawRequest 初始化TaobaoDeActivityLuckydrawAPIRequest对象
func NewTaobaoDeActivityLuckydrawRequest() *TaobaoDeActivityLuckydrawAPIRequest {
	return &TaobaoDeActivityLuckydrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDeActivityLuckydrawAPIRequest) GetApiMethodName() string {
	return "taobao.de.activity.luckydraw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDeActivityLuckydrawAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EventKey Setter
// 运营和cp约定的事件唯一标示
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetEventKey(_eventKey string) error {
	r._eventKey = _eventKey
	r.Set("event_key", _eventKey)
	return nil
}

// Get EventKey Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetEventKey() string {
	return r._eventKey
}

// Set is SequenceId Setter
// 时间戳
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetSequenceId(_sequenceId int64) error {
	r._sequenceId = _sequenceId
	r.Set("sequence_id", _sequenceId)
	return nil
}

// Get SequenceId Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetSequenceId() int64 {
	return r._sequenceId
}

// Set is AccountId Setter
// 用户的串ID
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// Get AccountId Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetAccountId() string {
	return r._accountId
}

// Set is MachineId Setter
// 机器设备号
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetMachineId(_machineId string) error {
	r._machineId = _machineId
	r.Set("machine_id", _machineId)
	return nil
}

// Get MachineId Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetMachineId() string {
	return r._machineId
}

// Set is ConfirmKey Setter
// 确认签名key
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetConfirmKey(_confirmKey string) error {
	r._confirmKey = _confirmKey
	r.Set("confirm_key", _confirmKey)
	return nil
}

// Get ConfirmKey Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetConfirmKey() string {
	return r._confirmKey
}

// Set is BehaviorKey Setter
// 行为Key
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetBehaviorKey(_behaviorKey string) error {
	r._behaviorKey = _behaviorKey
	r.Set("behavior_key", _behaviorKey)
	return nil
}

// Get BehaviorKey Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetBehaviorKey() string {
	return r._behaviorKey
}

// Set is Channel Setter
// 渠道
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// Get Channel Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetChannel() string {
	return r._channel
}

// Set is Market Setter
// 使用市场
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetMarket(_market string) error {
	r._market = _market
	r.Set("market", _market)
	return nil
}

// Get Market Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetMarket() string {
	return r._market
}

// Set is DeviceModel Setter
// 盒型号
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// Get DeviceModel Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

// Set is DistribChannel Setter
// 魔盒分发渠道
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetDistribChannel(_distribChannel string) error {
	r._distribChannel = _distribChannel
	r.Set("distrib_channel", _distribChannel)
	return nil
}

// Get DistribChannel Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetDistribChannel() string {
	return r._distribChannel
}

// Set is Uuid Setter
// 魔盒UUID
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetUuid() string {
	return r._uuid
}
