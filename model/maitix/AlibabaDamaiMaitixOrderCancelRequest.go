package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦-库存释放 APIRequest
alibaba.damai.maitix.order.cancel

库存释放
*/
type AlibabaDamaiMaitixOrderCancelRequest struct {
    model.Params

    // 库存释放入参
    param   *MoaUnlockTicketParam 

}

func NewAlibabaDamaiMaitixOrderCancelRequest() *AlibabaDamaiMaitixOrderCancelRequest{
    return &AlibabaDamaiMaitixOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixOrderCancelRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.order.cancel"
}

func (r AlibabaDamaiMaitixOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixOrderCancelRequest) SetParam(param *MoaUnlockTicketParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaDamaiMaitixOrderCancelRequest) GetParam() *MoaUnlockTicketParam {
    return r.param
}

