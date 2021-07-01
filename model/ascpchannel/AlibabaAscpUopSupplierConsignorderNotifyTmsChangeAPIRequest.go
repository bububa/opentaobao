package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest
商家修改运单号 API请求
alibaba.ascp.uop.supplier.consignorder.notify.tms.change

供应商可以通过此接口，对出库回告上报的运单号进行修改，目前一次调用只能支持一个运单号的修改 */
type AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest struct {
	model.Params
	// 修改运单号请求模型
	_modifyMailNoRequest *Modifymailnorequest
}

// New
