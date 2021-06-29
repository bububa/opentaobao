package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
优酷OTT广告素材审核 
yunos.tvpubadmin.adm.ott.audit

审核优酷OTT端广告素材
*/
func YunosTvpubadminAdmOttAudit(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminAdmOttAuditRequest, session string) (*tvupadmin.YunosTvpubadminAdmOttAuditAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminAdmOttAuditAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
