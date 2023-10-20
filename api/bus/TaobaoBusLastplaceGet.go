package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobuslastplaceget 获取目的地数据
// taobao.bus.lastplace.get
//
// 传入城市 获取对应的目的地
func Taobaobuslastplaceget(clt *core.SDKClient, req *bus.TaobaobuslastplacegetAPIRequest, session string) (*bus.TaobaobuslastplacegetAPIResponse, error) {
	var resp bus.TaobaobuslastplacegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
