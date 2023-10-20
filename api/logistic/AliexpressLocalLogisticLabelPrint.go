package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticLabelPrint 物流打印面单
// aliexpress.local.logistic.label.print
//
// 物流打印面单
func AliexpressLocalLogisticLabelPrint(clt *core.SDKClient, req *logistic.AliexpressLocalLogisticLabelPrintAPIRequest, resp *logistic.AliexpressLocalLogisticLabelPrintAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
