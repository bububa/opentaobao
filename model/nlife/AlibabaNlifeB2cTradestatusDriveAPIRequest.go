package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cTradestatusDriveAPIRequest b2c订单状态驱动 API请求
// alibaba.nlife.b2c.tradestatus.drive
//
// 用于驱动零售+订单状态
type AlibabaNlifeB2cTradestatusDriveAPIRequest struct {
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
	// 货流信息
	_logisticsInfo *LogisticsInfo
	// 扩展参数 JSON格式
	_extendParams string
}

// NewAlibabaNlifeB2cTradestatusDriveRequest 初始化AlibabaNlifeB2cTradestatusDriveAPIRequest对象
func NewAlibabaNlifeB2cTradestatusDriveRequest() *AlibabaNlifeB2cTradestatusDriveAPIRequest {
	return &AlibabaNlifeB2cTradestatusDriveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.tradestatus.drive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 零售门店在零售+平台的ID
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is Channel Setter
// APP:是指线上销售应用，POS:是指现场收银应用
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// Get Channel Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetChannel() string {
	return r._channel
}

// Set is OutTradeNo Setter
// 对零售+为外部订单号，对业务方为订单号
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// Get OutTradeNo Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// Set is TradeNo Setter
// 零售+平台订单号，和out_trade_no不能同时为空
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// Get TradeNo Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// Set is Action Setter
// 接口类型：CONFIRM（收货）DELIVER（发货）
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// Get Action Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetAction() string {
	return r._action
}

// Set is LogisticsInfo Setter
// 货流信息
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetLogisticsInfo(_logisticsInfo *LogisticsInfo) error {
	r._logisticsInfo = _logisticsInfo
	r.Set("logistics_info", _logisticsInfo)
	return nil
}

// Get LogisticsInfo Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetLogisticsInfo() *LogisticsInfo {
	return r._logisticsInfo
}

// Set is ExtendParams Setter
// 扩展参数 JSON格式
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetExtendParams(_extendParams string) error {
	r._extendParams = _extendParams
	r.Set("extend_params", _extendParams)
	return nil
}

// Get ExtendParams Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetExtendParams() string {
	return r._extendParams
}
