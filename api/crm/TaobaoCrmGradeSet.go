package crm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/crm"
)

/* 
卖家设置等级规则 
taobao.crm.grade.set

设置等级信息，可以设置层级等级，也可以单独设置一个等级。出于安全原因，折扣现最低只能设置到700即7折。
*/
func TaobaoCrmGradeSet(clt *core.SDKClient, req *crm.TaobaoCrmGradeSetRequest, session string) (*crm.TaobaoCrmGradeSetAPIResponse, error) {
    var resp crm.TaobaoCrmGradeSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
