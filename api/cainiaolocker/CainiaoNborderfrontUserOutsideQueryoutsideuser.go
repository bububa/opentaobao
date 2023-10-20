package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoNborderfrontUserOutsideQueryoutsideuser 查询外部小件员休息
// cainiao.nborderfront.user.outside.queryoutsideuser
//
// 采用SPI方式查询外部公司的小件员信息
func CainiaoNborderfrontUserOutsideQueryoutsideuser(clt *core.SDKClient, req *cainiaolocker.CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest, resp *cainiaolocker.CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
