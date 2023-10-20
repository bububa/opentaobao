package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetoplsydquerygetbyrefentid 根据企业唯一标识查看企业详细信息
// alibaba.alihealth.drugtrace.top.lsyd.query.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func Alibabaalihealthdrugtracetoplsydquerygetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetoplsydquerygetbyrefentidAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetoplsydquerygetbyrefentidAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetoplsydquerygetbyrefentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
