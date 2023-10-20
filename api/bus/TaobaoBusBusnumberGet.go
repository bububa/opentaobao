package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusbusnumberget 汽车票车次查询
// taobao.bus.busnumber.get
//
// 提供汽车票车次查询服务
func Taobaobusbusnumberget(clt *core.SDKClient, req *bus.TaobaobusbusnumbergetAPIRequest, session string) (*bus.TaobaobusbusnumbergetAPIResponse, error) {
	var resp bus.TaobaobusbusnumbergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
