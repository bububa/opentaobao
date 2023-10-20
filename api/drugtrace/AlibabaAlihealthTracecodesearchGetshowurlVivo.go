package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthtracecodesearchgetshowurlvivo 获取药品扫码落地页vivo
// alibaba.alihealth.tracecodesearch.getshowurl.vivo
//
// 获取药品扫码落地页vivo
func Alibabaalihealthtracecodesearchgetshowurlvivo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest, session string) (*drugtrace.AlibabaalihealthtracecodesearchgetshowurlvivoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthtracecodesearchgetshowurlvivoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
