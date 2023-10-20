package icbuseller

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// Alibabasellervendorserviceprocess 服务商客户关联信息
// alibaba.seller.vendor.service.process
//
// 服务商客户关联信息
func Alibabasellervendorserviceprocess(clt *core.SDKClient, req *icbuseller.AlibabasellervendorserviceprocessAPIRequest, session string) (*icbuseller.AlibabasellervendorserviceprocessAPIResponse, error) {
	var resp icbuseller.AlibabasellervendorserviceprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
