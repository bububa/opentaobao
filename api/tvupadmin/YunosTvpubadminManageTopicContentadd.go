package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicContentadd 专题新增内容
// yunos.tvpubadmin.manage.topic.contentadd
//
// 专题新增内容
func YunosTvpubadminManageTopicContentadd(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicContentaddAPIRequest, resp *tvupadmin.YunosTvpubadminManageTopicContentaddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
