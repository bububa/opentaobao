package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消售后服务单 APIRequest
tmall.tmjlapp.sap.serviceorder.cancel

SAP跟天猫精灵app接口对接，用户在app取消sap售后服务工单
*/
type TmallTmjlappSapServiceorderCancelRequest struct {
    model.Params

    // 取消服务单请求
    cancelRequest   *Dtcancelrequest 

}

func NewTmallTmjlappSapServiceorderCancelRequest() *TmallTmjlappSapServiceorderCancelRequest{
    return &TmallTmjlappSapServiceorderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTmjlappSapServiceorderCancelRequest) GetApiMethodName() string {
    return "tmall.tmjlapp.sap.serviceorder.cancel"
}

func (r TmallTmjlappSapServiceorderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTmjlappSapServiceorderCancelRequest) SetCancelRequest(cancelRequest *Dtcancelrequest) error {
    r.cancelRequest = cancelRequest
    r.Set("cancel_request", cancelRequest)
    return nil
}

func (r TmallTmjlappSapServiceorderCancelRequest) GetCancelRequest() *Dtcancelrequest {
    return r.cancelRequest
}

