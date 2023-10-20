package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Aliexpresslocallogisticsorderinfoquery query order details
// aliexpress.local.logistics.order.info.query
//
// query order details
func Aliexpresslocallogisticsorderinfoquery(clt *core.SDKClient, req *logistic.AliexpresslocallogisticsorderinfoqueryAPIRequest, session string) (*logistic.AliexpresslocallogisticsorderinfoqueryAPIResponse, error) {
	var resp logistic.AliexpresslocallogisticsorderinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
