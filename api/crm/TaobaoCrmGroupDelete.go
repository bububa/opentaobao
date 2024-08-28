package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmGroupDelete 删除分组
// taobao.crm.group.delete
//
// 将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
func TaobaoCrmGroupDelete(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmGroupDeleteAPIRequest, resp *crm.TaobaoCrmGroupDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
