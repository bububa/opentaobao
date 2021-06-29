package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优酷OTT广告素材审核 APIRequest
yunos.tvpubadmin.adm.ott.audit

审核优酷OTT端广告素材
*/
type YunosTvpubadminAdmOttAuditRequest struct {
    model.Params

    // 广告审核内容，json格式
    data   string 

}

func NewYunosTvpubadminAdmOttAuditRequest() *YunosTvpubadminAdmOttAuditRequest{
    return &YunosTvpubadminAdmOttAuditRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminAdmOttAuditRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.adm.ott.audit"
}

func (r YunosTvpubadminAdmOttAuditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminAdmOttAuditRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r YunosTvpubadminAdmOttAuditRequest) GetData() string {
    return r.data
}

