package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消售后服务单 API请求
tmall.tmjlapp.sap.serviceorder.cancel

SAP跟天猫精灵app接口对接，用户在app取消sap售后服务工单
*/
type TmallTmjlappSapServiceorderCancelRequest struct {
    model.Params
    // 取消服务单请求
    _cancelRequest   *Dtcancelrequest
}

// 初始化TmallTmjlappSapServiceorderCancelRequest对象
func NewTmallTmjlappSapServiceorderCancelRequest() *TmallTmjlappSapServiceorderCancelRequest{
    return &TmallTmjlappSapServiceorderCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTmjlappSapServiceorderCancelRequest) GetApiMethodName() string {
    return "tmall.tmjlapp.sap.serviceorder.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TmallTmjlappSapServiceorderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CancelRequest Setter
// 取消服务单请求
func (r *TmallTmjlappSapServiceorderCancelRequest) SetCancelRequest(_cancelRequest *Dtcancelrequest) error {
    r._cancelRequest = _cancelRequest
    r.Set("cancel_request", _cancelRequest)
    return nil
}

// CancelRequest Getter
func (r TmallTmjlappSapServiceorderCancelRequest) GetCancelRequest() *Dtcancelrequest {
    return r._cancelRequest
}
