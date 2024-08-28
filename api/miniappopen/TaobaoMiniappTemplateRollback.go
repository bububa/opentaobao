package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappTemplateRollback 回滚实例化应用
// taobao.miniapp.template.rollback
//
// 将实例化小程序回滚到指定版本
func TaobaoMiniappTemplateRollback(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappTemplateRollbackAPIRequest, resp *miniappopen.TaobaoMiniappTemplateRollbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
