package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对一笔交易添加备注 API请求
taobao.trade.memo.add

根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update
*/
type TaobaoTradeMemoAddRequest struct {
    model.Params
    // 交易编号
    _tid   int64
    // 交易备注。最大长度: 1000个字节
    _memo   string
    // 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
    _flag   int64
}

// 初始化TaobaoTradeMemoAddRequest对象
func NewTaobaoTradeMemoAddRequest() *TaobaoTradeMemoAddRequest{
    return &TaobaoTradeMemoAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeMemoAddRequest) GetApiMethodName() string {
    return "taobao.trade.memo.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeMemoAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 交易编号
func (r *TaobaoTradeMemoAddRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoTradeMemoAddRequest) GetTid() int64 {
    return r._tid
}
// Memo Setter
// 交易备注。最大长度: 1000个字节
func (r *TaobaoTradeMemoAddRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r TaobaoTradeMemoAddRequest) GetMemo() string {
    return r._memo
}
// Flag Setter
// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
func (r *TaobaoTradeMemoAddRequest) SetFlag(_flag int64) error {
    r._flag = _flag
    r.Set("flag", _flag)
    return nil
}

// Flag Getter
func (r TaobaoTradeMemoAddRequest) GetFlag() int64 {
    return r._flag
}
