package vaccin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/vaccin"
)

/* 
福州疫苗签到成功通知 
alibaba.health.vaccin.notice.order.sign

福州疫苗用户签到成功记录
*/
func AlibabaHealthVaccinNoticeOrderSign(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeOrderSignRequest, session string) (*vaccin.AlibabaHealthVaccinNoticeOrderSignAPIResponse, error) {
    var resp vaccin.AlibabaHealthVaccinNoticeOrderSignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
