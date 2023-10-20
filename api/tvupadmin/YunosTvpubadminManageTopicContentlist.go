package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicContentlist 查看专题内容列表
// yunos.tvpubadmin.manage.topic.contentlist
//
// 查看专题内容列表
func YunosTvpubadminManageTopicContentlist(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicContentlistAPIRequest, resp *tvupadmin.YunosTvpubadminManageTopicContentlistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
