package degoperation

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/degoperation"
)

/* 
通用用户抽奖次数限制 
taobao.degoperation.get.by.eventkey

通用用户抽奖次数限制
*/
func TaobaoDegoperationGetByEventkey(clt *core.SDKClient, req *degoperation.TaobaoDegoperationGetByEventkeyAPIRequest, session string) (*degoperation.TaobaoDegoperationGetByEventkeyAPIResponse, error) {
    var resp degoperation.TaobaoDegoperationGetByEventkeyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
