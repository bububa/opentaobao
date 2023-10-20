package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeletePaperformat 大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat
// alibaba.damai.mev.open.delete.paperformat
//
// deletePaperFormat
func AlibabaDamaiMevOpenDeletePaperformat(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletePaperformatAPIRequest, resp *damai.AlibabaDamaiMevOpenDeletePaperformatAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
