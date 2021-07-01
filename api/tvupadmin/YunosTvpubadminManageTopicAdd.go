package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
新增专题 
yunos.tvpubadmin.manage.topic.add

专题新增
*/
func YunosTvpubadminManageTopicAdd(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicAddAPIRequest, session string) (*tvupadmin.YunosTvpubadminManageTopicAddAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminManageTopicAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
