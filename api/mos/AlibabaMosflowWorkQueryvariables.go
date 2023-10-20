package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosflowWorkQueryvariables 获取指定流程上下文参数
// alibaba.mosflow.work.queryvariables
//
// 业务查询指定流程上下文内容
func AlibabaMosflowWorkQueryvariables(clt *core.SDKClient, req *mos.AlibabaMosflowWorkQueryvariablesAPIRequest, session string) (*mos.AlibabaMosflowWorkQueryvariablesAPIResponse, error) {
	var resp mos.AlibabaMosflowWorkQueryvariablesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
