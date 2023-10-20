package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// Alibabanlifeb2ctradestatusdriveAPIRequest b2c订单状态驱动 API请求
// alibaba.nlife.b2c.tradestatus.drive
//
// 用于驱动零售+订单状态
type Alibabanlifeb2ctradestatusdriveAPIRequest struct {
	model.Params
	// 零售门店在零售+平台的ID
	_storeId string
	// APP:是指线上销售应用，POS:是指现场收银应用
	_channel string
	// 对零售+为外部订单号，对业务方为订单号
	_outTradeNo string
	// 零售+平台订单号，和out_trade_no不能同时为空
	_tradeNo string
	// 接口类型：CONFIRM（收货）DELIVER（发货）
	_action string
	// 扩展参数 JSON格式
	_extendParams string
	// 货流信息
	_logisticsInfo *LogisticsInfo
}

// NewAlibabanlifeb2ctradestatusdriveRequest 初始化Alibabanlifeb2ctradestatusdriveAPIRequest对象
func NewAlibabanlifeb2ctradestatusdriveRequest() *Alibabanlifeb2ctradestatusdriveAPIRequest {
	return &Alibabanlifeb2ctradestatusdriveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r Alibabanlifeb2ctradestatusdriveAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.tradestatus.drive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r Alibabanlifeb2ctradestatusdriveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r Alibabanlifeb2ctradestatusdriveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 零售门店在零售+平台的ID
func (r *Alibabanlifeb2ctradestatusdriveAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r Alibabanlifeb2ctradestatusdriveAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetChannel is Channel Setter
// APP:是指线上销售应用，POS:是指现场收银应用
func (r *Alibabanlifeb2ctradestatusdriveAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r Alibabanlifeb2ctradestatusdriveAPIRequest) GetChannel() string {
	return r._channel
}

// SetOutTradeNo is OutTradeNo Setter
// 对零售+为外部订单号，对业务方为订单号
func (r *Alibabanlifeb2ctradestatusdriveAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r Alibabanlifeb2ctradestatusdriveAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// SetTradeNo is TradeNo Setter
// 零售+平台订单号，和out_trade_no不能同时为空
func (r *Alibabanlifeb2ctradestatusdriveAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r Alibabanlifeb2ctradestatusdriveAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetAction is Action Setter
// 接口类型：CONFIRM（收货）DELIVER（发货）
func (r *Alibabanlifeb2ctradestatusdriveAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r Alibabanlifeb2ctradestatusdriveAPIRequest) GetAction() string {
	return r._action
}

// SetExtendParams is ExtendParams Setter
// 扩展参数 JSON格式
func (r *Alibabanlifeb2ctradestatusdriveAPIRequest) SetExtendParams(_extendParams string) error {
	r._extendParams = _extendParams
	r.Set("extend_params", _extendParams)
	return nil
}

// GetExtendParams ExtendParams Getter
func (r Alibabanlifeb2ctradestatusdriveAPIRequest) GetExtendParams() string {
	return r._extendParams
}

// SetLogisticsInfo is LogisticsInfo Setter
// 货流信息
func (r *Alibabanlifeb2ctradestatusdriveAPIRequest) SetLogisticsInfo(_logisticsInfo *LogisticsInfo) error {
	r._logisticsInfo = _logisticsInfo
	r.Set("logistics_info", _logisticsInfo)
	return nil
}

// GetLogisticsInfo LogisticsInfo Getter
func (r Alibabanlifeb2ctradestatusdriveAPIRequest) GetLogisticsInfo() *LogisticsInfo {
	return r._logisticsInfo
}
