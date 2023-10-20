package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoLocationRelationEdit 地点关联关系增量编辑
// taobao.location.relation.edit
//
// 地点关联关系增量编辑
func TaobaoLocationRelationEdit(clt *core.SDKClient, req *inventory.TaobaoLocationRelationEditAPIRequest, resp *inventory.TaobaoLocationRelationEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
