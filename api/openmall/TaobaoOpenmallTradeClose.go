package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmalltradeclose 关闭订单
// taobao.openmall.trade.close
//
// 关闭订单
func Taobaoopenmalltradeclose(clt *core.SDKClient, req *openmall.TaobaoopenmalltradecloseAPIRequest, session string) (*openmall.TaobaoopenmalltradecloseAPIResponse, error) {
	var resp openmall.TaobaoopenmalltradecloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
