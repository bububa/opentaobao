package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对一笔交易添加备注 APIRequest
taobao.trade.memo.add

根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update
*/
type TaobaoTradeMemoAddRequest struct {
    model.Params

    // 交易编号
    tid   int64 

    // 交易备注。最大长度: 1000个字节
    memo   string 

    // 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
    flag   int64 

}

func NewTaobaoTradeMemoAddRequest() *TaobaoTradeMemoAddRequest{
    return &TaobaoTradeMemoAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeMemoAddRequest) GetApiMethodName() string {
    return "taobao.trade.memo.add"
}

func (r TaobaoTradeMemoAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeMemoAddRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoTradeMemoAddRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoTradeMemoAddRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r TaobaoTradeMemoAddRequest) GetMemo() string {
    return r.memo
}

func (r *TaobaoTradeMemoAddRequest) SetFlag(flag int64) error {
    r.flag = flag
    r.Set("flag", flag)
    return nil
}

func (r TaobaoTradeMemoAddRequest) GetFlag() int64 {
    return r.flag
}

