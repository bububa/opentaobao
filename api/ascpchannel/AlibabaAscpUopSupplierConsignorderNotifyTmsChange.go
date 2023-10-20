package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopSupplierConsignorderNotifyTmsChange 商家修改运单号
// alibaba.ascp.uop.supplier.consignorder.notify.tms.change
//
// 供应商可以通过此接口，对出库回告上报的运单号进行修改，目前一次调用只能支持一个运单号的修改
func AlibabaAscpUopSupplierConsignorderNotifyTmsChange(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest, resp *ascpchannel.AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
