package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
变更项目 
alitrip.btrip.project.modify

变更项目
*/
func AlitripBtripProjectModify(clt *core.SDKClient, req *btrip.AlitripBtripProjectModifyAPIRequest, session string) (*btrip.AlitripBtripProjectModifyAPIResponse, error) {
    var resp btrip.AlitripBtripProjectModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
