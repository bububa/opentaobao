package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
添加项目 
alitrip.btrip.project.add

添加项目
*/
func AlitripBtripProjectAdd(clt *core.SDKClient, req *btrip.AlitripBtripProjectAddRequest, session string) (*btrip.AlitripBtripProjectAddAPIResponse, error) {
    var resp btrip.AlitripBtripProjectAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
