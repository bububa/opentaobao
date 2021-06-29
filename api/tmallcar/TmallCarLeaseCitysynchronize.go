package tmallcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcar"
)

/* 
天猫开新车租后分期城市信息同步 
tmall.car.lease.citysynchronize

天猫开新车租后分期城市信息同步
*/
func TmallCarLeaseCitysynchronize(clt *core.SDKClient, req *tmallcar.TmallCarLeaseCitysynchronizeRequest, session string) (*tmallcar.TmallCarLeaseCitysynchronizeAPIResponse, error) {
    var resp tmallcar.TmallCarLeaseCitysynchronizeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
