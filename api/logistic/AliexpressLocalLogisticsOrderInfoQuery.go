package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsOrderInfoQuery query order details
// aliexpress.local.logistics.order.info.query
//
// query order details
func AliexpressLocalLogisticsOrderInfoQuery(clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsOrderInfoQueryAPIRequest, resp *logistic.AliexpressLocalLogisticsOrderInfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
