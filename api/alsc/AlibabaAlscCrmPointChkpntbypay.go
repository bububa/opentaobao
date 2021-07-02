package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmPointChkpntbypay 校验支付链路中的积分抵扣是否合法
// alibaba.alsc.crm.point.chkpntbypay
//
// 校验支付链路中的积分抵扣是否合法
func AlibabaAlscCrmPointChkpntbypay(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointChkpntbypayAPIRequest, session string) (*alsc.AlibabaAlscCrmPointChkpntbypayAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmPointChkpntbypayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
