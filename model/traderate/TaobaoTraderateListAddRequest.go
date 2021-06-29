package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
针对父子订单新增批量评价 API请求
taobao.traderate.list.add

针对父子订单新增批量评价(<font color="red">注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不用再通过该接口进行评价</font>)
*/
type TaobaoTraderateListAddRequest struct {
    model.Params
    // 交易ID
    _tid   int64
    // 评价结果。可选值:good(好评),neutral(中评),bad(差评)
    _result   string
    // 评价者角色。可选值:seller(卖家),buyer(买家)
    _role   string
    // 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
    _content   string
    // 是否匿名，卖家评不能匿名。可选值:true(匿名),false(非匿名)。 注意：如果买家匿名购买，那么买家的评价默认匿名
    _anony   bool
}

// 初始化TaobaoTraderateListAddRequest对象
func NewTaobaoTraderateListAddRequest() *TaobaoTraderateListAddRequest{
    return &TaobaoTraderateListAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTraderateListAddRequest) GetApiMethodName() string {
    return "taobao.traderate.list.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTraderateListAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 交易ID
func (r *TaobaoTraderateListAddRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoTraderateListAddRequest) GetTid() int64 {
    return r._tid
}
// Result Setter
// 评价结果。可选值:good(好评),neutral(中评),bad(差评)
func (r *TaobaoTraderateListAddRequest) SetResult(_result string) error {
    r._result = _result
    r.Set("result", _result)
    return nil
}

// Result Getter
func (r TaobaoTraderateListAddRequest) GetResult() string {
    return r._result
}
// Role Setter
// 评价者角色。可选值:seller(卖家),buyer(买家)
func (r *TaobaoTraderateListAddRequest) SetRole(_role string) error {
    r._role = _role
    r.Set("role", _role)
    return nil
}

// Role Getter
func (r TaobaoTraderateListAddRequest) GetRole() string {
    return r._role
}
// Content Setter
// 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
func (r *TaobaoTraderateListAddRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoTraderateListAddRequest) GetContent() string {
    return r._content
}
// Anony Setter
// 是否匿名，卖家评不能匿名。可选值:true(匿名),false(非匿名)。 注意：如果买家匿名购买，那么买家的评价默认匿名
func (r *TaobaoTraderateListAddRequest) SetAnony(_anony bool) error {
    r._anony = _anony
    r.Set("anony", _anony)
    return nil
}

// Anony Getter
func (r TaobaoTraderateListAddRequest) GetAnony() bool {
    return r._anony
}
