package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest
履约单纬度的仓缺货回告服务 API请求
alibaba.ascp.uop.supplier.consignorder.outofstock.callback

商家仓履约单纬度的仓缺货回告接口 */
type AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest struct {
	model.Params
	// 缺货回告请求模型
	_consignorderOutofstockCallbackRequest *Consignorderoutofstockcallbackrequest
}

// New
