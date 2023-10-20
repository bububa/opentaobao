package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpIndustryUopSupplierConsignoder 商家推单
// alibaba.ascp.industry.uop.supplier.consignoder
//
// 商家推单
func AlibabaAscpIndustryUopSupplierConsignoder(clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryUopSupplierConsignoderAPIRequest, resp *ascpchannel.AlibabaAscpIndustryUopSupplierConsignoderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
