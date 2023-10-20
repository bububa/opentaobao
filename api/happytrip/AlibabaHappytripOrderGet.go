package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripOrderGet 获取欢行统一订单模型
// alibaba.happytrip.order.get
//
// 通过订单id获取欢行统一订单模型数据
func AlibabaHappytripOrderGet(clt *core.SDKClient, req *happytrip.AlibabaHappytripOrderGetAPIRequest, resp *happytrip.AlibabaHappytripOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
