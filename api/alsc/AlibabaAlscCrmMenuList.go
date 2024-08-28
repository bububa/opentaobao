package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmMenuList 获取特价菜单
// alibaba.alsc.crm.menu.list
//
// 获取特价菜单
func AlibabaAlscCrmMenuList(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmMenuListAPIRequest, resp *alsc.AlibabaAlscCrmMenuListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
