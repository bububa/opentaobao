package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosBunkBunkinfoQuerybunk 根据合同号查询铺位信息
// alibaba.mos.bunk.bunkinfo.querybunk
//
// 根据合同号查询铺位信息
func AlibabaMosBunkBunkinfoQuerybunk(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosBunkBunkinfoQuerybunkAPIRequest, resp *mos.AlibabaMosBunkBunkinfoQuerybunkAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
