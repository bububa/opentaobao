package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoLanpeiUopCreate 阿里巴巴.天猫家装.揽配.履约订单.创建
// alibaba.tianmao.lanpei.uop.create
//
// 阿里巴巴.天猫家装.揽配.履约订单.创建
func AlibabaTianmaoLanpeiUopCreate(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaTianmaoLanpeiUopCreateAPIRequest, resp *ascp.AlibabaTianmaoLanpeiUopCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
