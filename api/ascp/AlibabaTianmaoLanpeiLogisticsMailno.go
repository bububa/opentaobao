package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoLanpeiLogisticsMailno 阿里巴巴.天猫家装.揽配.物流.获取运单号
// alibaba.tianmao.lanpei.logistics.mailno
//
// 阿里巴巴.天猫家装.揽配.物流.获取运单号
func AlibabaTianmaoLanpeiLogisticsMailno(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest, resp *ascp.AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
