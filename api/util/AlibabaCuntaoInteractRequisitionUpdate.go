package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
更新物料制作状态 
alibaba.cuntao.interact.requisition.update

村淘物料下沉，更新物料制作状态
*/
func AlibabaCuntaoInteractRequisitionUpdate(clt *core.SDKClient, req *util.AlibabaCuntaoInteractRequisitionUpdateAPIRequest, session string) (*util.AlibabaCuntaoInteractRequisitionUpdateAPIResponse, error) {
    var resp util.AlibabaCuntaoInteractRequisitionUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
