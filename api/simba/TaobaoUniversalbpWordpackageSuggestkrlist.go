package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpwordpackagesuggestkrlist 关键词包建议
// taobao.universalbp.wordpackage.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词包
func Taobaouniversalbpwordpackagesuggestkrlist(clt *core.SDKClient, req *simba.TaobaouniversalbpwordpackagesuggestkrlistAPIRequest, session string) (*simba.TaobaouniversalbpwordpackagesuggestkrlistAPIResponse, error) {
	var resp simba.TaobaouniversalbpwordpackagesuggestkrlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
