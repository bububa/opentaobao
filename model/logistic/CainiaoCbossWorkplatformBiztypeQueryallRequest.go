package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型 API请求
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

// 初始化CainiaoCbossWorkplatformBiztypeQueryallRequest对象
func NewCainiaoCbossWorkplatformBiztypeQueryallRequest() *CainiaoCbossWorkplatformBiztypeQueryallRequest{
    return &CainiaoCbossWorkplatformBiztypeQueryallRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformBiztypeQueryallRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.biztype.queryall"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformBiztypeQueryallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Level Setter
// level
func (r *CainiaoCbossWorkplatformBiztypeQueryallRequest) SetLevel(level int64) error {
    r.level = level
    r.Set("level", level)
    return nil
}

// Level Getter
func (r CainiaoCbossWorkplatformBiztypeQueryallRequest) GetLevel() int64 {
    return r.level
}
// TradeId Setter
// tradeId
func (r *CainiaoCbossWorkplatformBiztypeQueryallRequest) SetTradeId(tradeId string) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

// TradeId Getter
func (r CainiaoCbossWorkplatformBiztypeQueryallRequest) GetTradeId() string {
    return r.tradeId
}
