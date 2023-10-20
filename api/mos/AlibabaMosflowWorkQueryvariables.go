package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosflowworkqueryvariables 获取指定流程上下文参数
// alibaba.mosflow.work.queryvariables
//
// 业务查询指定流程上下文内容
func Alibabamosflowworkqueryvariables(clt *core.SDKClient, req *mos.AlibabamosflowworkqueryvariablesAPIRequest, session string) (*mos.AlibabamosflowworkqueryvariablesAPIResponse, error) {
	var resp mos.AlibabamosflowworkqueryvariablesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
