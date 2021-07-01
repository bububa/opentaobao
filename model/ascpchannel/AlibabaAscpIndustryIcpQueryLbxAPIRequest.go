package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpIndustryIcpQueryLbxAPIRequest
icp订单号查询lbx订单号 API请求
alibaba.ascp.industry.icp.query.lbx

根据icp订单号查询lbx订单号 */
type AlibabaAscpIndustryIcpQueryLbxAPIRequest struct {
	model.Params
	// icps订单号
	_icpOrderCode string
}

// New
