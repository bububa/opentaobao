package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpWordpackageSuggestkrlist 关键词包建议
// taobao.universalbp.wordpackage.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词包
func TaobaoUniversalbpWordpackageSuggestkrlist(clt *core.SDKClient, req *simba.TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest, resp *simba.TaobaoUniversalbpWordpackageSuggestkrlistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
