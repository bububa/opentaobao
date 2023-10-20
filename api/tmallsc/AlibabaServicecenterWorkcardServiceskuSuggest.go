package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaServicecenterWorkcardServiceskuSuggest 服务商反馈需要履行的服务项
// alibaba.servicecenter.workcard.servicesku.suggest
//
// 服务商反馈需要履行的服务项
func AlibabaServicecenterWorkcardServiceskuSuggest(clt *core.SDKClient, req *tmallsc.AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest, resp *tmallsc.AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
