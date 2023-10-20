package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpwordpackagefindlist 词包列表查询
// taobao.universalbp.wordpackage.findlist
//
// 根据计划+单元id，查绑定的词包列表
func Taobaouniversalbpwordpackagefindlist(clt *core.SDKClient, req *simba.TaobaouniversalbpwordpackagefindlistAPIRequest, session string) (*simba.TaobaouniversalbpwordpackagefindlistAPIResponse, error) {
	var resp simba.TaobaouniversalbpwordpackagefindlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
