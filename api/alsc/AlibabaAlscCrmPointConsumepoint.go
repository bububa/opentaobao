package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* AlibabaAlscCrmPointConsumepoint
积分抵现
alibaba.alsc.crm.point.consumepoint

积分抵现 */
func AlibabaAlscCrmPointConsumepoint(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointConsumepointAPIRequest, session string) (*alsc.AlibabaAlscCrmPointConsumepointAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmPointConsumepointAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
