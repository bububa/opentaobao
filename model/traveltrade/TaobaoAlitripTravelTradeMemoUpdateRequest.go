package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改一笔交易备注 API请求
taobao.alitrip.travel.trade.memo.update

更新一笔交易备注
*/
type TaobaoAlitripTravelTradeMemoUpdateRequest struct {
    model.Params
    // 交易ID
    _tid   int64
    // 交易备注。最大长度: 1000个字节
    _memo   string
    // 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
    _flag   int64
    // 是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值
    _reset   bool
}

// 初始化TaobaoAlitripTravelTradeMemoUpdateRequest对象
func NewTaobaoAlitripTravelTradeMemoUpdateRequest() *TaobaoAlitripTravelTradeMemoUpdateRequest{
    return &TaobaoAlitripTravelTradeMemoUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelTradeMemoUpdateRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.trade.memo.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelTradeMemoUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 交易ID
func (r *TaobaoAlitripTravelTradeMemoUpdateRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoAlitripTravelTradeMemoUpdateRequest) GetTid() int64 {
    return r._tid
}
// Memo Setter
// 交易备注。最大长度: 1000个字节
func (r *TaobaoAlitripTravelTradeMemoUpdateRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r TaobaoAlitripTravelTradeMemoUpdateRequest) GetMemo() string {
    return r._memo
}
// Flag Setter
// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
func (r *TaobaoAlitripTravelTradeMemoUpdateRequest) SetFlag(_flag int64) error {
    r._flag = _flag
    r.Set("flag", _flag)
    return nil
}

// Flag Getter
func (r TaobaoAlitripTravelTradeMemoUpdateRequest) GetFlag() int64 {
    return r._flag
}
// Reset Setter
// 是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值
func (r *TaobaoAlitripTravelTradeMemoUpdateRequest) SetReset(_reset bool) error {
    r._reset = _reset
    r.Set("reset", _reset)
    return nil
}

// Reset Getter
func (r TaobaoAlitripTravelTradeMemoUpdateRequest) GetReset() bool {
    return r._reset
}
