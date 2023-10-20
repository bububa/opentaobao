package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoregionwarehousemanage 编辑仓库覆盖范围
// taobao.region.warehouse.manage
//
// 编辑仓库覆盖范围
func Taobaoregionwarehousemanage(clt *core.SDKClient, req *fenxiao.TaobaoregionwarehousemanageAPIRequest, session string) (*fenxiao.TaobaoregionwarehousemanageAPIResponse, error) {
	var resp fenxiao.TaobaoregionwarehousemanageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
