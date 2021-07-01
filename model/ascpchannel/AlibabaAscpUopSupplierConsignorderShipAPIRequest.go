package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopSupplierConsignorderShipAPIRequest
履约单商家仓发货结果回传服务 API请求
alibaba.ascp.uop.supplier.consignorder.ship

ERP通过该接口通知商家仓声明销售订单出库信息,支持履约单纬度全部发货的回传（目前不支持分批回传) */
type AlibabaAscpUopSupplierConsignorderShipAPIRequest struct {
	model.Params
	// 发货回传请求模型
	_consignorderShipRequest *Consignordershiprequest
}

// New
