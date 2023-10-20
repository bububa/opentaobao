package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpbidwordfindlist 词列表查询
// taobao.universalbp.bidword.findlist
//
// 根据计划+单元id，查绑定的词列表
func Taobaouniversalbpbidwordfindlist(clt *core.SDKClient, req *simba.TaobaouniversalbpbidwordfindlistAPIRequest, session string) (*simba.TaobaouniversalbpbidwordfindlistAPIResponse, error) {
	var resp simba.TaobaouniversalbpbidwordfindlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
