package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Aliexpresslocallogisticsordercreate create logistics order
// aliexpress.local.logistics.order.create
//
// create logistics order
func Aliexpresslocallogisticsordercreate(clt *core.SDKClient, req *logistic.AliexpresslocallogisticsordercreateAPIRequest, session string) (*logistic.AliexpresslocallogisticsordercreateAPIResponse, error) {
	var resp logistic.AliexpresslocallogisticsordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
