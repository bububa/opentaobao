package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiapptemplateinstantiate 构建实例化应用
// taobao.miniapp.template.instantiate
//
// 实例化saas化的小程序
func Taobaominiapptemplateinstantiate(clt *core.SDKClient, req *miniappopen.TaobaominiapptemplateinstantiateAPIRequest, session string) (*miniappopen.TaobaominiapptemplateinstantiateAPIResponse, error) {
	var resp miniappopen.TaobaominiapptemplateinstantiateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
