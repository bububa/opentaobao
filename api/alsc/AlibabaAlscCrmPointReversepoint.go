package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmPointReversepoint 积分消费回退
// alibaba.alsc.crm.point.reversepoint
//
// 积分消费回退
func AlibabaAlscCrmPointReversepoint(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointReversepointAPIRequest, session string) (*alsc.AlibabaAlscCrmPointReversepointAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmPointReversepointAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
