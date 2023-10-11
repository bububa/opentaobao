package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpWordpackageSuggestdefaultlist 建议默认关键词包
// taobao.universalbp.wordpackage.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词包
func TaobaoUniversalbpWordpackageSuggestdefaultlist(clt *core.SDKClient, req *simba.TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest, session string) (*simba.TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse, error) {
	var resp simba.TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
