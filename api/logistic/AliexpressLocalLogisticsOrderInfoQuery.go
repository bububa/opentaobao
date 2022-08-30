package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsOrderInfoQuery query order details
// aliexpress.local.logistics.order.info.query
//
// query order details
func AliexpressLocalLogisticsOrderInfoQuery(clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsOrderInfoQueryAPIRequest, session string) (*logistic.AliexpressLocalLogisticsOrderInfoQueryAPIResponse, error) {
	var resp logistic.AliexpressLocalLogisticsOrderInfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
