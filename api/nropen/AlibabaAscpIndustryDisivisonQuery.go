package nropen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nropen"
)

// AlibabaAscpIndustryDisivisonQuery 查询服务支持地区列表
// alibaba.ascp.industry.disivison.query
//
// 商家获取服务支持地区
func AlibabaAscpIndustryDisivisonQuery(clt *core.SDKClient, req *nropen.AlibabaAscpIndustryDisivisonQueryAPIRequest, session string) (*nropen.AlibabaAscpIndustryDisivisonQueryAPIResponse, error) {
	var resp nropen.AlibabaAscpIndustryDisivisonQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
