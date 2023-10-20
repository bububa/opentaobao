package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaorecycleofnpreredpackettpdeductsuccess 回收商同步前置补贴红包的代扣成功事件
// taobao.recycle.ofnpreredpacket.tpdeductsuccess
//
// 回收商-&gt;天猫后端，同步前置补贴红包的代扣成功事件
func Taobaorecycleofnpreredpackettpdeductsuccess(clt *core.SDKClient, req *servicecenter.TaobaorecycleofnpreredpackettpdeductsuccessAPIRequest, session string) (*servicecenter.TaobaorecycleofnpreredpackettpdeductsuccessAPIResponse, error) {
	var resp servicecenter.TaobaorecycleofnpreredpackettpdeductsuccessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
