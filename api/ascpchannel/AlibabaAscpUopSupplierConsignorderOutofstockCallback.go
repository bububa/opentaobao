package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopSupplierConsignorderOutofstockCallback 履约单纬度的仓缺货回告服务
// alibaba.ascp.uop.supplier.consignorder.outofstock.callback
//
// 商家仓履约单纬度的仓缺货回告接口
func AlibabaAscpUopSupplierConsignorderOutofstockCallback(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest, resp *ascpchannel.AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
