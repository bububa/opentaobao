package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharityuserexternalauth 外部用户授权
// alibaba.charity.user.external.auth
//
// 外部用户授权
func Alibabacharityuserexternalauth(clt *core.SDKClient, req *charity.AlibabacharityuserexternalauthAPIRequest, session string) (*charity.AlibabacharityuserexternalauthAPIResponse, error) {
	var resp charity.AlibabacharityuserexternalauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
