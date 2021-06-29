package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦-出票 APIRequest
alibaba.damai.maitix.order.confirm

出票
*/
type AlibabaDamaiMaitixOrderConfirmRequest struct {
    model.Params

    // 出票入参
    param   *MoaConfirmOrderParam 

}

func NewAlibabaDamaiMaitixOrderConfirmRequest() *AlibabaDamaiMaitixOrderConfirmRequest{
    return &AlibabaDamaiMaitixOrderConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixOrderConfirmRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.order.confirm"
}

func (r AlibabaDamaiMaitixOrderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixOrderConfirmRequest) SetParam(param *MoaConfirmOrderParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaDamaiMaitixOrderConfirmRequest) GetParam() *MoaConfirmOrderParam {
    return r.param
}

