package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpBidwordSuggestkrlist 关键词建议
// taobao.universalbp.bidword.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词
func TaobaoUniversalbpBidwordSuggestkrlist(clt *core.SDKClient, req *simba.TaobaoUniversalbpBidwordSuggestkrlistAPIRequest, resp *simba.TaobaoUniversalbpBidwordSuggestkrlistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
