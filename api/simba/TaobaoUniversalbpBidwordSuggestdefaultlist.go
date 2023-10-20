package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpbidwordsuggestdefaultlist 建议默认关键词
// taobao.universalbp.bidword.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词
func Taobaouniversalbpbidwordsuggestdefaultlist(clt *core.SDKClient, req *simba.TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest, session string) (*simba.TaobaouniversalbpbidwordsuggestdefaultlistAPIResponse, error) {
	var resp simba.TaobaouniversalbpbidwordsuggestdefaultlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
