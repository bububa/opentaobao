package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
天猫服务平台服务核销 
tmall.service.code.consume

天猫服务平台－服务核销
*/
func TmallServiceCodeConsume(clt *core.SDKClient, req *tmallservice.TmallServiceCodeConsumeRequest, session string) (*tmallservice.TmallServiceCodeConsumeAPIResponse, error) {
    var resp tmallservice.TmallServiceCodeConsumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
