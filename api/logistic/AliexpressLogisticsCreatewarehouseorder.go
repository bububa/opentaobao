package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLogisticsCreatewarehouseorder 创建线上物流订单
// aliexpress.logistics.createwarehouseorder
//
// 创建线上发货物流订单
func AliexpressLogisticsCreatewarehouseorder(clt *core.SDKClient, req *logistic.AliexpressLogisticsCreatewarehouseorderAPIRequest, resp *logistic.AliexpressLogisticsCreatewarehouseorderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
