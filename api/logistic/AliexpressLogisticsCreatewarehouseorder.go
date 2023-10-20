package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Aliexpresslogisticscreatewarehouseorder 创建线上物流订单
// aliexpress.logistics.createwarehouseorder
//
// 创建线上发货物流订单
func Aliexpresslogisticscreatewarehouseorder(clt *core.SDKClient, req *logistic.AliexpresslogisticscreatewarehouseorderAPIRequest, session string) (*logistic.AliexpresslogisticscreatewarehouseorderAPIResponse, error) {
	var resp logistic.AliexpresslogisticscreatewarehouseorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
