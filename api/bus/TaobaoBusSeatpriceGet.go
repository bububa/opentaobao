package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusseatpriceget 汽车票余票接口
// taobao.bus.seatprice.get
//
// 提供给商家，查询汽车票班次余票
func Taobaobusseatpriceget(clt *core.SDKClient, req *bus.TaobaobusseatpricegetAPIRequest, session string) (*bus.TaobaobusseatpricegetAPIResponse, error) {
	var resp bus.TaobaobusseatpricegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
