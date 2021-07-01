package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

/* TmallDeviceTradePrecreate
智能硬件上预创建天猫订单
tmall.device.trade.precreate

智能硬件上预创建天猫订单。
1，use_open_price不再需要传入，使用unit_price传入单价。
2，订单默认5分钟自动关闭，没有付款的订单在手机淘宝不可见。
3，同一个码只运行一个用户扫码，多个用户扫一个码会报错 订单不存在。 */
func TmallDeviceTradePrecreate(clt *core.SDKClient, req *iot.TmallDeviceTradePrecreateAPIRequest, session string) (*iot.TmallDeviceTradePrecreateAPIResponse, error) {
	var resp iot.TmallDeviceTradePrecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
