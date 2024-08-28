package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosCommonAuthOperatorInfo 获取当前人员信息
// alibaba.mos.common.auth.operator.info
//
// 获取当前人员信息
func AlibabaMosCommonAuthOperatorInfo(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosCommonAuthOperatorInfoAPIRequest, resp *mos.AlibabaMosCommonAuthOperatorInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
