package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeMianUserBind 主账号入驻
// alibaba.alihouse.existinghome.mian.user.bind
//
// 主账号入驻
func AlibabaAlihouseExistinghomeMianUserBind(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeMianUserBindAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeMianUserBindAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeMianUserBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
