package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallFuwuHomedecorationWorkerlevelList 查询工人分层数据接口
// tmall.fuwu.homedecoration.workerlevel.list
//
// 查询工人分层数据接口
func TmallFuwuHomedecorationWorkerlevelList(clt *core.SDKClient, req *tmallservice.TmallFuwuHomedecorationWorkerlevelListAPIRequest, session string) (*tmallservice.TmallFuwuHomedecorationWorkerlevelListAPIResponse, error) {
	var resp tmallservice.TmallFuwuHomedecorationWorkerlevelListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
