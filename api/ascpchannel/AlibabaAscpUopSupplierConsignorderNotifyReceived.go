package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopSupplierConsignorderNotifyReceived 商家仓物流发货推单接单回告
// alibaba.ascp.uop.supplier.consignorder.notify.received
//
// ASCP通过该接口接收商家仓开始接单生产订单对应的物流订单信息
func AlibabaAscpUopSupplierConsignorderNotifyReceived(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest, session string) (*ascpchannel.AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
