package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeSubAccountBind 子账号入驻
// alibaba.alihouse.existinghome.sub.account.bind
//
// 子账号入驻
func AlibabaAlihouseExistinghomeSubAccountBind(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeSubAccountBindAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeSubAccountBindAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeSubAccountBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
