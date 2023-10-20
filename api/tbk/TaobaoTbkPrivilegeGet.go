package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkprivilegeget 淘宝客-服务商-单品券高效转链
// taobao.tbk.privilege.get
//
// 单品券高效转链API
func Taobaotbkprivilegeget(clt *core.SDKClient, req *tbk.TaobaotbkprivilegegetAPIRequest, session string) (*tbk.TaobaotbkprivilegegetAPIResponse, error) {
	var resp tbk.TaobaotbkprivilegegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
