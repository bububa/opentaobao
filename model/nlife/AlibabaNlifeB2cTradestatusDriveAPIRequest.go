package nlife

import (
	"net/url"
	"sync"

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
	// 扩展参数 JSON格式
	_extendParams string
	// 货流信息
	_logisticsInfo *LogisticsInfo
}

// NewAlibabaNlifeB2cTradestatusDriveRequest 初始化AlibabaNlifeB2cTradestatusDriveAPIRequest对象
func NewAlibabaNlifeB2cTradestatusDriveRequest() *AlibabaNlifeB2cTradestatusDriveAPIRequest {
	return &AlibabaNlifeB2cTradestatusDriveAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) Reset() {
	r._storeId = ""
	r._channel = ""
	r._outTradeNo = ""
	r._tradeNo = ""
	r._action = ""
	r._extendParams = ""
	r._logisticsInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.tradestatus.drive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 零售门店在零售+平台的ID
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetChannel is Channel Setter
// APP:是指线上销售应用，POS:是指现场收银应用
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetChannel() string {
	return r._channel
}

// SetOutTradeNo is OutTradeNo Setter
// 对零售+为外部订单号，对业务方为订单号
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// SetTradeNo is TradeNo Setter
// 零售+平台订单号，和out_trade_no不能同时为空
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetAction is Action Setter
// 接口类型：CONFIRM（收货）DELIVER（发货）
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetAction() string {
	return r._action
}

// SetExtendParams is ExtendParams Setter
// 扩展参数 JSON格式
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetExtendParams(_extendParams string) error {
	r._extendParams = _extendParams
	r.Set("extend_params", _extendParams)
	return nil
}

// GetExtendParams ExtendParams Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetExtendParams() string {
	return r._extendParams
}

// SetLogisticsInfo is LogisticsInfo Setter
// 货流信息
func (r *AlibabaNlifeB2cTradestatusDriveAPIRequest) SetLogisticsInfo(_logisticsInfo *LogisticsInfo) error {
	r._logisticsInfo = _logisticsInfo
	r.Set("logistics_info", _logisticsInfo)
	return nil
}

// GetLogisticsInfo LogisticsInfo Getter
func (r AlibabaNlifeB2cTradestatusDriveAPIRequest) GetLogisticsInfo() *LogisticsInfo {
	return r._logisticsInfo
}

var poolAlibabaNlifeB2cTradestatusDriveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaNlifeB2cTradestatusDriveRequest()
	},
}

// GetAlibabaNlifeB2cTradestatusDriveRequest 从 sync.Pool 获取 AlibabaNlifeB2cTradestatusDriveAPIRequest
func GetAlibabaNlifeB2cTradestatusDriveAPIRequest() *AlibabaNlifeB2cTradestatusDriveAPIRequest {
	return poolAlibabaNlifeB2cTradestatusDriveAPIRequest.Get().(*AlibabaNlifeB2cTradestatusDriveAPIRequest)
}

// ReleaseAlibabaNlifeB2cTradestatusDriveAPIRequest 将 AlibabaNlifeB2cTradestatusDriveAPIRequest 放入 sync.Pool
func ReleaseAlibabaNlifeB2cTradestatusDriveAPIRequest(v *AlibabaNlifeB2cTradestatusDriveAPIRequest) {
	v.Reset()
	poolAlibabaNlifeB2cTradestatusDriveAPIRequest.Put(v)
}
