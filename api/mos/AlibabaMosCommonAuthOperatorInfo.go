package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosCommonAuthOperatorInfo 获取当前人员信息
// alibaba.mos.common.auth.operator.info
//
// 获取当前人员信息
func AlibabaMosCommonAuthOperatorInfo(clt *core.SDKClient, req *mos.AlibabaMosCommonAuthOperatorInfoAPIRequest, session string) (*mos.AlibabaMosCommonAuthOperatorInfoAPIResponse, error) {
	var resp mos.AlibabaMosCommonAuthOperatorInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
