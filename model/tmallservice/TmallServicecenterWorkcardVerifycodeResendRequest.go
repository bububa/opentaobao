package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
重发核销码 API请求
tmall.servicecenter.workcard.verifycode.resend

重发核销码
*/
type TmallServicecenterWorkcardVerifycodeResendRequest struct {
    model.Params
    // 工单id
    _workcardId   int64
    // 门店/网点id
    _serviceStoreId   int64
}

// 初始化TmallServicecenterWorkcardVerifycodeResendRequest对象
func NewTmallServicecenterWorkcardVerifycodeResendRequest() *TmallServicecenterWorkcardVerifycodeResendRequest{
    return &TmallServicecenterWorkcardVerifycodeResendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardVerifycodeResendRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.verifycode.resend"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardVerifycodeResendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardVerifycodeResendRequest) SetWorkcardId(_workcardId int64) error {
    r._workcardId = _workcardId
    r.Set("workcard_id", _workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardVerifycodeResendRequest) GetWorkcardId() int64 {
    return r._workcardId
}
// ServiceStoreId Setter
// 门店/网点id
func (r *TmallServicecenterWorkcardVerifycodeResendRequest) SetServiceStoreId(_serviceStoreId int64) error {
    r._serviceStoreId = _serviceStoreId
    r.Set("service_store_id", _serviceStoreId)
    return nil
}

// ServiceStoreId Getter
func (r TmallServicecenterWorkcardVerifycodeResendRequest) GetServiceStoreId() int64 {
    return r._serviceStoreId
}
