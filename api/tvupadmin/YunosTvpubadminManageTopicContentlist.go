package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagetopiccontentlist 查看专题内容列表
// yunos.tvpubadmin.manage.topic.contentlist
//
// 查看专题内容列表
func Yunostvpubadminmanagetopiccontentlist(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagetopiccontentlistAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagetopiccontentlistAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagetopiccontentlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
