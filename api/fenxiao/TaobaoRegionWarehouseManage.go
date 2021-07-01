package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
编辑仓库覆盖范围 
taobao.region.warehouse.manage

编辑仓库覆盖范围
*/
func TaobaoRegionWarehouseManage(clt *core.SDKClient, req *fenxiao.TaobaoRegionWarehouseManageAPIRequest, session string) (*fenxiao.TaobaoRegionWarehouseManageAPIResponse, error) {
    var resp fenxiao.TaobaoRegionWarehouseManageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
