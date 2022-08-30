package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsShippingMethodQuery query shipping method
// aliexpress.local.logistics.shipping.method.query
//
// query shipping method
func AliexpressLocalLogisticsShippingMethodQuery(clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsShippingMethodQueryAPIRequest, session string) (*logistic.AliexpressLocalLogisticsShippingMethodQueryAPIResponse, error) {
	var resp logistic.AliexpressLocalLogisticsShippingMethodQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
