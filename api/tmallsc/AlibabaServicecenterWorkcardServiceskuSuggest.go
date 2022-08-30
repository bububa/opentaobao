package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaServicecenterWorkcardServiceskuSuggest 服务商反馈需要履行的服务项
// alibaba.servicecenter.workcard.servicesku.suggest
//
// 服务商反馈需要履行的服务项
func AlibabaServicecenterWorkcardServiceskuSuggest(clt *core.SDKClient, req *tmallsc.AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest, session string) (*tmallsc.AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse, error) {
	var resp tmallsc.AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
