package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
修改技能 
yunos.tvpubadmin.device.yks.skill.modify

修改技能
*/
func YunosTvpubadminDeviceYksSkillModify(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillModifyAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceYksSkillModifyAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDeviceYksSkillModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
