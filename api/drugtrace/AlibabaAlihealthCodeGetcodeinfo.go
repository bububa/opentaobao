package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthCodeGetcodeinfo 码查询功能
// alibaba.alihealth.code.getcodeinfo
//
// 码查询功能
func AlibabaAlihealthCodeGetcodeinfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthCodeGetcodeinfoAPIRequest, resp *drugtrace.AlibabaAlihealthCodeGetcodeinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
