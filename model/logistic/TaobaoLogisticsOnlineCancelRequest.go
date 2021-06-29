package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消物流订单接口 API请求
taobao.logistics.online.cancel

调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
*/
type TaobaoLogisticsOnlineCancelRequest struct {
    model.Params
    // 淘宝交易ID
    _tid   int64
}

// 初始化TaobaoLogisticsOnlineCancelRequest对象
func NewTaobaoLogisticsOnlineCancelRequest() *TaobaoLogisticsOnlineCancelRequest{
    return &TaobaoLogisticsOnlineCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOnlineCancelRequest) GetApiMethodName() string {
    return "taobao.logistics.online.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOnlineCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝交易ID
func (r *TaobaoLogisticsOnlineCancelRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoLogisticsOnlineCancelRequest) GetTid() int64 {
    return r._tid
}
