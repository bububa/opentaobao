package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushPaperformat 大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat
// alibaba.damai.mev.open.push.paperformat
//
// pushPaperFormat
func AlibabaDamaiMevOpenPushPaperformat(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushPaperformatAPIRequest, resp *damai.AlibabaDamaiMevOpenPushPaperformatAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
