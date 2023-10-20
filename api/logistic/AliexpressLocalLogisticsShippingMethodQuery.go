package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsShippingMethodQuery query shipping method
// aliexpress.local.logistics.shipping.method.query
//
// query shipping method
func AliexpressLocalLogisticsShippingMethodQuery(clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsShippingMethodQueryAPIRequest, resp *logistic.AliexpressLocalLogisticsShippingMethodQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
