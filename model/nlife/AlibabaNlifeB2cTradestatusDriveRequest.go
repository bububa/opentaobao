package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
b2c订单状态驱动 API请求
alibaba.nlife.b2c.tradestatus.drive

用于驱动零售+订单状态
*/
type AlibabaNlifeB2cTradestatusDriveRequest struct {
    model.Params
    // 零售门店在零售+平台的ID
    _storeId   string
    // APP:是指线上销售应用，POS:是指现场收银应用
    _channel   string
    // 对零售+为外部订单号，对业务方为订单号
    _outTradeNo   string
    // 零售+平台订单号，和out_trade_no不能同时为空
    _tradeNo   string
    // 接口类型：CONFIRM（收货）DELIVER（发货）
    _action   string
    // 货流信息
    _logisticsInfo   *LogisticsInfo
    // 扩展参数 JSON格式
    _extendParams   string
}

// 初始化AlibabaNlifeB2cTradestatusDriveRequest对象
func NewAlibabaNlifeB2cTradestatusDriveRequest() *AlibabaNlifeB2cTradestatusDriveRequest{
    return &AlibabaNlifeB2cTradestatusDriveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradestatusDriveRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.tradestatus.drive"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradestatusDriveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 零售门店在零售+平台的ID
func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeB2cTradestatusDriveRequest) GetStoreId() string {
    return r._storeId
}
// Channel Setter
// APP:是指线上销售应用，POS:是指现场收银应用
func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r AlibabaNlifeB2cTradestatusDriveRequest) GetChannel() string {
    return r._channel
}
// OutTradeNo Setter
// 对零售+为外部订单号，对业务方为订单号
func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetOutTradeNo(_outTradeNo string) error {
    r._outTradeNo = _outTradeNo
    r.Set("out_trade_no", _outTradeNo)
    return nil
}

// OutTradeNo Getter
func (r AlibabaNlifeB2cTradestatusDriveRequest) GetOutTradeNo() string {
    return r._outTradeNo
}
// TradeNo Setter
// 零售+平台订单号，和out_trade_no不能同时为空
func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetTradeNo(_tradeNo string) error {
    r._tradeNo = _tradeNo
    r.Set("trade_no", _tradeNo)
    return nil
}

// TradeNo Getter
func (r AlibabaNlifeB2cTradestatusDriveRequest) GetTradeNo() string {
    return r._tradeNo
}
// Action Setter
// 接口类型：CONFIRM（收货）DELIVER（发货）
func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r AlibabaNlifeB2cTradestatusDriveRequest) GetAction() string {
    return r._action
}
// LogisticsInfo Setter
// 货流信息
func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetLogisticsInfo(_logisticsInfo *LogisticsInfo) error {
    r._logisticsInfo = _logisticsInfo
    r.Set("logistics_info", _logisticsInfo)
    return nil
}

// LogisticsInfo Getter
func (r AlibabaNlifeB2cTradestatusDriveRequest) GetLogisticsInfo() *LogisticsInfo {
    return r._logisticsInfo
}
// ExtendParams Setter
// 扩展参数 JSON格式
func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetExtendParams(_extendParams string) error {
    r._extendParams = _extendParams
    r.Set("extend_params", _extendParams)
    return nil
}

// ExtendParams Getter
func (r AlibabaNlifeB2cTradestatusDriveRequest) GetExtendParams() string {
    return r._extendParams
}
