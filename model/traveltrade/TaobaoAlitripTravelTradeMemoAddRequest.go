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
type TaobaoAlitripTravelTradeMemoAddRequest struct {
    model.Params
    // 交易ID
    tid   int64
    // 交易备注。最大长度: 1000个字节
    memo   string
    // 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
    flag   int64
}

// 初始化TaobaoAlitripTravelTradeMemoAddRequest对象
func NewTaobaoAlitripTravelTradeMemoAddRequest() *TaobaoAlitripTravelTradeMemoAddRequest{
    return &TaobaoAlitripTravelTradeMemoAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelTradeMemoAddRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.trade.memo.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelTradeMemoAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 交易ID
func (r *TaobaoAlitripTravelTradeMemoAddRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoAlitripTravelTradeMemoAddRequest) GetTid() int64 {
    return r.tid
}
// Memo Setter
// 交易备注。最大长度: 1000个字节
func (r *TaobaoAlitripTravelTradeMemoAddRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

// Memo Getter
func (r TaobaoAlitripTravelTradeMemoAddRequest) GetMemo() string {
    return r.memo
}
// Flag Setter
// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
func (r *TaobaoAlitripTravelTradeMemoAddRequest) SetFlag(flag int64) error {
    r.flag = flag
    r.Set("flag", flag)
    return nil
}

// Flag Getter
func (r TaobaoAlitripTravelTradeMemoAddRequest) GetFlag() int64 {
    return r.flag
}
