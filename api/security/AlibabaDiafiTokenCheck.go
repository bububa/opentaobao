package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabadiafitokencheck 天朗token校验API
// alibaba.diafi.token.check
//
// 天朗token校验
func Alibabadiafitokencheck(clt *core.SDKClient, req *security.AlibabadiafitokencheckAPIRequest, session string) (*security.AlibabadiafitokencheckAPIResponse, error) {
	var resp security.AlibabadiafitokencheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
