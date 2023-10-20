package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeMianUserBind 主账号入驻
// alibaba.alihouse.existinghome.mian.user.bind
//
// 主账号入驻
func AlibabaAlihouseExistinghomeMianUserBind(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeMianUserBindAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeMianUserBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
