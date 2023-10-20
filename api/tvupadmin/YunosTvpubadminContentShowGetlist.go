package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentshowgetlist 节目审核获取节目列表
// yunos.tvpubadmin.content.show.getlist
//
// 节目审核获取节目列表
func Yunostvpubadmincontentshowgetlist(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentshowgetlistAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentshowgetlistAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentshowgetlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
