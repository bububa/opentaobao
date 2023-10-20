package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// AlibabaretaildevicepayUrlget 贩卖机支付二维链接获取
// alibaba.retail.device.payUrl.get
//
// 贩卖机支付二维链接获取
func AlibabaretaildevicepayUrlget(clt *core.SDKClient, req *iot.AlibabaretaildevicepayUrlgetAPIRequest, session string) (*iot.AlibabaretaildevicepayUrlgetAPIResponse, error) {
	var resp iot.AlibabaretaildevicepayUrlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
