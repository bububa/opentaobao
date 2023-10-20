package yunosdm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosdm"
)

// YunosDmSysGetDomain 获取动态域名
// yunos.dm.sys.get.domain
//
// 返回alios ucp后端域名
func YunosDmSysGetDomain(clt *core.SDKClient, req *yunosdm.YunosDmSysGetDomainAPIRequest, resp *yunosdm.YunosDmSysGetDomainAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
