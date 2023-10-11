package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpLabelFindconfiglist 查询可用标签id信息
// taobao.universalbp.label.findconfiglist
//
// 入参账号信息，出参可用标签id，用于下游接口入参
func TaobaoUniversalbpLabelFindconfiglist(clt *core.SDKClient, req *simba.TaobaoUniversalbpLabelFindconfiglistAPIRequest, session string) (*simba.TaobaoUniversalbpLabelFindconfiglistAPIResponse, error) {
	var resp simba.TaobaoUniversalbpLabelFindconfiglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
