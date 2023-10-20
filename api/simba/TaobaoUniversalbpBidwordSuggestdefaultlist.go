package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpBidwordSuggestdefaultlist 建议默认关键词
// taobao.universalbp.bidword.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词
func TaobaoUniversalbpBidwordSuggestdefaultlist(clt *core.SDKClient, req *simba.TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest, session string) (*simba.TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse, error) {
	var resp simba.TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
