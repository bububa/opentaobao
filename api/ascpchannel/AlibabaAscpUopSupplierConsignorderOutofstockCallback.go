package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopSupplierConsignorderOutofstockCallback 履约单纬度的仓缺货回告服务
// alibaba.ascp.uop.supplier.consignorder.outofstock.callback
//
// 商家仓履约单纬度的仓缺货回告接口
func AlibabaAscpUopSupplierConsignorderOutofstockCallback(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest, resp *ascpchannel.AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
