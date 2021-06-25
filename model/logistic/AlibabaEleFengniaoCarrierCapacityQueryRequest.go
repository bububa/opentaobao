package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按照门店查询骑手运力状态查询 APIRequest
alibaba.ele.fengniao.carrier.capacity.query

提供给大润发，用于按照站点纬度查询大润发每个配送站的实时上班骑手数、到店骑手数、活跃骑手数量
*/
type AlibabaEleFengniaoCarrierCapacityQueryRequest struct {
    model.Params

    // 系统自动生成
    param   *Param 

}

func NewAlibabaEleFengniaoCarrierCapacityQueryRequest() *AlibabaEleFengniaoCarrierCapacityQueryRequest{
    return &AlibabaEleFengniaoCarrierCapacityQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoCarrierCapacityQueryRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.carrier.capacity.query"
}

func (r AlibabaEleFengniaoCarrierCapacityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoCarrierCapacityQueryRequest) SetParam(param *Param) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaEleFengniaoCarrierCapacityQueryRequest) GetParam() *Param {
    return r.param
}

