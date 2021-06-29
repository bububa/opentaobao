package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦-查询分销单 APIRequest
alibaba.damai.maitix.order.query

查询分销单
*/
type AlibabaDamaiMaitixOrderQueryRequest struct {
    model.Params

    // 分销单查询入参
    param   *MoaOrderQueryParam 

}

func NewAlibabaDamaiMaitixOrderQueryRequest() *AlibabaDamaiMaitixOrderQueryRequest{
    return &AlibabaDamaiMaitixOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.order.query"
}

func (r AlibabaDamaiMaitixOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixOrderQueryRequest) SetParam(param *MoaOrderQueryParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaDamaiMaitixOrderQueryRequest) GetParam() *MoaOrderQueryParam {
    return r.param
}

