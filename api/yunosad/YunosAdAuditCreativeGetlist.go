package yunosad

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/yunosad"
)

/* 
批量获取创意审核状态 
yunos.ad.audit.creative.getlist

批量获取创意审核状态
*/
func YunosAdAuditCreativeGetlist(clt *core.SDKClient, req *yunosad.YunosAdAuditCreativeGetlistRequest, session string) (*yunosad.YunosAdAuditCreativeGetlistAPIResponse, error) {
    var resp yunosad.YunosAdAuditCreativeGetlistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
