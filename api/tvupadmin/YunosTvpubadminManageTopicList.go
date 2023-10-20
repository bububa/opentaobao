package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicList 专题内容操作列表
// yunos.tvpubadmin.manage.topic.list
//
// 获取外部可操作编辑的专题列表
func YunosTvpubadminManageTopicList(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicListAPIRequest, resp *tvupadmin.YunosTvpubadminManageTopicListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
