package gameact

import (
	"net/url"
	"sync"

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
	// 时间戳
	_sequenceId int64
}

// NewTaobaoDeActivityLuckydrawRequest 初始化TaobaoDeActivityLuckydrawAPIRequest对象
func NewTaobaoDeActivityLuckydrawRequest() *TaobaoDeActivityLuckydrawAPIRequest {
	return &TaobaoDeActivityLuckydrawAPIRequest{
		Params: model.NewParams(11),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDeActivityLuckydrawAPIRequest) Reset() {
	r._eventKey = ""
	r._accountId = ""
	r._machineId = ""
	r._confirmKey = ""
	r._behaviorKey = ""
	r._channel = ""
	r._market = ""
	r._deviceModel = ""
	r._distribChannel = ""
	r._uuid = ""
	r._sequenceId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDeActivityLuckydrawAPIRequest) GetApiMethodName() string {
	return "taobao.de.activity.luckydraw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDeActivityLuckydrawAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDeActivityLuckydrawAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEventKey is EventKey Setter
// 运营和cp约定的事件唯一标示
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetEventKey(_eventKey string) error {
	r._eventKey = _eventKey
	r.Set("event_key", _eventKey)
	return nil
}

// GetEventKey EventKey Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetEventKey() string {
	return r._eventKey
}

// SetAccountId is AccountId Setter
// 用户的串ID
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetMachineId is MachineId Setter
// 机器设备号
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetMachineId(_machineId string) error {
	r._machineId = _machineId
	r.Set("machine_id", _machineId)
	return nil
}

// GetMachineId MachineId Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetMachineId() string {
	return r._machineId
}

// SetConfirmKey is ConfirmKey Setter
// 确认签名key
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetConfirmKey(_confirmKey string) error {
	r._confirmKey = _confirmKey
	r.Set("confirm_key", _confirmKey)
	return nil
}

// GetConfirmKey ConfirmKey Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetConfirmKey() string {
	return r._confirmKey
}

// SetBehaviorKey is BehaviorKey Setter
// 行为Key
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetBehaviorKey(_behaviorKey string) error {
	r._behaviorKey = _behaviorKey
	r.Set("behavior_key", _behaviorKey)
	return nil
}

// GetBehaviorKey BehaviorKey Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetBehaviorKey() string {
	return r._behaviorKey
}

// SetChannel is Channel Setter
// 渠道
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetChannel() string {
	return r._channel
}

// SetMarket is Market Setter
// 使用市场
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetMarket(_market string) error {
	r._market = _market
	r.Set("market", _market)
	return nil
}

// GetMarket Market Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetMarket() string {
	return r._market
}

// SetDeviceModel is DeviceModel Setter
// 盒型号
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// GetDeviceModel DeviceModel Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

// SetDistribChannel is DistribChannel Setter
// 魔盒分发渠道
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetDistribChannel(_distribChannel string) error {
	r._distribChannel = _distribChannel
	r.Set("distrib_channel", _distribChannel)
	return nil
}

// GetDistribChannel DistribChannel Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetDistribChannel() string {
	return r._distribChannel
}

// SetUuid is Uuid Setter
// 魔盒UUID
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetUuid() string {
	return r._uuid
}

// SetSequenceId is SequenceId Setter
// 时间戳
func (r *TaobaoDeActivityLuckydrawAPIRequest) SetSequenceId(_sequenceId int64) error {
	r._sequenceId = _sequenceId
	r.Set("sequence_id", _sequenceId)
	return nil
}

// GetSequenceId SequenceId Getter
func (r TaobaoDeActivityLuckydrawAPIRequest) GetSequenceId() int64 {
	return r._sequenceId
}

var poolTaobaoDeActivityLuckydrawAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDeActivityLuckydrawRequest()
	},
}

// GetTaobaoDeActivityLuckydrawRequest 从 sync.Pool 获取 TaobaoDeActivityLuckydrawAPIRequest
func GetTaobaoDeActivityLuckydrawAPIRequest() *TaobaoDeActivityLuckydrawAPIRequest {
	return poolTaobaoDeActivityLuckydrawAPIRequest.Get().(*TaobaoDeActivityLuckydrawAPIRequest)
}

// ReleaseTaobaoDeActivityLuckydrawAPIRequest 将 TaobaoDeActivityLuckydrawAPIRequest 放入 sync.Pool
func ReleaseTaobaoDeActivityLuckydrawAPIRequest(v *TaobaoDeActivityLuckydrawAPIRequest) {
	v.Reset()
	poolTaobaoDeActivityLuckydrawAPIRequest.Put(v)
}
