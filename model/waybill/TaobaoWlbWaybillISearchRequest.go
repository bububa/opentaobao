package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询面单服务订购及面单使用情况v1.0 API请求
taobao.wlb.waybill.i.search

获取发货地&CP开通状态&账户的使用情况
*/
type TaobaoWlbWaybillISearchRequest struct {
    model.Params
    // 查询网点信息
    _waybillApplyRequest   *WaybillApplyRequest
}

// 初始化TaobaoWlbWaybillISearchRequest对象
func NewTaobaoWlbWaybillISearchRequest() *TaobaoWlbWaybillISearchRequest{
    return &TaobaoWlbWaybillISearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillISearchRequest) GetApiMethodName() string {
    return "taobao.wlb.waybill.i.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillISearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WaybillApplyRequest Setter
// 查询网点信息
func (r *TaobaoWlbWaybillISearchRequest) SetWaybillApplyRequest(_waybillApplyRequest *WaybillApplyRequest) error {
    r._waybillApplyRequest = _waybillApplyRequest
    r.Set("waybill_apply_request", _waybillApplyRequest)
    return nil
}

// WaybillApplyRequest Getter
func (r TaobaoWlbWaybillISearchRequest) GetWaybillApplyRequest() *WaybillApplyRequest {
    return r._waybillApplyRequest
}
