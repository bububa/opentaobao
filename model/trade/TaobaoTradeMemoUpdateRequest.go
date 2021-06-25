package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改交易备注 APIRequest
taobao.trade.memo.update

需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能
*/
type TaobaoTradeMemoUpdateRequest struct {
    model.Params

    // 交易编号
    tid   int64 

    // 卖家交易备注。最大长度: 1000个字节
    memo   string 

    // 卖家交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
    flag   int64 

    // 是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值
    reset   bool 

}

func NewTaobaoTradeMemoUpdateRequest() *TaobaoTradeMemoUpdateRequest{
    return &TaobaoTradeMemoUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeMemoUpdateRequest) GetApiMethodName() string {
    return "taobao.trade.memo.update"
}

func (r TaobaoTradeMemoUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeMemoUpdateRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoTradeMemoUpdateRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoTradeMemoUpdateRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r TaobaoTradeMemoUpdateRequest) GetMemo() string {
    return r.memo
}

func (r *TaobaoTradeMemoUpdateRequest) SetFlag(flag int64) error {
    r.flag = flag
    r.Set("flag", flag)
    return nil
}

func (r TaobaoTradeMemoUpdateRequest) GetFlag() int64 {
    return r.flag
}

func (r *TaobaoTradeMemoUpdateRequest) SetReset(reset bool) error {
    r.reset = reset
    r.Set("reset", reset)
    return nil
}

func (r TaobaoTradeMemoUpdateRequest) GetReset() bool {
    return r.reset
}

