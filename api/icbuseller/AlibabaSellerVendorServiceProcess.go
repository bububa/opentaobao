package icbuseller

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// AlibabaSellerVendorServiceProcess 服务商客户关联信息
// alibaba.seller.vendor.service.process
//
// 服务商客户关联信息
func AlibabaSellerVendorServiceProcess(ctx context.Context, clt *core.SDKClient, req *icbuseller.AlibabaSellerVendorServiceProcessAPIRequest, resp *icbuseller.AlibabaSellerVendorServiceProcessAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
