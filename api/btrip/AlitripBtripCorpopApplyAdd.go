package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
【商旅】isv添加审批单 
alitrip.btrip.corpop.apply.add

【商旅】isv添加审批单
*/
func AlitripBtripCorpopApplyAdd(clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplyAddRequest, session string) (*btrip.AlitripBtripCorpopApplyAddAPIResponse, error) {
    var resp btrip.AlitripBtripCorpopApplyAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
