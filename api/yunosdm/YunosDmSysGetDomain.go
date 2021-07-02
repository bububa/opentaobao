package yunosdm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosdm"
)

// YunosDmSysGetDomain 获取动态域名
// yunos.dm.sys.get.domain
//
// 返回alios ucp后端域名
func YunosDmSysGetDomain(clt *core.SDKClient, req *yunosdm.YunosDmSysGetDomainAPIRequest, session string) (*yunosdm.YunosDmSysGetDomainAPIResponse, error) {
	var resp yunosdm.YunosDmSysGetDomainAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
