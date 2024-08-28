package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopSupplierConsignorderNotifyReceived 商家仓物流发货推单接单回告
// alibaba.ascp.uop.supplier.consignorder.notify.received
//
// ASCP通过该接口接收商家仓开始接单生产订单对应的物流订单信息
func AlibabaAscpUopSupplierConsignorderNotifyReceived(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest, resp *ascpchannel.AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
