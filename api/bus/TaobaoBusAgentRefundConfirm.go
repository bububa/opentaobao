package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusagentrefundconfirm 汽车票退票和退款二合一接口
// taobao.bus.agent.refund.confirm
//
// 1.商家退票成功后，回调飞猪平台汽车票退票接口，平台进行退票和退款操作。
func Taobaobusagentrefundconfirm(clt *core.SDKClient, req *bus.TaobaobusagentrefundconfirmAPIRequest, session string) (*bus.TaobaobusagentrefundconfirmAPIResponse, error) {
	var resp bus.TaobaobusagentrefundconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
