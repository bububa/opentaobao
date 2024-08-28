package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthTracecodesearchGetshowurlVivo 获取药品扫码落地页vivo
// alibaba.alihealth.tracecodesearch.getshowurl.vivo
//
// 获取药品扫码落地页vivo
func AlibabaAlihealthTracecodesearchGetshowurlVivo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest, resp *drugtrace.AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
