package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqloginpreventionresultfetch 获取登录保护结果
// alibaba.security.jaq.loginprevention.result.fetch
//
// 获取登录保护结果
func Alibabasecurityjaqloginpreventionresultfetch(clt *core.SDKClient, req *security.AlibabasecurityjaqloginpreventionresultfetchAPIRequest, session string) (*security.AlibabasecurityjaqloginpreventionresultfetchAPIResponse, error) {
	var resp security.AlibabasecurityjaqloginpreventionresultfetchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
