package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
更新一个推广计划的平台设置 
taobao.simba.campaign.platform.update

更新一个推广计划的平台设置
*/
func TaobaoSimbaCampaignPlatformUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignPlatformUpdateRequest, session string) (*simba.TaobaoSimbaCampaignPlatformUpdateResponse, error) {
    var resp simba.TaobaoSimbaCampaignPlatformUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
