package crm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/crm"
)

/* 
修改一个已经存在的分组 
taobao.crm.group.update

修改一个已经存在的分组，接口返回分组的修改是否成功
*/
func TaobaoCrmGroupUpdate(clt *core.SDKClient, req *crm.TaobaoCrmGroupUpdateRequest, session string) (*crm.TaobaoCrmGroupUpdateAPIResponse, error) {
    var resp crm.TaobaoCrmGroupUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
