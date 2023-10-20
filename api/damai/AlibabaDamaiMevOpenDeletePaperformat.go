package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopendeletepaperformat 大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat
// alibaba.damai.mev.open.delete.paperformat
//
// deletePaperFormat
func Alibabadamaimevopendeletepaperformat(clt *core.SDKClient, req *damai.AlibabadamaimevopendeletepaperformatAPIRequest, session string) (*damai.AlibabadamaimevopendeletepaperformatAPIResponse, error) {
	var resp damai.AlibabadamaimevopendeletepaperformatAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
