package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuCategoryGetNew （新）ICBU类目树获取接口
// alibaba.icbu.category.get.new
//
// 获取商品发布类目
func AlibabaIcbuCategoryGetNew(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryGetNewAPIRequest, resp *icbu.AlibabaIcbuCategoryGetNewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
