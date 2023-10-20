package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsLabelPrint print label
// aliexpress.local.logistics.label.print
//
// print label
func AliexpressLocalLogisticsLabelPrint(clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsLabelPrintAPIRequest, resp *logistic.AliexpressLocalLogisticsLabelPrintAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
