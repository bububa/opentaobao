package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiapptemplatequeryapp 查询实例化应用版本
// taobao.miniapp.template.queryapp
//
// 根据模板id和商家信息，查询实例化小程序版本查询
func Taobaominiapptemplatequeryapp(clt *core.SDKClient, req *miniappopen.TaobaominiapptemplatequeryappAPIRequest, session string) (*miniappopen.TaobaominiapptemplatequeryappAPIResponse, error) {
	var resp miniappopen.TaobaominiapptemplatequeryappAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
