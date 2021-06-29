package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
专题内容操作列表 
yunos.tvpubadmin.manage.topic.list

获取外部可操作编辑的专题列表
*/
func YunosTvpubadminManageTopicList(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicListRequest, session string) (*tvupadmin.YunosTvpubadminManageTopicListAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminManageTopicListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
