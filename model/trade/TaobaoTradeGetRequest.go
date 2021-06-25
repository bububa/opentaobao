package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔交易的部分信息(性能高) APIRequest
taobao.trade.get

获取单笔交易的部分信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong>
*/
type TaobaoTradeGetRequest struct {
    model.Params

    // 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
    fields   string 

    // 交易编号
    tid   int64 

}

func NewTaobaoTradeGetRequest() *TaobaoTradeGetRequest{
    return &TaobaoTradeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeGetRequest) GetApiMethodName() string {
    return "taobao.trade.get"
}

func (r TaobaoTradeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoTradeGetRequest) GetFields() string {
    return r.fields
}

func (r *TaobaoTradeGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoTradeGetRequest) GetTid() int64 {
    return r.tid
}

