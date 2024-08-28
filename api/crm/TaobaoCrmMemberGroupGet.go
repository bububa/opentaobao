package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMemberGroupGet 获取买家身上的标签
// taobao.crm.member.group.get
//
// 获取买家身上的标签，不返回标签的总人数
func TaobaoCrmMemberGroupGet(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmMemberGroupGetAPIRequest, resp *crm.TaobaoCrmMemberGroupGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
