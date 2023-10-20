package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Alibabaretaildevicevendingregister 贩卖机设备注册
// alibaba.retail.device.vending.register
//
// 贩卖机注册
func Alibabaretaildevicevendingregister(clt *core.SDKClient, req *nrt.AlibabaretaildevicevendingregisterAPIRequest, session string) (*nrt.AlibabaretaildevicevendingregisterAPIResponse, error) {
	var resp nrt.AlibabaretaildevicevendingregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
