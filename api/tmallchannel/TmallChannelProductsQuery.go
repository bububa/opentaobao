package tmallchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// TmallChannelProductsQuery 渠道中心-查询产品列表
// tmall.channel.products.query
//
// 渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息
func TmallChannelProductsQuery(ctx context.Context, clt *core.SDKClient, req *tmallchannel.TmallChannelProductsQueryAPIRequest, resp *tmallchannel.TmallChannelProductsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
