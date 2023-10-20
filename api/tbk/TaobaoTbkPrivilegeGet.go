package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkPrivilegeGet 淘宝客-服务商-单品券高效转链
// taobao.tbk.privilege.get
//
// 单品券高效转链API
func TaobaoTbkPrivilegeGet(clt *core.SDKClient, req *tbk.TaobaoTbkPrivilegeGetAPIRequest, session string) (*tbk.TaobaoTbkPrivilegeGetAPIResponse, error) {
	var resp tbk.TaobaoTbkPrivilegeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
