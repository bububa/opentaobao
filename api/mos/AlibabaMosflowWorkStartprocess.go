package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosflowWorkStartprocess 发起流程
// alibaba.mosflow.work.startprocess
//
// 业务发起流程审批
func AlibabaMosflowWorkStartprocess(clt *core.SDKClient, req *mos.AlibabaMosflowWorkStartprocessAPIRequest, resp *mos.AlibabaMosflowWorkStartprocessAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
