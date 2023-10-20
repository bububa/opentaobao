package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetoplsydservicegetenddate 获取企业服务截止时间
// alibaba.alihealth.drugtrace.top.lsyd.service.getenddate
//
// 获取企业服务截止时间
func Alibabaalihealthdrugtracetoplsydservicegetenddate(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetoplsydservicegetenddateAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetoplsydservicegetenddateAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetoplsydservicegetenddateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
