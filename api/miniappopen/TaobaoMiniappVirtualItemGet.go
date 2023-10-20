package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappvirtualitemget 小程序关联虚拟商品查询
// taobao.miniapp.virtual.item.get
//
// 小程序关联虚拟商品查询
func Taobaominiappvirtualitemget(clt *core.SDKClient, req *miniappopen.TaobaominiappvirtualitemgetAPIRequest, session string) (*miniappopen.TaobaominiappvirtualitemgetAPIResponse, error) {
	var resp miniappopen.TaobaominiappvirtualitemgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
