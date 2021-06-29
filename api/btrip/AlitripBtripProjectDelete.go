package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
删除项目 
alitrip.btrip.project.delete

删除项目
*/
func AlitripBtripProjectDelete(clt *core.SDKClient, req *btrip.AlitripBtripProjectDeleteRequest, session string) (*btrip.AlitripBtripProjectDeleteAPIResponse, error) {
    var resp btrip.AlitripBtripProjectDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
