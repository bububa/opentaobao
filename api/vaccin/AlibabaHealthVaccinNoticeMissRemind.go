package vaccin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/vaccin"
)

/* 
疫苗漏种提醒 
alibaba.health.vaccin.notice.miss.remind

医生消息提醒适龄儿童按计划接种
*/
func AlibabaHealthVaccinNoticeMissRemind(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeMissRemindRequest, session string) (*vaccin.AlibabaHealthVaccinNoticeMissRemindAPIResponse, error) {
    var resp vaccin.AlibabaHealthVaccinNoticeMissRemindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
