package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthCodeGetcodeinfo 码查询功能
// alibaba.alihealth.code.getcodeinfo
//
// 码查询功能
func AlibabaAlihealthCodeGetcodeinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthCodeGetcodeinfoAPIRequest, session string) (*drugtrace.AlibabaAlihealthCodeGetcodeinfoAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthCodeGetcodeinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
