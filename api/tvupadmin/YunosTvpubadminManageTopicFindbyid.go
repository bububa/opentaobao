package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicFindbyid 根据id获取专题信息
// yunos.tvpubadmin.manage.topic.findbyid
//
// 根据id获取专题信息
func YunosTvpubadminManageTopicFindbyid(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicFindbyidAPIRequest, resp *tvupadmin.YunosTvpubadminManageTopicFindbyidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
