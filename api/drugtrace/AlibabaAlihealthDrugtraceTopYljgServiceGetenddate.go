package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljgservicegetenddate 获取服务截止日期
// alibaba.alihealth.drugtrace.top.yljg.service.getenddate
//
// 获取企业服务截止时间
func Alibabaalihealthdrugtracetopyljgservicegetenddate(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljgservicegetenddateAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljgservicegetenddateAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljgservicegetenddateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
