package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmPointChkpntbypay 校验支付链路中的积分抵扣是否合法
// alibaba.alsc.crm.point.chkpntbypay
//
// 校验支付链路中的积分抵扣是否合法
func AlibabaAlscCrmPointChkpntbypay(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointChkpntbypayAPIRequest, resp *alsc.AlibabaAlscCrmPointChkpntbypayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
