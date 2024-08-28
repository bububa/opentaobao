package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopSupplierConsignorderShip 履约单商家仓发货结果回传服务
// alibaba.ascp.uop.supplier.consignorder.ship
//
// ERP通过该接口通知商家仓声明销售订单出库信息,支持履约单纬度全部发货的回传（目前不支持分批回传)
func AlibabaAscpUopSupplierConsignorderShip(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierConsignorderShipAPIRequest, resp *ascpchannel.AlibabaAscpUopSupplierConsignorderShipAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
