package icbuseller

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// Alibabasellervendorservicevendorprocess 服务商客户关联信息
// alibaba.seller.vendor.service.vendorprocess
//
// 服务商客户关联信息
func Alibabasellervendorservicevendorprocess(clt *core.SDKClient, req *icbuseller.AlibabasellervendorservicevendorprocessAPIRequest, session string) (*icbuseller.AlibabasellervendorservicevendorprocessAPIResponse, error) {
	var resp icbuseller.AlibabasellervendorservicevendorprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
