package yunosdm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosdm"
)

// YunosDmSysGetDomain 获取动态域名
// yunos.dm.sys.get.domain
//
// 返回alios ucp后端域名
func YunosDmSysGetDomain(ctx context.Context, clt *core.SDKClient, req *yunosdm.YunosDmSysGetDomainAPIRequest, resp *yunosdm.YunosDmSysGetDomainAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
