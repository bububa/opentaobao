package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Aliexpresslocallogisticsshippingmethodquery query shipping method
// aliexpress.local.logistics.shipping.method.query
//
// query shipping method
func Aliexpresslocallogisticsshippingmethodquery(clt *core.SDKClient, req *logistic.AliexpresslocallogisticsshippingmethodqueryAPIRequest, session string) (*logistic.AliexpresslocallogisticsshippingmethodqueryAPIResponse, error) {
	var resp logistic.AliexpresslocallogisticsshippingmethodqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
