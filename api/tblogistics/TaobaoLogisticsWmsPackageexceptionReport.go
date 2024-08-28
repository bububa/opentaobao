package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsWmsPackageexceptionReport 无主件回告
// taobao.logistics.wms.packageexception.report
//
// 无主件回告
func TaobaoLogisticsWmsPackageexceptionReport(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoLogisticsWmsPackageexceptionReportAPIRequest, resp *tblogistics.TaobaoLogisticsWmsPackageexceptionReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
