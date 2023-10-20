package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsWmsPackagedeliveryorderPull 包裹出库单拉单
// taobao.logistics.wms.packagedeliveryorder.pull
//
// 包裹出库单拉单
func TaobaoLogisticsWmsPackagedeliveryorderPull(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest, session string) (*tblogistics.TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse, error) {
	var resp tblogistics.TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
