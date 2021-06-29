package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分佣数据传输 APIRequest
alibaba.retail.commission.order.sync

同步分佣结果
*/
type AlibabaRetailCommissionOrderSyncRequest struct {
    model.Params

    // 请求参数
    param0   *UniverseOrderVo 

}

func NewAlibabaRetailCommissionOrderSyncRequest() *AlibabaRetailCommissionOrderSyncRequest{
    return &AlibabaRetailCommissionOrderSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailCommissionOrderSyncRequest) GetApiMethodName() string {
    return "alibaba.retail.commission.order.sync"
}

func (r AlibabaRetailCommissionOrderSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailCommissionOrderSyncRequest) SetParam0(param0 *UniverseOrderVo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaRetailCommissionOrderSyncRequest) GetParam0() *UniverseOrderVo {
    return r.param0
}

