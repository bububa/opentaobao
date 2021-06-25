package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
取得一个推广计划的投放平台设置 
taobao.simba.campaign.platform.get

获得一个推广计划的投放平台设置
*/
func TaobaoSimbaCampaignPlatformGet(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignPlatformGetRequest, session string) (*simba.TaobaoSimbaCampaignPlatformGetResponse, error) {
    var resp simba.TaobaoSimbaCampaignPlatformGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
