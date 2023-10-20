package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpBidwordSuggestkrlist 关键词建议
// taobao.universalbp.bidword.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词
func TaobaoUniversalbpBidwordSuggestkrlist(clt *core.SDKClient, req *simba.TaobaoUniversalbpBidwordSuggestkrlistAPIRequest, session string) (*simba.TaobaoUniversalbpBidwordSuggestkrlistAPIResponse, error) {
	var resp simba.TaobaoUniversalbpBidwordSuggestkrlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
