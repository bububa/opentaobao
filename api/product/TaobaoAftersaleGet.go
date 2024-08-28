package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoAftersaleGet 查询用户售后服务模板
// taobao.aftersale.get
//
// 查询用户设置的售后服务模板，仅返回标题和id
func TaobaoAftersaleGet(ctx context.Context, clt *core.SDKClient, req *product.TaobaoAftersaleGetAPIRequest, resp *product.TaobaoAftersaleGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
