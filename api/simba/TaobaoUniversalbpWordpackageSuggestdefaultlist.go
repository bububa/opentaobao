package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpwordpackagesuggestdefaultlist 建议默认关键词包
// taobao.universalbp.wordpackage.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词包
func Taobaouniversalbpwordpackagesuggestdefaultlist(clt *core.SDKClient, req *simba.TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest, session string) (*simba.TaobaouniversalbpwordpackagesuggestdefaultlistAPIResponse, error) {
	var resp simba.TaobaouniversalbpwordpackagesuggestdefaultlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
