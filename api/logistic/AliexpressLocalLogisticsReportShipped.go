package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsReportShipped report as shipped
// aliexpress.local.logistics.report.shipped
//
// report as shipped
func AliexpressLocalLogisticsReportShipped(clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsReportShippedAPIRequest, session string) (*logistic.AliexpressLocalLogisticsReportShippedAPIResponse, error) {
	var resp logistic.AliexpressLocalLogisticsReportShippedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
