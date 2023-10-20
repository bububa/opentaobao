package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoLocationRelationQuery 地点关联关系查询
// taobao.location.relation.query
//
// 地点关联关系查询
// 门店和仓库关联关系查询
func TaobaoLocationRelationQuery(clt *core.SDKClient, req *inventory.TaobaoLocationRelationQueryAPIRequest, resp *inventory.TaobaoLocationRelationQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
