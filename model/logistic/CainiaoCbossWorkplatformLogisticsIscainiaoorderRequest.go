package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据交易单号判断是否为菜鸟发货订单 APIRequest
cainiao.cboss.workplatform.logistics.iscainiaoorder

根据交易单号判断是否为菜鸟发货订单
*/
type CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest struct {
    model.Params

    // 交易单号
    tradeId   string 

}

func NewCainiaoCbossWorkplatformLogisticsIscainiaoorderRequest() *CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest{
    return &CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.logistics.iscainiaoorder"
}

func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest) SetTradeId(tradeId string) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderRequest) GetTradeId() string {
    return r.tradeId
}

