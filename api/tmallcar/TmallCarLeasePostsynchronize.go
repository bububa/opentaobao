package tmallcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcar"
)

/* 
天猫开新车租后信息同步 
tmall.car.lease.postsynchronize

商家同步天猫开新车租后方案
*/
func TmallCarLeasePostsynchronize(clt *core.SDKClient, req *tmallcar.TmallCarLeasePostsynchronizeAPIRequest, session string) (*tmallcar.TmallCarLeasePostsynchronizeAPIResponse, error) {
    var resp tmallcar.TmallCarLeasePostsynchronizeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
