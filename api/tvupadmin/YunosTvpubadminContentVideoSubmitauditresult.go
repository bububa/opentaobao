package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
迎客松提交视频审核结果 
yunos.tvpubadmin.content.video.submitauditresult

迎客松提交视频审核结果
*/
func YunosTvpubadminContentVideoSubmitauditresult(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentVideoSubmitauditresultRequest, session string) (*tvupadmin.YunosTvpubadminContentVideoSubmitauditresultAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentVideoSubmitauditresultAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}