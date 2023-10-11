package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsWmsPackageexceptionReport 无主件回告
// taobao.logistics.wms.packageexception.report
//
// 无主件回告
func TaobaoLogisticsWmsPackageexceptionReport(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsWmsPackageexceptionReportAPIRequest, session string) (*tblogistics.TaobaoLogisticsWmsPackageexceptionReportAPIResponse, error) {
	var resp tblogistics.TaobaoLogisticsWmsPackageexceptionReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
