package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseStorePunishInfoQuery 门店处罚信息查询
// alibaba.alihouse.store.punish.info.query
//
// 门店处罚信息查询
func AlibabaAlihouseStorePunishInfoQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseStorePunishInfoQueryAPIRequest, resp *alihouse.AlibabaAlihouseStorePunishInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
