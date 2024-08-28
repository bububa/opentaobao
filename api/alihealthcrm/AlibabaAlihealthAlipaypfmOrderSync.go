package alihealthcrm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthAlipaypfmOrderSync 订单数据回传接口
// alibaba.alihealth.alipaypfm.order.sync
//
// 订单数据回传接口，各个isv通过我们渠道产生订单需要回传进行统计
func AlibabaAlihealthAlipaypfmOrderSync(ctx context.Context, clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthAlipaypfmOrderSyncAPIRequest, resp *alihealthcrm.AlibabaAlihealthAlipaypfmOrderSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
