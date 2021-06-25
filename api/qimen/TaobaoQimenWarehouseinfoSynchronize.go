package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
仓库同步接口 
taobao.qimen.warehouseinfo.synchronize

仓库同步接口
*/
func TaobaoQimenWarehouseinfoSynchronize(clt *core.SDKClient, req *qimen.TaobaoQimenWarehouseinfoSynchronizeRequest, session string) (*qimen.TaobaoQimenWarehouseinfoSynchronizeResponse, error) {
    var resp qimen.TaobaoQimenWarehouseinfoSynchronizeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
