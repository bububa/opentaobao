package nropen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nropen"
)

// Alibabaascpindustrydisivisonquery 查询服务支持地区列表
// alibaba.ascp.industry.disivison.query
//
// 商家获取服务支持地区
func Alibabaascpindustrydisivisonquery(clt *core.SDKClient, req *nropen.AlibabaascpindustrydisivisonqueryAPIRequest, session string) (*nropen.AlibabaascpindustrydisivisonqueryAPIResponse, error) {
	var resp nropen.AlibabaascpindustrydisivisonqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
