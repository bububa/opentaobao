package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// AlibabaRetailCommissionOrderSync 分佣数据传输
// alibaba.retail.commission.order.sync
//
// 同步分佣结果
func AlibabaRetailCommissionOrderSync(clt *core.SDKClient, req *omniorder.AlibabaRetailCommissionOrderSyncAPIRequest, resp *omniorder.AlibabaRetailCommissionOrderSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
