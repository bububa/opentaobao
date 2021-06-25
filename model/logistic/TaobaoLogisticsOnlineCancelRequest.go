package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消物流订单接口 APIRequest
taobao.logistics.online.cancel

调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
*/
type TaobaoLogisticsOnlineCancelRequest struct {
    model.Params

    // 淘宝交易ID
    tid   int64 

}

func NewTaobaoLogisticsOnlineCancelRequest() *TaobaoLogisticsOnlineCancelRequest{
    return &TaobaoLogisticsOnlineCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsOnlineCancelRequest) GetApiMethodName() string {
    return "taobao.logistics.online.cancel"
}

func (r TaobaoLogisticsOnlineCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsOnlineCancelRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoLogisticsOnlineCancelRequest) GetTid() int64 {
    return r.tid
}

