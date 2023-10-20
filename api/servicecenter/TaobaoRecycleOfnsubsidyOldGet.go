package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecycleOfnsubsidyOldGet 回收单旧机款及补贴查询
// taobao.recycle.ofnsubsidy.old.get
//
// 回收单旧机款及补贴查询
func TaobaoRecycleOfnsubsidyOldGet(clt *core.SDKClient, req *servicecenter.TaobaoRecycleOfnsubsidyOldGetAPIRequest, resp *servicecenter.TaobaoRecycleOfnsubsidyOldGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
