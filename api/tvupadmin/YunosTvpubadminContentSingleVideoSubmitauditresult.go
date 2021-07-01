package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
单视频审核提交审核结果 
yunos.tvpubadmin.content.single.video.submitauditresult

单视频审核提交审核结果
*/
func YunosTvpubadminContentSingleVideoSubmitauditresult(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentSingleVideoSubmitauditresultAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentSingleVideoSubmitauditresultAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
