package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoRegionWarehouseManage 编辑仓库覆盖范围
// taobao.region.warehouse.manage
//
// 编辑仓库覆盖范围
func TaobaoRegionWarehouseManage(clt *core.SDKClient, req *fenxiao.TaobaoRegionWarehouseManageAPIRequest, resp *fenxiao.TaobaoRegionWarehouseManageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
