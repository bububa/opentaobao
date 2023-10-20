package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Aliexpresslocallogisticsreportshipped report as shipped
// aliexpress.local.logistics.report.shipped
//
// report as shipped
func Aliexpresslocallogisticsreportshipped(clt *core.SDKClient, req *logistic.AliexpresslocallogisticsreportshippedAPIRequest, session string) (*logistic.AliexpresslocallogisticsreportshippedAPIResponse, error) {
	var resp logistic.AliexpresslocallogisticsreportshippedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
