package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaorecycleofnpreredpacketget 服务商查询前置补贴红包的最新数据
// taobao.recycle.ofnpreredpacket.get
//
// 服务商查询前置补贴红包的最新数据
func Taobaorecycleofnpreredpacketget(clt *core.SDKClient, req *servicecenter.TaobaorecycleofnpreredpacketgetAPIRequest, session string) (*servicecenter.TaobaorecycleofnpreredpacketgetAPIResponse, error) {
	var resp servicecenter.TaobaorecycleofnpreredpacketgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
