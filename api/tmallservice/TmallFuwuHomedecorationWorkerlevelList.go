package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallFuwuHomedecorationWorkerlevelList 查询工人分层数据接口
// tmall.fuwu.homedecoration.workerlevel.list
//
// 查询工人分层数据接口
func TmallFuwuHomedecorationWorkerlevelList(clt *core.SDKClient, req *tmallservice.TmallFuwuHomedecorationWorkerlevelListAPIRequest, resp *tmallservice.TmallFuwuHomedecorationWorkerlevelListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
