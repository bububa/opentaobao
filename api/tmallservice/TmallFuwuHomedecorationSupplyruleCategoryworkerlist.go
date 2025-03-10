package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallFuwuHomedecorationSupplyruleCategoryworkerlist 基于规则查品牌品类工人接口
// tmall.fuwu.homedecoration.supplyrule.categoryworkerlist
//
// 基于规则查品牌品类工人接口
func TmallFuwuHomedecorationSupplyruleCategoryworkerlist(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIRequest, resp *tmallservice.TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
