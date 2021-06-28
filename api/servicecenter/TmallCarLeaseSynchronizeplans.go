package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
同步租赁方案 
tmall.car.lease.synchronizeplans

租赁公司同步还款计划
*/
func TmallCarLeaseSynchronizeplans(clt *core.SDKClient, req *servicecenter.TmallCarLeaseSynchronizeplansRequest, session string) (*servicecenter.TmallCarLeaseSynchronizeplansAPIResponse, error) {
    var resp servicecenter.TmallCarLeaseSynchronizeplansAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
