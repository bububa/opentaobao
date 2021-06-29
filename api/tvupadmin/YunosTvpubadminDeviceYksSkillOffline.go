package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
技能下架 
yunos.tvpubadmin.device.yks.skill.offline

迎客松平台技能下架
*/
func YunosTvpubadminDeviceYksSkillOffline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillOfflineRequest, session string) (*tvupadmin.YunosTvpubadminDeviceYksSkillOfflineAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDeviceYksSkillOfflineAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
