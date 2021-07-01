package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

/* AlibabaAscpUopSupplierConsignorderShip
履约单商家仓发货结果回传服务
alibaba.ascp.uop.supplier.consignorder.ship

ERP通过该接口通知商家仓声明销售订单出库信息,支持履约单纬度全部发货的回传（目前不支持分批回传) */
func AlibabaAscpUopSupplierConsignorderShip(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierConsignorderShipAPIRequest, session string) (*ascpchannel.AlibabaAscpUopSupplierConsignorderShipAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpUopSupplierConsignorderShipAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
