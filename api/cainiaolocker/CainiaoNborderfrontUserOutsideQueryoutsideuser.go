package cainiaolocker

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoNborderfrontUserOutsideQueryoutsideuser 查询外部小件员休息
// cainiao.nborderfront.user.outside.queryoutsideuser
//
// 采用SPI方式查询外部公司的小件员信息
func CainiaoNborderfrontUserOutsideQueryoutsideuser(ctx context.Context, clt *core.SDKClient, req *cainiaolocker.CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest, resp *cainiaolocker.CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
