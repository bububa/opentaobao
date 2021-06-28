package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
供应商获取物料申请单列表 
alibaba.cuntao.interact.requisition.get

供应商获取物料申请单列表
*/
func AlibabaCuntaoInteractRequisitionGet(clt *core.SDKClient, req *util.AlibabaCuntaoInteractRequisitionGetRequest, session string) (*util.AlibabaCuntaoInteractRequisitionGetAPIResponse, error) {
    var resp util.AlibabaCuntaoInteractRequisitionGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
