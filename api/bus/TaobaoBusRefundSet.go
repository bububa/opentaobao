package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusrefundset B2B退票申请接口
// taobao.bus.refund.set
//
// B2B业务支持退票
func Taobaobusrefundset(clt *core.SDKClient, req *bus.TaobaobusrefundsetAPIRequest, session string) (*bus.TaobaobusrefundsetAPIResponse, error) {
	var resp bus.TaobaobusrefundsetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
