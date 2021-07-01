package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpIndustryUopSupplierConsignoderAPIRequest
商家推单 API请求
alibaba.ascp.industry.uop.supplier.consignoder

商家推单 */
type AlibabaAscpIndustryUopSupplierConsignoderAPIRequest struct {
	model.Params
	// 发货主单信息
	_erpNormalConsignOrderRequest *Erpnormalconsignorderrequest
}

// New
