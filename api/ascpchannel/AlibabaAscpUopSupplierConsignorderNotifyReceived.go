package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpuopsupplierconsignordernotifyreceived 商家仓物流发货推单接单回告
// alibaba.ascp.uop.supplier.consignorder.notify.received
//
// ASCP通过该接口接收商家仓开始接单生产订单对应的物流订单信息
func Alibabaascpuopsupplierconsignordernotifyreceived(clt *core.SDKClient, req *ascpchannel.AlibabaascpuopsupplierconsignordernotifyreceivedAPIRequest, session string) (*ascpchannel.AlibabaascpuopsupplierconsignordernotifyreceivedAPIResponse, error) {
	var resp ascpchannel.AlibabaascpuopsupplierconsignordernotifyreceivedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
