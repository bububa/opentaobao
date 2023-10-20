package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopenpushpaperformat 大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat
// alibaba.damai.mev.open.push.paperformat
//
// pushPaperFormat
func Alibabadamaimevopenpushpaperformat(clt *core.SDKClient, req *damai.AlibabadamaimevopenpushpaperformatAPIRequest, session string) (*damai.AlibabadamaimevopenpushpaperformatAPIResponse, error) {
	var resp damai.AlibabadamaimevopenpushpaperformatAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
