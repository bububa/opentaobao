package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusticketset 出票接口
// taobao.bus.ticket.set
//
// 提供给汽车票商家出票使用
func Taobaobusticketset(clt *core.SDKClient, req *bus.TaobaobusticketsetAPIRequest, session string) (*bus.TaobaobusticketsetAPIResponse, error) {
	var resp bus.TaobaobusticketsetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
