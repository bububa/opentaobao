package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsLabelPrint print label
// aliexpress.local.logistics.label.print
//
// print label
func AliexpressLocalLogisticsLabelPrint(ctx context.Context, clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsLabelPrintAPIRequest, resp *logistic.AliexpressLocalLogisticsLabelPrintAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
