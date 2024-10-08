package guoguo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// CainiaoGuoguoCpNborderfrontrUpdateuser 小件员信息变更
// cainiao.guoguo.cp.nborderfrontr.updateuser
//
// 小件员信息变更
func CainiaoGuoguoCpNborderfrontrUpdateuser(ctx context.Context, clt *core.SDKClient, req *guoguo.CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest, resp *guoguo.CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
