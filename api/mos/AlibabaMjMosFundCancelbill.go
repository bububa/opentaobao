package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjMosFundCancelbill 取消付款单
// alibaba.mj.mos.fund.cancelbill
//
// 取消付款单
func AlibabaMjMosFundCancelbill(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjMosFundCancelbillAPIRequest, resp *mos.AlibabaMjMosFundCancelbillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
