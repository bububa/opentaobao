package crm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/crm"
)

/* 
编辑会员资料 
taobao.crm.memberinfo.update

编辑会员的基本资料，接口返回会员信息修改是否成功
*/
func TaobaoCrmMemberinfoUpdate(clt *core.SDKClient, req *crm.TaobaoCrmMemberinfoUpdateRequest, session string) (*crm.TaobaoCrmMemberinfoUpdateResponse, error) {
    var resp crm.TaobaoCrmMemberinfoUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
