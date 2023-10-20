package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosflowWorkStartprocess 发起流程
// alibaba.mosflow.work.startprocess
//
// 业务发起流程审批
func AlibabaMosflowWorkStartprocess(clt *core.SDKClient, req *mos.AlibabaMosflowWorkStartprocessAPIRequest, session string) (*mos.AlibabaMosflowWorkStartprocessAPIResponse, error) {
	var resp mos.AlibabaMosflowWorkStartprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
