package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicContentlist 查看专题内容列表
// yunos.tvpubadmin.manage.topic.contentlist
//
// 查看专题内容列表
func YunosTvpubadminManageTopicContentlist(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicContentlistAPIRequest, resp *tvupadmin.YunosTvpubadminManageTopicContentlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
