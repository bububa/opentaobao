package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpbidwordsuggestkrlist 关键词建议
// taobao.universalbp.bidword.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词
func Taobaouniversalbpbidwordsuggestkrlist(clt *core.SDKClient, req *simba.TaobaouniversalbpbidwordsuggestkrlistAPIRequest, session string) (*simba.TaobaouniversalbpbidwordsuggestkrlistAPIResponse, error) {
	var resp simba.TaobaouniversalbpbidwordsuggestkrlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
