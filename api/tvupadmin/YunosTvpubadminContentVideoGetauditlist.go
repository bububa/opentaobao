package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentvideogetauditlist 迎客松视频审核记录查询
// yunos.tvpubadmin.content.video.getauditlist
//
// 迎客松视频审核记录查询
func Yunostvpubadmincontentvideogetauditlist(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentvideogetauditlistAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentvideogetauditlistAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentvideogetauditlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
