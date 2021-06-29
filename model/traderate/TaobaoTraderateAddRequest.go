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
    _tid   int64
    // 子订单ID
    _oid   int64
    // 评价结果,可选值:good(好评),neutral(中评),bad(差评)
    _result   string
    // 评价者角色,可选值:seller(卖家),buyer(买家)
    _role   string
    // 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
    _content   string
    // 是否匿名,卖家评不能匿名。可选值:true(匿名),false(非匿名)。注意：如果交易订单匿名，则评价也匿名
    _anony   bool
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
func (r *TaobaoTraderateAddRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoTraderateAddRequest) GetTid() int64 {
    return r._tid
}
// Oid Setter
// 子订单ID
func (r *TaobaoTraderateAddRequest) SetOid(_oid int64) error {
    r._oid = _oid
    r.Set("oid", _oid)
    return nil
}

// Oid Getter
func (r TaobaoTraderateAddRequest) GetOid() int64 {
    return r._oid
}
// Result Setter
// 评价结果,可选值:good(好评),neutral(中评),bad(差评)
func (r *TaobaoTraderateAddRequest) SetResult(_result string) error {
    r._result = _result
    r.Set("result", _result)
    return nil
}

// Result Getter
func (r TaobaoTraderateAddRequest) GetResult() string {
    return r._result
}
// Role Setter
// 评价者角色,可选值:seller(卖家),buyer(买家)
func (r *TaobaoTraderateAddRequest) SetRole(_role string) error {
    r._role = _role
    r.Set("role", _role)
    return nil
}

// Role Getter
func (r TaobaoTraderateAddRequest) GetRole() string {
    return r._role
}
// Content Setter
// 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
func (r *TaobaoTraderateAddRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoTraderateAddRequest) GetContent() string {
    return r._content
}
// Anony Setter
// 是否匿名,卖家评不能匿名。可选值:true(匿名),false(非匿名)。注意：如果交易订单匿名，则评价也匿名
func (r *TaobaoTraderateAddRequest) SetAnony(_anony bool) error {
    r._anony = _anony
    r.Set("anony", _anony)
    return nil
}

// Anony Getter
func (r TaobaoTraderateAddRequest) GetAnony() bool {
    return r._anony
}
