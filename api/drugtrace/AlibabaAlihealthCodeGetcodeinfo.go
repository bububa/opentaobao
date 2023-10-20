package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthcodegetcodeinfo 码查询功能
// alibaba.alihealth.code.getcodeinfo
//
// 码查询功能
func Alibabaalihealthcodegetcodeinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthcodegetcodeinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthcodegetcodeinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthcodegetcodeinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
