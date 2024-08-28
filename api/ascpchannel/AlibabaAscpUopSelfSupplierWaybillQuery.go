package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopSelfSupplierWaybillQuery 商家仓自营配电子面单取号
// alibaba.ascp.uop.self.supplier.waybill.query
//
// ERP调用打印面单取号接口
func AlibabaAscpUopSelfSupplierWaybillQuery(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSelfSupplierWaybillQueryAPIRequest, resp *ascpchannel.AlibabaAscpUopSelfSupplierWaybillQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
