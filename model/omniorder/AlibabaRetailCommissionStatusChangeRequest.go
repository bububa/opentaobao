package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分佣状态变更 APIRequest
alibaba.retail.commission.status.change

分佣系统，分佣状态变更接口
*/
type AlibabaRetailCommissionStatusChangeRequest struct {
    model.Params

    // 请求参数
    param0   *UniverseOrderVo 

}

func NewAlibabaRetailCommissionStatusChangeRequest() *AlibabaRetailCommissionStatusChangeRequest{
    return &AlibabaRetailCommissionStatusChangeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailCommissionStatusChangeRequest) GetApiMethodName() string {
    return "alibaba.retail.commission.status.change"
}

func (r AlibabaRetailCommissionStatusChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailCommissionStatusChangeRequest) SetParam0(param0 *UniverseOrderVo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaRetailCommissionStatusChangeRequest) GetParam0() *UniverseOrderVo {
    return r.param0
}

