package inventory

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/inventory"
)

/* 
地点关联关系增量编辑 
taobao.location.relation.edit

地点关联关系增量编辑
*/
func TaobaoLocationRelationEdit(clt *core.SDKClient, req *inventory.TaobaoLocationRelationEditRequest, session string) (*inventory.TaobaoLocationRelationEditAPIResponse, error) {
    var resp inventory.TaobaoLocationRelationEditAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
