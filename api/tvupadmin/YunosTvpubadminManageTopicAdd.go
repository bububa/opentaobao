package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicAdd 新增专题
// yunos.tvpubadmin.manage.topic.add
//
// 专题新增
func YunosTvpubadminManageTopicAdd(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicAddAPIRequest, resp *tvupadmin.YunosTvpubadminManageTopicAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
