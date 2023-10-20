package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicEdit 编辑专题
// yunos.tvpubadmin.manage.topic.edit
//
// 编辑专题
func YunosTvpubadminManageTopicEdit(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicEditAPIRequest, resp *tvupadmin.YunosTvpubadminManageTopicEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
