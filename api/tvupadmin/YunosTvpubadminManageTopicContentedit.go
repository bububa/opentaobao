package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicContentedit 专题关联内容编辑
// yunos.tvpubadmin.manage.topic.contentedit
//
// 编辑专题关联的内容
func YunosTvpubadminManageTopicContentedit(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicContenteditAPIRequest, resp *tvupadmin.YunosTvpubadminManageTopicContenteditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
