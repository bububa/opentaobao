package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicEdit 编辑专题
// yunos.tvpubadmin.manage.topic.edit
//
// 编辑专题
func YunosTvpubadminManageTopicEdit(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicEditAPIRequest, resp *tvupadmin.YunosTvpubadminManageTopicEditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
