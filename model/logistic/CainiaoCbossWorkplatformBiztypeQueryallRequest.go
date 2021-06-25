package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型 APIRequest
cainiao.cboss.workplatform.biztype.queryall

菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型。 目前调用者ISV
*/
type CainiaoCbossWorkplatformBiztypeQueryallRequest struct {
    model.Params

    // level
    level   int64 

    // tradeId
    tradeId   string 

}

func NewCainiaoCbossWorkplatformBiztypeQueryallRequest() *CainiaoCbossWorkplatformBiztypeQueryallRequest{
    return &CainiaoCbossWorkplatformBiztypeQueryallRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCbossWorkplatformBiztypeQueryallRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.biztype.queryall"
}

func (r CainiaoCbossWorkplatformBiztypeQueryallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCbossWorkplatformBiztypeQueryallRequest) SetLevel(level int64) error {
    r.level = level
    r.Set("level", level)
    return nil
}

func (r CainiaoCbossWorkplatformBiztypeQueryallRequest) GetLevel() int64 {
    return r.level
}

func (r *CainiaoCbossWorkplatformBiztypeQueryallRequest) SetTradeId(tradeId string) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r CainiaoCbossWorkplatformBiztypeQueryallRequest) GetTradeId() string {
    return r.tradeId
}

