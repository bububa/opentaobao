package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分佣结果查询 API请求
alibaba.retail.commission.result.query

查询导购分佣记录
*/
type AlibabaRetailCommissionResultQueryRequest struct {
    model.Params
    // 请求参数
    _param0   *CommissionResultQuery
}

// 初始化AlibabaRetailCommissionResultQueryRequest对象
func NewAlibabaRetailCommissionResultQueryRequest() *AlibabaRetailCommissionResultQueryRequest{
    return &AlibabaRetailCommissionResultQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailCommissionResultQueryRequest) GetApiMethodName() string {
    return "alibaba.retail.commission.result.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailCommissionResultQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 请求参数
func (r *AlibabaRetailCommissionResultQueryRequest) SetParam0(_param0 *CommissionResultQuery) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaRetailCommissionResultQueryRequest) GetParam0() *CommissionResultQuery {
    return r._param0
}
