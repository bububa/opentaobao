package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔交易的部分信息(商家应用使用) API请求
taobao.open.trade.get

获取单笔交易的部分信息</br>
1.入参fields中传入buyer_nick ，才能返回buyer_open_id
*/
type TaobaoOpenTradeGetRequest struct {
    model.Params
    // 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
    _fields   string
    // 交易编号
    _tid   int64
}

// 初始化TaobaoOpenTradeGetRequest对象
func NewTaobaoOpenTradeGetRequest() *TaobaoOpenTradeGetRequest{
    return &TaobaoOpenTradeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenTradeGetRequest) GetApiMethodName() string {
    return "taobao.open.trade.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenTradeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
func (r *TaobaoOpenTradeGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoOpenTradeGetRequest) GetFields() string {
    return r._fields
}
// Tid Setter
// 交易编号
func (r *TaobaoOpenTradeGetRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenTradeGetRequest) GetTid() int64 {
    return r._tid
}
