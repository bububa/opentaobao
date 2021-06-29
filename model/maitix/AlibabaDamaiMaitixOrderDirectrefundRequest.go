package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦-直接退票 APIRequest
alibaba.damai.maitix.order.directrefund

大麦-退票
*/
type AlibabaDamaiMaitixOrderDirectrefundRequest struct {
    model.Params

    // 退票入参
    param   *MoaRefundAuditParam 

}

func NewAlibabaDamaiMaitixOrderDirectrefundRequest() *AlibabaDamaiMaitixOrderDirectrefundRequest{
    return &AlibabaDamaiMaitixOrderDirectrefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixOrderDirectrefundRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.order.directrefund"
}

func (r AlibabaDamaiMaitixOrderDirectrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixOrderDirectrefundRequest) SetParam(param *MoaRefundAuditParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaDamaiMaitixOrderDirectrefundRequest) GetParam() *MoaRefundAuditParam {
    return r.param
}

