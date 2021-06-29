package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增单个评价 API请求
taobao.traderate.add

新增单个评价(<font color="red">注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不能再通过该接口进行评价</font>)
*/
type TaobaoTraderateAddRequest struct {
    model.Params
    // 交易ID
    tid   int64
    // 子订单ID
    oid   int64
    // 评价结果,可选值:good(好评),neutral(中评),bad(差评)
    result   string
    // 评价者角色,可选值:seller(卖家),buyer(买家)
    role   string
    // 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
    content   string
    // 是否匿名,卖家评不能匿名。可选值:true(匿名),false(非匿名)。注意：如果交易订单匿名，则评价也匿名
    anony   bool
}

// 初始化TaobaoTraderateAddRequest对象
func NewTaobaoTraderateAddRequest() *TaobaoTraderateAddRequest{
    return &TaobaoTraderateAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTraderateAddRequest) GetApiMethodName() string {
    return "taobao.traderate.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTraderateAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 交易ID
func (r *TaobaoTraderateAddRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoTraderateAddRequest) GetTid() int64 {
    return r.tid
}
// Oid Setter
// 子订单ID
func (r *TaobaoTraderateAddRequest) SetOid(oid int64) error {
    r.oid = oid
    r.Set("oid", oid)
    return nil
}

// Oid Getter
func (r TaobaoTraderateAddRequest) GetOid() int64 {
    return r.oid
}
// Result Setter
// 评价结果,可选值:good(好评),neutral(中评),bad(差评)
func (r *TaobaoTraderateAddRequest) SetResult(result string) error {
    r.result = result
    r.Set("result", result)
    return nil
}

// Result Getter
func (r TaobaoTraderateAddRequest) GetResult() string {
    return r.result
}
// Role Setter
// 评价者角色,可选值:seller(卖家),buyer(买家)
func (r *TaobaoTraderateAddRequest) SetRole(role string) error {
    r.role = role
    r.Set("role", role)
    return nil
}

// Role Getter
func (r TaobaoTraderateAddRequest) GetRole() string {
    return r.role
}
// Content Setter
// 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
func (r *TaobaoTraderateAddRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r TaobaoTraderateAddRequest) GetContent() string {
    return r.content
}
// Anony Setter
// 是否匿名,卖家评不能匿名。可选值:true(匿名),false(非匿名)。注意：如果交易订单匿名，则评价也匿名
func (r *TaobaoTraderateAddRequest) SetAnony(anony bool) error {
    r.anony = anony
    r.Set("anony", anony)
    return nil
}

// Anony Getter
func (r TaobaoTraderateAddRequest) GetAnony() bool {
    return r.anony
}
