package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
查看专题内容列表 
yunos.tvpubadmin.manage.topic.contentlist

查看专题内容列表
*/
func YunosTvpubadminManageTopicContentlist(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicContentlistRequest, session string) (*tvupadmin.YunosTvpubadminManageTopicContentlistAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminManageTopicContentlistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
