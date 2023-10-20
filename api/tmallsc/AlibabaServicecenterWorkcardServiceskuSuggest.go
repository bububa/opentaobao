package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabaservicecenterworkcardserviceskusuggest 服务商反馈需要履行的服务项
// alibaba.servicecenter.workcard.servicesku.suggest
//
// 服务商反馈需要履行的服务项
func Alibabaservicecenterworkcardserviceskusuggest(clt *core.SDKClient, req *tmallsc.AlibabaservicecenterworkcardserviceskusuggestAPIRequest, session string) (*tmallsc.AlibabaservicecenterworkcardserviceskusuggestAPIResponse, error) {
	var resp tmallsc.AlibabaservicecenterworkcardserviceskusuggestAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
