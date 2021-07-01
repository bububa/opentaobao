package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

/* TaobaoTmcAuthGet
TMC授权token
taobao.tmc.auth.get

TMC连接授权Token */
func TaobaoTmcAuthGet(clt *core.SDKClient, req *util.TaobaoTmcAuthGetAPIRequest, session string) (*util.TaobaoTmcAuthGetAPIResponse, error) {
	var resp util.TaobaoTmcAuthGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
