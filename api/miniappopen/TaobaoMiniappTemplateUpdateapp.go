package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiapptemplateupdateapp 更新实例化应用
// taobao.miniapp.template.updateapp
//
// 商家应用c端模板实例化小程序更新，生成新的版本，但不会自动上线新版本
func Taobaominiapptemplateupdateapp(clt *core.SDKClient, req *miniappopen.TaobaominiapptemplateupdateappAPIRequest, session string) (*miniappopen.TaobaominiapptemplateupdateappAPIResponse, error) {
	var resp miniappopen.TaobaominiapptemplateupdateappAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
