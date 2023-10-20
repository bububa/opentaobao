package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosCommonAuthOperatorInfo 获取当前人员信息
// alibaba.mos.common.auth.operator.info
//
// 获取当前人员信息
func AlibabaMosCommonAuthOperatorInfo(clt *core.SDKClient, req *mos.AlibabaMosCommonAuthOperatorInfoAPIRequest, resp *mos.AlibabaMosCommonAuthOperatorInfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
