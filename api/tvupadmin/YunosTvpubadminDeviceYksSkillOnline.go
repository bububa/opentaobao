package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
迎客松技能上架接口 
yunos.tvpubadmin.device.yks.skill.online

迎客松技能上架接口
*/
func YunosTvpubadminDeviceYksSkillOnline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillOnlineAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceYksSkillOnlineAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDeviceYksSkillOnlineAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
