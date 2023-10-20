package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecycleOfnpreredpacketGet 服务商查询前置补贴红包的最新数据
// taobao.recycle.ofnpreredpacket.get
//
// 服务商查询前置补贴红包的最新数据
func TaobaoRecycleOfnpreredpacketGet(clt *core.SDKClient, req *servicecenter.TaobaoRecycleOfnpreredpacketGetAPIRequest, session string) (*servicecenter.TaobaoRecycleOfnpreredpacketGetAPIResponse, error) {
	var resp servicecenter.TaobaoRecycleOfnpreredpacketGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
