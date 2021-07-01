package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest
商家仓物流发货推单接单回告 API请求
alibaba.ascp.uop.supplier.consignorder.notify.received

ASCP通过该接口接收商家仓开始接单生产订单对应的物流订单信息 */
type AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest struct {
	model.Params
	// qimen.alibaba.ascp.uop.consignorder.notify报文中的supplierId字段值
	_supplierId string
	// qimen.alibaba.ascp.uop.consignorder.notify报文中bizOrderCode履约单号
	_bizOrderCode string
	// 业务请求时间
	_bizTime string
}

// New
