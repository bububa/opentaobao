package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopSupplierReverseorderCreate 商家ERP发起创建销退单服务
// alibaba.ascp.uop.supplier.reverseorder.create
//
// 商家在收到消费者实物退货后，在ERP发起创建销退单服务
func AlibabaAscpUopSupplierReverseorderCreate(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierReverseorderCreateAPIRequest, resp *ascpchannel.AlibabaAscpUopSupplierReverseorderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
