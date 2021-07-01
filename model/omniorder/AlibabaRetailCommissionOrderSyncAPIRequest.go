package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分佣数据传输 API请求
alibaba.retail.commission.order.sync

同步分佣结果
*/
type AlibabaRetailCommissionOrderSyncAPIRequest struct {
    model.Params
    // 请求参数
    _param0   *UniverseOrderVo
}

// 初始化AlibabaRetailCommissionOrderSyncAPIRequest对象
func NewAlibabaRetailCommissionOrderSyncRequest() *AlibabaRetailCommissionOrderSyncAPIRequest{
    return &AlibabaRetailCommissionOrderSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailCommissionOrderSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.retail.commission.order.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailCommissionOrderSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 请求参数
func (r *AlibabaRetailCommissionOrderSyncAPIRequest) SetParam0(_param0 *UniverseOrderVo) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaRetailCommissionOrderSyncAPIRequest) GetParam0() *UniverseOrderVo {
    return r._param0
}
