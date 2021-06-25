package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
删除智能发货引擎仓策略 
cainiao.smartdelivery.strategy.warehouse.i.delete

删除智能发货引擎仓策略
*/
func CainiaoSmartdeliveryStrategyWarehouseIDelete(clt *core.SDKClient, req *waybill.CainiaoSmartdeliveryStrategyWarehouseIDeleteRequest, session string) (*waybill.CainiaoSmartdeliveryStrategyWarehouseIDeleteResponse, error) {
    var resp waybill.CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
