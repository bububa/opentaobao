package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbplabeldmpfinddmpmoduleconfig 查询dmp浮层配置
// taobao.universalbp.label.dmp.finddmpmoduleconfig
//
// 入参账号信息，出参达摩盘相关配置信息
func Taobaouniversalbplabeldmpfinddmpmoduleconfig(clt *core.SDKClient, req *simba.TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIRequest, session string) (*simba.TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIResponse, error) {
	var resp simba.TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
