package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据交易单号判断是否为菜鸟发货订单 API请求
cainiao.cboss.workplatform.logistics.iscainiaoorder

根据交易单号判断是否为菜鸟发货订单
*/
type CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest struct {
    model.Params
    // 交易单号
    _tradeId   string
}

// 初始化CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest对象
func NewCainiaoCbossWorkplatformLogisticsIscainiaoorderRequest() *CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest{
    return &CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.logistics.iscainiaoorder"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeId Setter
// 交易单号
func (r *CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest) SetTradeId(_tradeId string) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest) GetTradeId() string {
    return r._tradeId
}
