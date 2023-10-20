package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoRegionWarehouseQuery 查询仓库覆盖范围
// taobao.region.warehouse.query
//
// 查询仓库覆盖范围
func TaobaoRegionWarehouseQuery(clt *core.SDKClient, req *fenxiao.TaobaoRegionWarehouseQueryAPIRequest, resp *fenxiao.TaobaoRegionWarehouseQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
