package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecycleOfnpreredpacketTpdeductsuccess 回收商同步前置补贴红包的代扣成功事件
// taobao.recycle.ofnpreredpacket.tpdeductsuccess
//
// 回收商-&gt;天猫后端，同步前置补贴红包的代扣成功事件
func TaobaoRecycleOfnpreredpacketTpdeductsuccess(clt *core.SDKClient, req *servicecenter.TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest, resp *servicecenter.TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
