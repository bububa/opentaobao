package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripOrderGet 获取欢行统一订单模型
// alibaba.happytrip.order.get
//
// 通过订单id获取欢行统一订单模型数据
func AlibabaHappytripOrderGet(clt *core.SDKClient, req *happytrip.AlibabaHappytripOrderGetAPIRequest, session string) (*happytrip.AlibabaHappytripOrderGetAPIResponse, error) {
	var resp happytrip.AlibabaHappytripOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
