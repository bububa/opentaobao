package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosflowworkstartprocess 发起流程
// alibaba.mosflow.work.startprocess
//
// 业务发起流程审批
func Alibabamosflowworkstartprocess(clt *core.SDKClient, req *mos.AlibabamosflowworkstartprocessAPIRequest, session string) (*mos.AlibabamosflowworkstartprocessAPIResponse, error) {
	var resp mos.AlibabamosflowworkstartprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
