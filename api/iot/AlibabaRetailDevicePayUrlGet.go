package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

/* AlibabaRetailDevicePayUrlGet
贩卖机支付二维链接获取
alibaba.retail.device.payUrl.get

贩卖机支付二维链接获取 */
func AlibabaRetailDevicePayUrlGet(clt *core.SDKClient, req *iot.AlibabaRetailDevicePayUrlGetAPIRequest, session string) (*iot.AlibabaRetailDevicePayUrlGetAPIResponse, error) {
	var resp iot.AlibabaRetailDevicePayUrlGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
