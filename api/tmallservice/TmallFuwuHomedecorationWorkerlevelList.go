package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallfuwuhomedecorationworkerlevellist 查询工人分层数据接口
// tmall.fuwu.homedecoration.workerlevel.list
//
// 查询工人分层数据接口
func Tmallfuwuhomedecorationworkerlevellist(clt *core.SDKClient, req *tmallservice.TmallfuwuhomedecorationworkerlevellistAPIRequest, session string) (*tmallservice.TmallfuwuhomedecorationworkerlevellistAPIResponse, error) {
	var resp tmallservice.TmallfuwuhomedecorationworkerlevellistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
