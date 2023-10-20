package ship

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ship"
)

// AlitripShipOrderNotify 订单信息回填(出票回调)
// alitrip.ship.order.notify
//
// 此接口为接入商调用飞猪旅行接口回填票号、密码(验证码)等订单信息。接口根据alitripOrderId幂等。若第一次调用失败，后续调用仍然可以回填票号、密码(验证码)成功。第一次调用成功后，后续调用会直接返回第一次的调用结果，不会再产生更新操作。多张票同时出票回填时，保证原子性，只允许全部成功或者全部失败，不能存在部分成功或者失败
func AlitripShipOrderNotify(clt *core.SDKClient, req *ship.AlitripShipOrderNotifyAPIRequest, resp *ship.AlitripShipOrderNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
