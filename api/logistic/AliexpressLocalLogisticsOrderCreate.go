package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsOrderCreate create logistics order
// aliexpress.local.logistics.order.create
//
// create logistics order
func AliexpressLocalLogisticsOrderCreate(clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsOrderCreateAPIRequest, session string) (*logistic.AliexpressLocalLogisticsOrderCreateAPIResponse, error) {
	var resp logistic.AliexpressLocalLogisticsOrderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
