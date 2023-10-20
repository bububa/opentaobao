package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusagentmultiplerefundconfirm 综合交通多次退款接口
// taobao.bus.agent.multiple.refund.confirm
//
// 此接口支持多次按照单客进行多次退款操作，只进行退款操作。
func Taobaobusagentmultiplerefundconfirm(clt *core.SDKClient, req *bus.TaobaobusagentmultiplerefundconfirmAPIRequest, session string) (*bus.TaobaobusagentmultiplerefundconfirmAPIResponse, error) {
	var resp bus.TaobaobusagentmultiplerefundconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
