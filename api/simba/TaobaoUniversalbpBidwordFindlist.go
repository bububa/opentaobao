package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpBidwordFindlist 词列表查询
// taobao.universalbp.bidword.findlist
//
// 根据计划+单元id，查绑定的词列表
func TaobaoUniversalbpBidwordFindlist(clt *core.SDKClient, req *simba.TaobaoUniversalbpBidwordFindlistAPIRequest, session string) (*simba.TaobaoUniversalbpBidwordFindlistAPIResponse, error) {
	var resp simba.TaobaoUniversalbpBidwordFindlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
