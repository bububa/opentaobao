package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
编辑专题 
yunos.tvpubadmin.manage.topic.edit

编辑专题
*/
func YunosTvpubadminManageTopicEdit(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicEditRequest, session string) (*tvupadmin.YunosTvpubadminManageTopicEditAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminManageTopicEditAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
