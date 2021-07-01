package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加一笔交易备注 API请求
taobao.alitrip.travel.trade.memo.add

对一笔交易添加备注
*/
type TaobaoAlitripTravelTradeMemoAddAPIRequest struct {
    model.Params
    // 交易ID
    _tid   int64
    // 交易备注。最大长度: 1000个字节
    _memo   string
    // 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
    _flag   int64
}

// 初始化TaobaoAlitripTravelTradeMemoAddAPIRequest对象
func NewTaobaoAlitripTravelTradeMemoAddRequest() *TaobaoAlitripTravelTradeMemoAddAPIRequest{
    return &TaobaoAlitripTravelTradeMemoAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelTradeMemoAddAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.trade.memo.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelTradeMemoAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 交易ID
func (r *TaobaoAlitripTravelTradeMemoAddAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoAlitripTravelTradeMemoAddAPIRequest) GetTid() int64 {
    return r._tid
}
// Memo Setter
// 交易备注。最大长度: 1000个字节
func (r *TaobaoAlitripTravelTradeMemoAddAPIRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r TaobaoAlitripTravelTradeMemoAddAPIRequest) GetMemo() string {
    return r._memo
}
// Flag Setter
// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
func (r *TaobaoAlitripTravelTradeMemoAddAPIRequest) SetFlag(_flag int64) error {
    r._flag = _flag
    r.Set("flag", _flag)
    return nil
}

// Flag Getter
func (r TaobaoAlitripTravelTradeMemoAddAPIRequest) GetFlag() int64 {
    return r._flag
}
