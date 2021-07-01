package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
批量得到推广计划下的推广单元 
taobao.simba.adgroupsbycampaignid.get

根据推广计划ID分页获取推广计划下的推广单元信息
*/
func TaobaoSimbaAdgroupsbycampaignidGet(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupsbycampaignidGetAPIRequest, session string) (*simba.TaobaoSimbaAdgroupsbycampaignidGetAPIResponse, error) {
    var resp simba.TaobaoSimbaAdgroupsbycampaignidGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
