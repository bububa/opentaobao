package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushPaperformat 大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat
// alibaba.damai.mev.open.push.paperformat
//
// pushPaperFormat
func AlibabaDamaiMevOpenPushPaperformat(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushPaperformatAPIRequest, resp *damai.AlibabaDamaiMevOpenPushPaperformatAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
