package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkPrivilegeGet 淘宝客-服务商-单品券高效转链
// taobao.tbk.privilege.get
//
// 单品券高效转链API
func TaobaoTbkPrivilegeGet(clt *core.SDKClient, req *tbk.TaobaoTbkPrivilegeGetAPIRequest, resp *tbk.TaobaoTbkPrivilegeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
