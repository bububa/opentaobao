package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsWmsPackageentryorderConfirm 包裹入库单确认
// taobao.logistics.wms.packageentryorder.confirm
//
// 包裹入库单确认
func TaobaoLogisticsWmsPackageentryorderConfirm(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest, session string) (*tblogistics.TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse, error) {
	var resp tblogistics.TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
