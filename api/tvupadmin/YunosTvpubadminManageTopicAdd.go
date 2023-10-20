package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageTopicAdd 新增专题
// yunos.tvpubadmin.manage.topic.add
//
// 专题新增
func YunosTvpubadminManageTopicAdd(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicAddAPIRequest, resp *tvupadmin.YunosTvpubadminManageTopicAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
