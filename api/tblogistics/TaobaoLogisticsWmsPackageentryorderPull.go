package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsWmsPackageentryorderPull 包裹入库单拉单
// taobao.logistics.wms.packageentryorder.pull
//
// 包裹入库单拉单
func TaobaoLogisticsWmsPackageentryorderPull(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsWmsPackageentryorderPullAPIRequest, session string) (*tblogistics.TaobaoLogisticsWmsPackageentryorderPullAPIResponse, error) {
	var resp tblogistics.TaobaoLogisticsWmsPackageentryorderPullAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
