package vaccin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/vaccin"
)

/* 
支付宝疫苗绑定接种人 
alibaba.health.vaccin.notice.user.bind

支付宝疫苗绑定接种人
*/
func AlibabaHealthVaccinNoticeUserBind(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeUserBindAPIRequest, session string) (*vaccin.AlibabaHealthVaccinNoticeUserBindAPIResponse, error) {
    var resp vaccin.AlibabaHealthVaccinNoticeUserBindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
