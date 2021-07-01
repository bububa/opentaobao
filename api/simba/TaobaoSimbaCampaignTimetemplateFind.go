package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取分时折扣模板 
taobao.simba.campaign.timetemplate.find

批量得到智能推广推广计划下的推广组
*/
func TaobaoSimbaCampaignTimetemplateFind(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignTimetemplateFindAPIRequest, session string) (*simba.TaobaoSimbaCampaignTimetemplateFindAPIResponse, error) {
    var resp simba.TaobaoSimbaCampaignTimetemplateFindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
