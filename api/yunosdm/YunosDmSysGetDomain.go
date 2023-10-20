package yunosdm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosdm"
)

// Yunosdmsysgetdomain 获取动态域名
// yunos.dm.sys.get.domain
//
// 返回alios ucp后端域名
func Yunosdmsysgetdomain(clt *core.SDKClient, req *yunosdm.YunosdmsysgetdomainAPIRequest, session string) (*yunosdm.YunosdmsysgetdomainAPIResponse, error) {
	var resp yunosdm.YunosdmsysgetdomainAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
