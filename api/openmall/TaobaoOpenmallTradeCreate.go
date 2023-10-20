package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmalltradecreate 创建订单
// taobao.openmall.trade.create
//
// 创建Openmall订单
func Taobaoopenmalltradecreate(clt *core.SDKClient, req *openmall.TaobaoopenmalltradecreateAPIRequest, session string) (*openmall.TaobaoopenmalltradecreateAPIResponse, error) {
	var resp openmall.TaobaoopenmalltradecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
