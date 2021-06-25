package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
针对父子订单新增批量评价 APIRequest
taobao.traderate.list.add

针对父子订单新增批量评价(<font color="red">注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不用再通过该接口进行评价</font>)
*/
type TaobaoTraderateListAddRequest struct {
    model.Params

    // 交易ID
    tid   int64 

    // 评价结果。可选值:good(好评),neutral(中评),bad(差评)
    result   string 

    // 评价者角色。可选值:seller(卖家),buyer(买家)
    role   string 

    // 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
    content   string 

    // 是否匿名，卖家评不能匿名。可选值:true(匿名),false(非匿名)。 注意：如果买家匿名购买，那么买家的评价默认匿名
    anony   bool 

}

func NewTaobaoTraderateListAddRequest() *TaobaoTraderateListAddRequest{
    return &TaobaoTraderateListAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTraderateListAddRequest) GetApiMethodName() string {
    return "taobao.traderate.list.add"
}

func (r TaobaoTraderateListAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTraderateListAddRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoTraderateListAddRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoTraderateListAddRequest) SetResult(result string) error {
    r.result = result
    r.Set("result", result)
    return nil
}

func (r TaobaoTraderateListAddRequest) GetResult() string {
    return r.result
}

func (r *TaobaoTraderateListAddRequest) SetRole(role string) error {
    r.role = role
    r.Set("role", role)
    return nil
}

func (r TaobaoTraderateListAddRequest) GetRole() string {
    return r.role
}

func (r *TaobaoTraderateListAddRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoTraderateListAddRequest) GetContent() string {
    return r.content
}

func (r *TaobaoTraderateListAddRequest) SetAnony(anony bool) error {
    r.anony = anony
    r.Set("anony", anony)
    return nil
}

func (r TaobaoTraderateListAddRequest) GetAnony() bool {
    return r.anony
}

