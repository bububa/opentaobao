package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
迎客松节目设置免审开关 
yunos.tvpubadmin.content.show.setexemptaudit

迎客松节目设置免审开关
*/
func YunosTvpubadminContentShowSetexemptaudit(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowSetexemptauditRequest, session string) (*tvupadmin.YunosTvpubadminContentShowSetexemptauditAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentShowSetexemptauditAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
