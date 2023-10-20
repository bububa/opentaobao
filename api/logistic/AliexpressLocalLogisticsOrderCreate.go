package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsOrderCreate create logistics order
// aliexpress.local.logistics.order.create
//
// create logistics order
func AliexpressLocalLogisticsOrderCreate(clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsOrderCreateAPIRequest, resp *logistic.AliexpressLocalLogisticsOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
