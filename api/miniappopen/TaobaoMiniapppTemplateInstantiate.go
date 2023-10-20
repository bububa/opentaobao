package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappptemplateinstantiate （已废弃）构建实例化应用
// taobao.miniappp.template.instantiate
//
// 实例化saas化的小程序
func Taobaominiappptemplateinstantiate(clt *core.SDKClient, req *miniappopen.TaobaominiappptemplateinstantiateAPIRequest, session string) (*miniappopen.TaobaominiappptemplateinstantiateAPIResponse, error) {
	var resp miniappopen.TaobaominiappptemplateinstantiateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
