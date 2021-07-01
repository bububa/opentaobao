package cloudgame

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cloudgame"
)

/* 
云游戏混淆用户ID校验 
alibaba.cloudgame.user.mixuserid.check

验证混淆用户ID是否合法
*/
func AlibabaCloudgameUserMixuseridCheck(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameUserMixuseridCheckAPIRequest, session string) (*cloudgame.AlibabaCloudgameUserMixuseridCheckAPIResponse, error) {
    var resp cloudgame.AlibabaCloudgameUserMixuseridCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
