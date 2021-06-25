package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商城评价解释接口 APIRequest
taobao.traderate.explain.add

商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释）
*/
type TaobaoTraderateExplainAddRequest struct {
    model.Params

    // 子订单ID
    oid   int64 

    // 评价解释内容，最大长度：500个汉字
    reply   string 

}

func NewTaobaoTraderateExplainAddRequest() *TaobaoTraderateExplainAddRequest{
    return &TaobaoTraderateExplainAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTraderateExplainAddRequest) GetApiMethodName() string {
    return "taobao.traderate.explain.add"
}

func (r TaobaoTraderateExplainAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTraderateExplainAddRequest) SetOid(oid int64) error {
    r.oid = oid
    r.Set("oid", oid)
    return nil
}

func (r TaobaoTraderateExplainAddRequest) GetOid() int64 {
    return r.oid
}

func (r *TaobaoTraderateExplainAddRequest) SetReply(reply string) error {
    r.reply = reply
    r.Set("reply", reply)
    return nil
}

func (r TaobaoTraderateExplainAddRequest) GetReply() string {
    return r.reply
}

