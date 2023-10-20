package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsReportShipped report as shipped
// aliexpress.local.logistics.report.shipped
//
// report as shipped
func AliexpressLocalLogisticsReportShipped(clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsReportShippedAPIRequest, resp *logistic.AliexpressLocalLogisticsReportShippedAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
