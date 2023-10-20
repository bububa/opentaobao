package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsWmsPackageentryorderPull 包裹入库单拉单
// taobao.logistics.wms.packageentryorder.pull
//
// 包裹入库单拉单
func TaobaoLogisticsWmsPackageentryorderPull(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsWmsPackageentryorderPullAPIRequest, resp *tblogistics.TaobaoLogisticsWmsPackageentryorderPullAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
