package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分佣结果查询 APIRequest
alibaba.retail.commission.result.query

查询导购分佣记录
*/
type AlibabaRetailCommissionResultQueryRequest struct {
    model.Params

    // 请求参数
    param0   *CommissionResultQuery 

}

func NewAlibabaRetailCommissionResultQueryRequest() *AlibabaRetailCommissionResultQueryRequest{
    return &AlibabaRetailCommissionResultQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailCommissionResultQueryRequest) GetApiMethodName() string {
    return "alibaba.retail.commission.result.query"
}

func (r AlibabaRetailCommissionResultQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailCommissionResultQueryRequest) SetParam0(param0 *CommissionResultQuery) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaRetailCommissionResultQueryRequest) GetParam0() *CommissionResultQuery {
    return r.param0
}

