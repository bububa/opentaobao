package crm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/crm"
)

/* 
删除分组 
taobao.crm.group.delete

将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
*/
func TaobaoCrmGroupDelete(clt *core.SDKClient, req *crm.TaobaoCrmGroupDeleteAPIRequest, session string) (*crm.TaobaoCrmGroupDeleteAPIResponse, error) {
    var resp crm.TaobaoCrmGroupDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
