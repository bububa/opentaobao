package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsLabelPrint print label
// aliexpress.local.logistics.label.print
//
// print label
func AliexpressLocalLogisticsLabelPrint(clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsLabelPrintAPIRequest, session string) (*logistic.AliexpressLocalLogisticsLabelPrintAPIResponse, error) {
	var resp logistic.AliexpressLocalLogisticsLabelPrintAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
