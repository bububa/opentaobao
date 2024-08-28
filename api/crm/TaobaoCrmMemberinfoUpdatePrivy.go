package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMemberinfoUpdatePrivy 编辑会员资料
// taobao.crm.memberinfo.update.privy
//
// 编辑会员的基本资料，接口返回会员信息修改是否成功
func TaobaoCrmMemberinfoUpdatePrivy(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmMemberinfoUpdatePrivyAPIRequest, resp *crm.TaobaoCrmMemberinfoUpdatePrivyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
