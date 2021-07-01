package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

/* AlibabaDamaiMevOpenDeletePaperformat
大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat
alibaba.damai.mev.open.delete.paperformat

deletePaperFormat */
func AlibabaDamaiMevOpenDeletePaperformat(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletePaperformatAPIRequest, session string) (*damai.AlibabaDamaiMevOpenDeletePaperformatAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenDeletePaperformatAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
