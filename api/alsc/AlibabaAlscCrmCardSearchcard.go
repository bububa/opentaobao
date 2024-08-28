package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardSearchcard 搜索卡实例列表(支持号段查询)
// alibaba.alsc.crm.card.searchcard
//
// 搜索卡实例列表(支持号段查询)
func AlibabaAlscCrmCardSearchcard(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardSearchcardAPIRequest, resp *alsc.AlibabaAlscCrmCardSearchcardAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
