package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtBrandinfoQuery 品牌数据查询
// tmall.nrt.brandinfo.query
//
// 商家获取自己旗舰店授权的品牌id列表
func TmallNrtBrandinfoQuery(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtBrandinfoQueryAPIRequest, resp *nrt.TmallNrtBrandinfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
