package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbplabelfindconfiglist 查询可用标签id信息
// taobao.universalbp.label.findconfiglist
//
// 入参账号信息，出参可用标签id，用于下游接口入参
func Taobaouniversalbplabelfindconfiglist(clt *core.SDKClient, req *simba.TaobaouniversalbplabelfindconfiglistAPIRequest, session string) (*simba.TaobaouniversalbplabelfindconfiglistAPIResponse, error) {
	var resp simba.TaobaouniversalbplabelfindconfiglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
