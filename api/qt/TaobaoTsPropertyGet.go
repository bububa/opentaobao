package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// TaobaoTsPropertyGet 淘宝服务属性查询
// taobao.ts.property.get
//
// 淘宝服务属性查询
func TaobaoTsPropertyGet(clt *core.SDKClient, req *qt.TaobaoTsPropertyGetAPIRequest, resp *qt.TaobaoTsPropertyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
