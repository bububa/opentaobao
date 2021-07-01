package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
双11到店互动无线端抽奖接口鉴权 
alibaba.interact.wireless.draw

双11到店互动无线端mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 坯子
*/
func AlibabaInteractWirelessDraw(clt *core.SDKClient, req *interact.AlibabaInteractWirelessDrawAPIRequest, session string) (*interact.AlibabaInteractWirelessDrawAPIResponse, error) {
    var resp interact.AlibabaInteractWirelessDrawAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
