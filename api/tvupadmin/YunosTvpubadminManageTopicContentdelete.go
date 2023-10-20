package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagetopiccontentdelete 删除专题下内容
// yunos.tvpubadmin.manage.topic.contentdelete
//
// 删除专题下内容信息
func Yunostvpubadminmanagetopiccontentdelete(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagetopiccontentdeleteAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagetopiccontentdeleteAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagetopiccontentdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
