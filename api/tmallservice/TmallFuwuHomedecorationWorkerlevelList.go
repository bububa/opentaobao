package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallFuwuHomedecorationWorkerlevelList 查询工人分层数据接口
// tmall.fuwu.homedecoration.workerlevel.list
//
// 查询工人分层数据接口
func TmallFuwuHomedecorationWorkerlevelList(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallFuwuHomedecorationWorkerlevelListAPIRequest, resp *tmallservice.TmallFuwuHomedecorationWorkerlevelListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
