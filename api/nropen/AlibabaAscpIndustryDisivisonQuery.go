package nropen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nropen"
)

// AlibabaAscpIndustryDisivisonQuery 查询服务支持地区列表
// alibaba.ascp.industry.disivison.query
//
// 商家获取服务支持地区
func AlibabaAscpIndustryDisivisonQuery(clt *core.SDKClient, req *nropen.AlibabaAscpIndustryDisivisonQueryAPIRequest, resp *nropen.AlibabaAscpIndustryDisivisonQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
