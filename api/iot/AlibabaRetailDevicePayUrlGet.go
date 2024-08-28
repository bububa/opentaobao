package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// AlibabaRetailDevicePayUrlGet 贩卖机支付二维链接获取
// alibaba.retail.device.payUrl.get
//
// 贩卖机支付二维链接获取
func AlibabaRetailDevicePayUrlGet(ctx context.Context, clt *core.SDKClient, req *iot.AlibabaRetailDevicePayUrlGetAPIRequest, resp *iot.AlibabaRetailDevicePayUrlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
