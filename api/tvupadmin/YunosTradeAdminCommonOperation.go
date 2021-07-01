package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
交易迎客松通用服务接口 
yunos.trade.admin.common.operation

迎客松交易相关通用接口
*/
func YunosTradeAdminCommonOperation(clt *core.SDKClient, req *tvupadmin.YunosTradeAdminCommonOperationAPIRequest, session string) (*tvupadmin.YunosTradeAdminCommonOperationAPIResponse, error) {
    var resp tvupadmin.YunosTradeAdminCommonOperationAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
