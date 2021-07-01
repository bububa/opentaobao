package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
根据设备id获取技能列表 
yunos.tvpubadmin.device.yks.skills

根据设备id获取技能列表
*/
func YunosTvpubadminDeviceYksSkills(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillsAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceYksSkillsAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDeviceYksSkillsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
