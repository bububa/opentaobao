package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpuopsupplierconsignordernotifytmschange 商家修改运单号
// alibaba.ascp.uop.supplier.consignorder.notify.tms.change
//
// 供应商可以通过此接口，对出库回告上报的运单号进行修改，目前一次调用只能支持一个运单号的修改
func Alibabaascpuopsupplierconsignordernotifytmschange(clt *core.SDKClient, req *ascpchannel.AlibabaascpuopsupplierconsignordernotifytmschangeAPIRequest, session string) (*ascpchannel.AlibabaascpuopsupplierconsignordernotifytmschangeAPIResponse, error) {
	var resp ascpchannel.AlibabaascpuopsupplierconsignordernotifytmschangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
