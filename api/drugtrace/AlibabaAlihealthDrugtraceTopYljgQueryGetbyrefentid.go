package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljgquerygetbyrefentid 根据企业唯一标识查看企业详细信息
// alibaba.alihealth.drugtrace.top.yljg.query.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func Alibabaalihealthdrugtracetopyljgquerygetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljgquerygetbyrefentidAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljgquerygetbyrefentidAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljgquerygetbyrefentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
