package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家取消获取的电子面单号v1.0 API请求
taobao.wlb.waybill.i.cancel

面单号有误需要取消的时候，调用该接口取消获取的电子面单。
*/
type TaobaoWlbWaybillICancelRequest struct {
    model.Params
    // 取消接口入参
    _waybillApplyCancelRequest   *WaybillApplyCancelRequest
}

// 初始化TaobaoWlbWaybillICancelRequest对象
func NewTaobaoWlbWaybillICancelRequest() *TaobaoWlbWaybillICancelRequest{
    return &TaobaoWlbWaybillICancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillICancelRequest) GetApiMethodName() string {
    return "taobao.wlb.waybill.i.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillICancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WaybillApplyCancelRequest Setter
// 取消接口入参
func (r *TaobaoWlbWaybillICancelRequest) SetWaybillApplyCancelRequest(_waybillApplyCancelRequest *WaybillApplyCancelRequest) error {
    r._waybillApplyCancelRequest = _waybillApplyCancelRequest
    r.Set("waybill_apply_cancel_request", _waybillApplyCancelRequest)
    return nil
}

// WaybillApplyCancelRequest Getter
func (r TaobaoWlbWaybillICancelRequest) GetWaybillApplyCancelRequest() *WaybillApplyCancelRequest {
    return r._waybillApplyCancelRequest
}
