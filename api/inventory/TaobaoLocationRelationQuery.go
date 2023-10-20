package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// Taobaolocationrelationquery 地点关联关系查询
// taobao.location.relation.query
//
// 地点关联关系查询
// 门店和仓库关联关系查询
func Taobaolocationrelationquery(clt *core.SDKClient, req *inventory.TaobaolocationrelationqueryAPIRequest, session string) (*inventory.TaobaolocationrelationqueryAPIResponse, error) {
	var resp inventory.TaobaolocationrelationqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
