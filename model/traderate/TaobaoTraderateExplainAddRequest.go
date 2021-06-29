package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商城评价解释接口 API请求
taobao.traderate.explain.add

商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释）
*/
type TaobaoTraderateExplainAddRequest struct {
    model.Params
    // 子订单ID
    _oid   int64
    // 评价解释内容，最大长度：500个汉字
    _reply   string
}

// 初始化TaobaoTraderateExplainAddRequest对象
func NewTaobaoTraderateExplainAddRequest() *TaobaoTraderateExplainAddRequest{
    return &TaobaoTraderateExplainAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTraderateExplainAddRequest) GetApiMethodName() string {
    return "taobao.traderate.explain.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTraderateExplainAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Oid Setter
// 子订单ID
func (r *TaobaoTraderateExplainAddRequest) SetOid(_oid int64) error {
    r._oid = _oid
    r.Set("oid", _oid)
    return nil
}

// Oid Getter
func (r TaobaoTraderateExplainAddRequest) GetOid() int64 {
    return r._oid
}
// Reply Setter
// 评价解释内容，最大长度：500个汉字
func (r *TaobaoTraderateExplainAddRequest) SetReply(_reply string) error {
    r._reply = _reply
    r.Set("reply", _reply)
    return nil
}

// Reply Getter
func (r TaobaoTraderateExplainAddRequest) GetReply() string {
    return r._reply
}
