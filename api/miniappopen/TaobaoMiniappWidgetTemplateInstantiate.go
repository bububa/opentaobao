package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappwidgettemplateinstantiate 小部件实例化接口
// taobao.miniapp.widget.template.instantiate
//
// 小部件实例化接口
func Taobaominiappwidgettemplateinstantiate(clt *core.SDKClient, req *miniappopen.TaobaominiappwidgettemplateinstantiateAPIRequest, session string) (*miniappopen.TaobaominiappwidgettemplateinstantiateAPIResponse, error) {
	var resp miniappopen.TaobaominiappwidgettemplateinstantiateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
