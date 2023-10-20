package icbuseller

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// Alibabasellervendororderlist 国际站服务市场订单列表接口
// alibaba.seller.vendor.order.list
//
// 返回服务商在服务市场的客户订单
func Alibabasellervendororderlist(clt *core.SDKClient, req *icbuseller.AlibabasellervendororderlistAPIRequest, session string) (*icbuseller.AlibabasellervendororderlistAPIResponse, error) {
	var resp icbuseller.AlibabasellervendororderlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
