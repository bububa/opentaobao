package nropen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpIndustryDisivisonQueryAPIRequest
查询服务支持地区列表 API请求
alibaba.ascp.industry.disivison.query

商家获取服务支持地区 */
type AlibabaAscpIndustryDisivisonQueryAPIRequest struct {
	model.Params
	// 服务编码
	_serviceCode string
}

// New
