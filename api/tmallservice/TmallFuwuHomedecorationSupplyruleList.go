package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallFuwuHomedecorationSupplyruleList 查询供给规则接口
// tmall.fuwu.homedecoration.supplyrule.list
//
// 查询供给规则接口
func TmallFuwuHomedecorationSupplyruleList(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallFuwuHomedecorationSupplyruleListAPIRequest, resp *tmallservice.TmallFuwuHomedecorationSupplyruleListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
