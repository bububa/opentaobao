package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkPrivilegeTemporaryGet 淘宝客-服务商-单品券高效转链（临时接口）
// taobao.tbk.privilege.temporary.get
//
// 单品券高效转链API
func TaobaoTbkPrivilegeTemporaryGet(clt *core.SDKClient, req *tbk.TaobaoTbkPrivilegeTemporaryGetAPIRequest, session string) (*tbk.TaobaoTbkPrivilegeTemporaryGetAPIResponse, error) {
	var resp tbk.TaobaoTbkPrivilegeTemporaryGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
