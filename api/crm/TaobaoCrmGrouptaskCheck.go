package crm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/crm"
)

/* 
查询分组任务是否完成 
taobao.crm.grouptask.check

检查一个分组上是否有异步任务,异步任务包括1.将一个分组下的所有用户添加到另外一个分组2.将一个分组下的所有用户移动到另外一个分组3.删除某个分组<br/>若分组上有任务则该属性不能被操作。
*/
func TaobaoCrmGrouptaskCheck(clt *core.SDKClient, req *crm.TaobaoCrmGrouptaskCheckRequest, session string) (*crm.TaobaoCrmGrouptaskCheckAPIResponse, error) {
    var resp crm.TaobaoCrmGrouptaskCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
