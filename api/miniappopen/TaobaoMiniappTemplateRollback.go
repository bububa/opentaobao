package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

/* TaobaoMiniappTemplateRollback
回滚实例化应用
taobao.miniapp.template.rollback

将实例化小程序回滚到指定版本 */
func TaobaoMiniappTemplateRollback(clt *core.SDKClient, req *miniappopen.TaobaoMiniappTemplateRollbackAPIRequest, session string) (*miniappopen.TaobaoMiniappTemplateRollbackAPIResponse, error) {
	var resp miniappopen.TaobaoMiniappTemplateRollbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
