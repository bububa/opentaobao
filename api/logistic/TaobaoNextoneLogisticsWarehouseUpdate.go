package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
AG退货入仓状态写接口 
taobao.nextone.logistics.warehouse.update

商家上传退货入仓状态给ag
*/
func TaobaoNextoneLogisticsWarehouseUpdate(clt *core.SDKClient, req *logistic.TaobaoNextoneLogisticsWarehouseUpdateRequest, session string) (*logistic.TaobaoNextoneLogisticsWarehouseUpdateResponse, error) {
    var resp logistic.TaobaoNextoneLogisticsWarehouseUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
