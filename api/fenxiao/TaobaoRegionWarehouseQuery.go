package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
查询仓库覆盖范围 
taobao.region.warehouse.query

查询仓库覆盖范围
*/
func TaobaoRegionWarehouseQuery(clt *core.SDKClient, req *fenxiao.TaobaoRegionWarehouseQueryRequest, session string) (*fenxiao.TaobaoRegionWarehouseQueryAPIResponse, error) {
    var resp fenxiao.TaobaoRegionWarehouseQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
