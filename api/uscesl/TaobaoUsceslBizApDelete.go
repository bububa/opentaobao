package uscesl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/uscesl"
)

/* 
删除价签AP设备 
taobao.uscesl.biz.ap.delete

删除价签AP设备
*/
func TaobaoUsceslBizApDelete(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizApDeleteRequest, session string) (*uscesl.TaobaoUsceslBizApDeleteAPIResponse, error) {
    var resp uscesl.TaobaoUsceslBizApDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
