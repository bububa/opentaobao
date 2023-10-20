package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsWmsPackagedeliveryorderConfirm 包裹出库单确认
// taobao.logistics.wms.packagedeliveryorder.confirm
//
// 包裹出库单确认
func TaobaoLogisticsWmsPackagedeliveryorderConfirm(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest, session string) (*tblogistics.TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse, error) {
	var resp tblogistics.TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
