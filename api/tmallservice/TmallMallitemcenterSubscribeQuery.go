package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMallitemcenterSubscribeQuery 天猫服务订购信息查询接口
// tmall.mallitemcenter.subscribe.query
//
// 查询商家服务订购信息
func TmallMallitemcenterSubscribeQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallMallitemcenterSubscribeQueryAPIRequest, resp *tmallservice.TmallMallitemcenterSubscribeQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
