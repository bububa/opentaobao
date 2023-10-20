package icbuseller

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// Alibabasellervendororderdetail 国际站服务市场订单详情接口
// alibaba.seller.vendor.order.detail
//
// 国际站服务市场订单列表接口
func Alibabasellervendororderdetail(clt *core.SDKClient, req *icbuseller.AlibabasellervendororderdetailAPIRequest, session string) (*icbuseller.AlibabasellervendororderdetailAPIResponse, error) {
	var resp icbuseller.AlibabasellervendororderdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
