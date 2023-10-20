package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicContentdelete 删除专题下内容
// yunos.tvpubadmin.manage.topic.contentdelete
//
// 删除专题下内容信息
func YunosTvpubadminManageTopicContentdelete(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicContentdeleteAPIRequest, resp *tvupadmin.YunosTvpubadminManageTopicContentdeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
