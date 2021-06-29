package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
b2c订单状态驱动 APIRequest
alibaba.nlife.b2c.tradestatus.drive

用于驱动零售+订单状态
*/
type AlibabaNlifeB2cTradestatusDriveRequest struct {
    model.Params

    // 零售门店在零售+平台的ID
    storeId   string 

    // APP:是指线上销售应用，POS:是指现场收银应用
    channel   string 

    // 对零售+为外部订单号，对业务方为订单号
    outTradeNo   string 

    // 零售+平台订单号，和out_trade_no不能同时为空
    tradeNo   string 

    // 接口类型：CONFIRM（收货）DELIVER（发货）
    action   string 

    // 货流信息
    logisticsInfo   *LogisticsInfo 

    // 扩展参数 JSON格式
    extendParams   string 

}

func NewAlibabaNlifeB2cTradestatusDriveRequest() *AlibabaNlifeB2cTradestatusDriveRequest{
    return &AlibabaNlifeB2cTradestatusDriveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNlifeB2cTradestatusDriveRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.tradestatus.drive"
}

func (r AlibabaNlifeB2cTradestatusDriveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaNlifeB2cTradestatusDriveRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r AlibabaNlifeB2cTradestatusDriveRequest) GetChannel() string {
    return r.channel
}

func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetOutTradeNo(outTradeNo string) error {
    r.outTradeNo = outTradeNo
    r.Set("out_trade_no", outTradeNo)
    return nil
}

func (r AlibabaNlifeB2cTradestatusDriveRequest) GetOutTradeNo() string {
    return r.outTradeNo
}

func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

func (r AlibabaNlifeB2cTradestatusDriveRequest) GetTradeNo() string {
    return r.tradeNo
}

func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

func (r AlibabaNlifeB2cTradestatusDriveRequest) GetAction() string {
    return r.action
}

func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetLogisticsInfo(logisticsInfo *LogisticsInfo) error {
    r.logisticsInfo = logisticsInfo
    r.Set("logistics_info", logisticsInfo)
    return nil
}

func (r AlibabaNlifeB2cTradestatusDriveRequest) GetLogisticsInfo() *LogisticsInfo {
    return r.logisticsInfo
}

func (r *AlibabaNlifeB2cTradestatusDriveRequest) SetExtendParams(extendParams string) error {
    r.extendParams = extendParams
    r.Set("extend_params", extendParams)
    return nil
}

func (r AlibabaNlifeB2cTradestatusDriveRequest) GetExtendParams() string {
    return r.extendParams
}

