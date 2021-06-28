package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
货主仓库资源查询接口 
taobao.qimen.warehouseinfo.query

货主仓库资源查询
*/
func TaobaoQimenWarehouseinfoQuery(clt *core.SDKClient, req *qimen.TaobaoQimenWarehouseinfoQueryRequest, session string) (*qimen.TaobaoQimenWarehouseinfoQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenWarehouseinfoQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
