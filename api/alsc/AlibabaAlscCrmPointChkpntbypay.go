package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmPointChkpntbypay 校验支付链路中的积分抵扣是否合法
// alibaba.alsc.crm.point.chkpntbypay
//
// 校验支付链路中的积分抵扣是否合法
func AlibabaAlscCrmPointChkpntbypay(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointChkpntbypayAPIRequest, resp *alsc.AlibabaAlscCrmPointChkpntbypayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
